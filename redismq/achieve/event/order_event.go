package event

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"sports_service/dao"
	"sports_service/global/app/log"
	"sports_service/global/consts"
	"sports_service/global/rdskey"
	"sports_service/models/mappointment"
	"sports_service/models/morder"
	producer "sports_service/redismq/event"
	"sports_service/redismq/protocol"
	"sports_service/util"
	"time"
)

func LoopPopOrderEvent() {
	for !closing {
		conn := dao.RedisPool().Get()
		values, err := redis.Values(conn.Do("BRPOP", rdskey.MSG_ORDER_EVENT_KEY, 0))
		conn.Close()
		if err != nil {
			log.Log.Errorf("redisMq_trace: loopPop event fail, err:%s", err)
			// 防止出现错误时 频繁刷日志
			time.Sleep(time.Second)
			continue
		}

		if len(values) < 2 {
			log.Log.Errorf("redisMq_trace: invalid values, len:%d, values:%+v", len(values), values)
		}

		bts, ok := values[1].([]byte)
		if !ok {
			log.Log.Errorf("redisMq_trace: value[1] unSupport type")
			continue
		}

		if err := OrderEventConsumer(bts); err != nil {
			log.Log.Errorf("redisMq_trace: event consumer fail, err:%s, msg:%s", err, string(bts))
			// 重新投递消息
			producer.PushOrderEventMsg(bts)
		}

	}
}

func OrderEventConsumer(bts []byte) error {
	event := protocol.Event{}
	if err := util.JsonFast.Unmarshal(bts, &event); err != nil {
		log.Log.Errorf("redisMq_trace: proto unmarshal order event err:%s", err)
		return err
	}

	if err := handleOrderEvent(event); err != nil {
		log.Log.Errorf("handleOrderEvent err:%s", err)
		return err
	}

	return nil
}

// 订单相关事件处理
func handleOrderEvent(event protocol.Event) error {
	info := &protocol.OrderData{}
	if err := util.JsonFast.Unmarshal(event.Data, info); err != nil {
		log.Log.Errorf("redisMq_trace: proto unmarshal order data err:%s", err)
		return nil
	}
	log.Log.Infof("redisMq_trace: info:%+v", info)

	// 如果超时处理时间 > 当前时间 则重新入队列
	if info.ProcessTm > time.Now().Unix() {
		log.Log.Errorf("redisMq_trace: orderId:%s, requeue", info.OrderId)
		return errors.New("requeue")
	}

	switch event.EventType {
	// 预约场馆超时
	case consts.ORDER_EVENT_VENUE_TIME_OUT:
		if err := orderTimeOut(consts.APPOINTMENT_VENUE, info.OrderId); err != nil {
			return err
		}

	// 预约私教超时
	case consts.ORDER_EVENT_COACH_TIME_OUT:
		if err := orderTimeOut(consts.APPOINTMENT_COACH, info.OrderId); err != nil {
			return err
		}

	// 预约课程超时
	case consts.ORDER_EVENT_COURSE_TIME_OUT:
		if err := orderTimeOut(consts.APPOINTMENT_COURSE, info.OrderId); err != nil {
			return err
		}

	default:
		log.Log.Errorf("redisMq_trace: unsupported eventType, orderId:%s, eventType:%d", info.OrderId, event.EventType)
	}

	return nil
}

// 订单超时
func orderTimeOut(appointmentType int, orderId string) error {
	session := dao.VenueEngine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		log.Log.Errorf("redisMq_trace: session begin err:%s, orderId:%s", err, orderId)
		return err
	}

	orderModel := morder.NewOrderModel(session)
	ok, err := orderModel.GetOrder(orderId)
	if !ok || err != nil {
		log.Log.Errorf("redisMq_trace: get order info fail, err:%s, ok:%v, orderId:%s", err, ok, orderId)
		session.Rollback()
		return nil
	}

	// 订单状态 != 0 (待支付) 表示 订单 已设为超时/已支付/已完成 等等...
	if orderModel.Order.Status != consts.ORDER_TYPE_WAIT {
		log.Log.Errorf("redisMq_trace: don't need to change，orderId:%s, status:%d", orderId,
			orderModel.Order.Status)
		session.Rollback()
		return nil
	}

	now := int(time.Now().Unix())
	orderModel.Order.UpdateAt = now
	orderModel.Order.Status = consts.ORDER_TYPE_UNPAID
	// 更新订单状态为 超时未支付
	affected, err := orderModel.UpdateOrderStatus(orderId, consts.ORDER_TYPE_WAIT)
	if affected != 1 || err != nil {
		log.Log.Errorf("redisMq_trace: update order status fail, orderId:%s, err:%s", orderId, err)
		session.Rollback()
		return errors.New("update order status fail")
	}

	orderModel.OrderProduct.Status = consts.ORDER_TYPE_UNPAID
	orderModel.OrderProduct.UpdateAt = now
	// 更新订单商品流水状态
	if _, err = orderModel.UpdateOrderProductStatus(orderId, consts.ORDER_TYPE_WAIT); err != nil {
		log.Log.Errorf("redisMq_trace: update order product status fail, err:%s, affected:%d, orderId:%s", err, affected, orderId)
		session.Rollback()
		return errors.New("update order product status fail")
	}

	// 获取订单对应的预约流水
	amodel := mappointment.NewAppointmentModel(session)
	list, err := amodel.GetAppointmentRecordByOrderId(orderId)
	if err != nil {
		log.Log.Errorf("redisMq_trace: get appointment record by orderId fail, orderId:%s, err:%s", orderId, err)
		session.Rollback()
		return err
	}

	for _, record := range list {
		// 归还对应节点的冻结库存
		affected, err = amodel.RevertStockNum(record.TimeNode, record.Date, record.PurchasedNum*-1, now,
			record.AppointmentType, int(record.VenueId))
		if affected != 1 || err != nil {
			log.Log.Errorf("redisMq_trace: update stock info fail, orderId:%s, err:%s, affected:%d, id:%d", orderId, err, affected, record.Id)
			session.Rollback()
			return errors.New("update stock info fail")
		}
	}

	// 更新订单对应的预约流水状态
	//if err := amodel.UpdateAppointmentRecordStatus(orderId, now, 0); err != nil {
	//	log.Log.Errorf("payNotify_trace: update order product status fail, err:%s, orderId:%s", err, orderId)
	//	session.Rollback()
	//	return err
	//}

	// 更新标签状态[废弃]
	amodel.Labels.Status = 1
	if _, err = amodel.UpdateLabelsStatus(orderId, 0); err != nil {
		log.Log.Errorf("redisMq_trace: update labels status fail, orderId:%s, err:%s", orderId, err)
		session.Rollback()
		return errors.New("update label status fail")
	}

	session.Commit()
	return nil
}
