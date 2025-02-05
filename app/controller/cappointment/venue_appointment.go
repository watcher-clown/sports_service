package cappointment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"sports_service/dao"
	"sports_service/global/app/errdef"
	"sports_service/global/app/log"
	"sports_service/global/consts"
	"sports_service/models"
	"sports_service/models/mappointment"
	"sports_service/models/muser"
	"sports_service/models/mvenue"
	"sports_service/tools/tencentCloud"
	"sports_service/util"
	"time"
)

type VenueAppointmentModule struct {
	context *gin.Context
	engine  *xorm.Session
	user    *muser.UserModel
	venue   *mvenue.VenueModel
	*base
}

func NewVenue(c *gin.Context) *VenueAppointmentModule {
	venueSocket := dao.VenueEngine.NewSession()
	defer venueSocket.Close()
	appSocket := dao.AppEngine.NewSession()
	defer appSocket.Close()
	return &VenueAppointmentModule{
		context: c,
		user:    muser.NewUserModel(appSocket),
		venue:   mvenue.NewVenueModel(venueSocket),
		engine:  venueSocket,
		base:    New(venueSocket),
	}
}

// 场馆选项 todo：暂时只有一个场馆
func (svc *VenueAppointmentModule) Options(relatedId int64) (int, interface{}) {
	list, err := svc.venue.GetVenueList()
	if err != nil {
		return errdef.ERROR, nil
	}

	if list == nil {
		return errdef.SUCCESS, []interface{}{}
	}

	res := make([]*mappointment.Options, len(list))
	for index, item := range list {
		info := &mappointment.Options{
			Id:           item.Id,
			Name:         item.VenueName,
			Instructions: item.Instructions,
		}

		res[index] = info
	}

	return errdef.SUCCESS, res
}

// todo:
// 预约场馆
// 1 判断库存数据是否存在 不存在写入 如并发情况 遇写入失败[mysql err 1062 唯一索引约束错误] 说明已有用户成功预约 同时已写入库存数据
// 2 进行库存更新 判断库存是否足够 库存不够 直接返回各节点最新库存量
// 3 如库存均足够 则判断是否充值会员时长
// 4 如充值会员时长 预约的时间 按价格从高至低 抵扣时长 且每个时间节点最多只可抵扣一次 [能抵扣则进行抵扣 预约数量-1] 并 计算抵扣时长后的 订单总价
// 5 如未充值会员时长 或 会员时长不足 则剩余预约 按售价 * 预约数量 计算订单总价
// 6 记录订单、订单商品流水、预约流水
func (svc *VenueAppointmentModule) Appointment(params *mappointment.AppointmentReq) (int, interface{}) {
	log.Log.Infof("params:%+v", params)
	if err := svc.engine.Begin(); err != nil {
		log.Log.Errorf("venue_trace: session begin fail, err:%s", err)
		return errdef.ERROR, nil
	}

	if len(params.Infos) == 0 {
		log.Log.Errorf("venue_trace: infos len%d", len(params.Infos))
		svc.engine.Rollback()
		return errdef.APPOINTMENT_INVALID_INFO, nil
	}

	user := svc.user.FindUserByUserid(params.UserId)
	if user == nil {
		svc.engine.Rollback()
		return errdef.USER_NOT_EXISTS, nil
	}

	list, err := svc.GetAppointmentConfByIds(params.Ids)
	if err != nil {
		log.Log.Errorf("venue_trace: get appointment conf by ids fail, err:%s, ids:%v", err, params.Ids)
		svc.engine.Rollback()
		return errdef.APPOINTMENT_QUERY_NODE_FAIL, nil
	}

	if len(list) != len(params.Infos) {
		log.Log.Errorf("venue_trace: length not match, list len:%d, infos len:%d", len(list), len(params.Infos))
		svc.engine.Rollback()
		return errdef.APPOINTMENT_INVALID_NODE_ID, nil
	}

	ok, err := svc.venue.GetVenueInfoById(fmt.Sprint(params.RelatedId))
	if !ok || err != nil {
		log.Log.Errorf("venue_trace: get venue info fail, err:%s", err)
		svc.engine.Rollback()
		return errdef.VENUE_NOT_EXISTS, nil
	}

	orderId := util.NewOrderId()
	now := int(time.Now().Unix())
	svc.Extra.ProductImg = tencentCloud.BucketURI(svc.venue.Venue.PromotionPic)

	if err := svc.AppointmentProcess(user.UserId, orderId, params.RelatedId, params.WeekNum, params.LabelIds, params.Infos); err != nil {
		log.Log.Errorf("venue_trace: appointment fail, err:%s", err)
		svc.engine.Rollback()
		return errdef.APPOINTMENT_PROCESS_FAIL, nil
	}

	svc.Extra.Id = params.RelatedId
	svc.Extra.Name = svc.venue.Venue.VenueName
	svc.Extra.VenueName = svc.venue.Venue.VenueName
	svc.Extra.WeekCn = util.GetWeekCn(params.WeekNum)
	svc.Extra.MobileNum = util.HideMobileNum(fmt.Sprint(user.MobileNum))
	svc.Extra.TmCn = util.ResolveTime(svc.Extra.TotalTm)
	svc.Extra.TotalSalesPrice = svc.Extra.TotalAmount
	svc.Extra.Channel = params.Channel

	// 用户选择抵扣时长
	if params.IsDiscount == 1 {
		// 库存都足够 则判断用户是否充值会员时长
		// 存在会员数据 则需要先抵扣时间
		if err := svc.VipDeductionProcess(user.UserId, list); err != nil {
			log.Log.Errorf("venue_trace: vip deduction process fail, err:%s", err)
			svc.engine.Rollback()
			return errdef.APPOINTMENT_VIP_DEDUCTION, nil
		}
	}

	log.Log.Errorf("venue_trace: extra Info:%+v", svc.Extra)
	// 库存不足 返回最新数据 事务回滚
	if !svc.Extra.IsEnough {
		log.Log.Errorf("venue_trace: rollback, isEnough:%v, reqType:%d", svc.Extra.IsEnough, params.ReqType)
		svc.engine.Rollback()
		return errdef.APPOINTMENT_NOT_ENOUGH_STOCK, svc.Extra
	}

	// 查询数据 则返回200
	if params.ReqType != 2 {
		svc.engine.Rollback()
		return errdef.SUCCESS, svc.Extra
	}

	svc.Extra.WriteOffCode = fmt.Sprint(util.GetSnowId())
	// 添加订单
	if err := svc.AddOrder(orderId, user.UserId, "预约场馆", now, consts.ORDER_TYPE_APPOINTMENT_VENUE); err != nil {
		log.Log.Errorf("venue_trace: add order fail, err:%s", err)
		svc.engine.Rollback()
		return errdef.ORDER_ADD_FAIL, nil
	}

	// 添加订单商品流水
	if err := svc.AddOrderProducts(); err != nil {
		log.Log.Errorf("venue_trace: add order products fail, err:%s", err)
		svc.engine.Rollback()
		return errdef.ORDER_PRODUCT_ADD_FAIL, nil
	}

	// 添加预约记录流水
	if err := svc.AddAppointmentRecord(); err != nil {
		log.Log.Errorf("venue_trace: add appointment record fail, err:%s", err)
		svc.engine.Rollback()
		return errdef.APPOINTMENT_ADD_RECORD_FAIL, nil
	}

	// 记录需处理支付超时的订单
	if _, err := svc.order.RecordOrderId(orderId); err != nil {
		log.Log.Errorf("venue_trace: record orderId fail, err:%s", err)
		svc.engine.Rollback()
		return errdef.APPOINTMENT_RECORD_ORDER_FAIL, nil
	}

	svc.engine.Commit()

	svc.Extra.OrderId = orderId
	svc.Extra.PayDuration = consts.PAYMENT_DURATION
	// 超时
	//redismq.PushOrderEventMsg(redismq.NewOrderEvent(user.UserId, svc.Extra.OrderId, int64(svc.order.Order.CreateAt) + svc.Extra.PayDuration,
	//	consts.ORDER_EVENT_VENUE_TIME_OUT))
	return errdef.SUCCESS, svc.Extra
}

// 预约场馆时间选项
func (svc *VenueAppointmentModule) AppointmentOptions() (int, interface{}) {
	date := svc.GetDateById(svc.DateId, consts.FORMAT_DATE)
	if date == "" {
		return errdef.ERROR, nil
	}

	condition, err := svc.GetQueryCondition()
	if err != nil {
		log.Log.Errorf("venue_trace: get query condition fail, err:%s", err)
		return errdef.ERROR, nil
	}

	list, err := svc.GetAppointmentOptions(condition)
	if err != nil {
		log.Log.Errorf("venue_trace: get options fail, err:%s", err)
		return errdef.ERROR, nil
	}

	if len(list) == 0 {
		return errdef.SUCCESS, []interface{}{}
	}

	res := make([]*mappointment.OptionsInfo, 0)
	for _, item := range list {
		info := svc.SetAppointmentOptionsRes(date, item)
		if info == nil {
			log.Log.Error("venue_trace: options res nil")
			continue
		}

		ok, err := svc.venue.GetVenueInfoById(fmt.Sprint(item.VenueId))
		if err != nil {
			log.Log.Errorf("venue_trace: get venue info by id fail, err:%s", err)
		}

		if ok {
			info.Name = svc.venue.Venue.VenueName
		}

		svc.appointment.Labels.TimeNode = item.TimeNode
		svc.appointment.Labels.Date = date
		svc.appointment.Labels.VenueId = item.VenueId
		labels, err := svc.appointment.GetVenueUserLabels()
		if err != nil {
			log.Log.Errorf("venue_trace: get venue user lables fail, err:%s", err)
		}

		info.Labels = make([]*mappointment.LabelInfo, len(labels))
		for key, val := range labels {
			label := &mappointment.LabelInfo{
				UserId:    val.UserId,
				LabelId:   val.LabelId,
				LabelName: val.LabelName,
			}

			info.Labels[key] = label
		}

		svc.appointment.Record.AppointmentType = 0
		svc.appointment.Record.TimeNode = item.TimeNode
		svc.appointment.Record.VenueId = item.VenueId
		svc.appointment.Record.Date = date
		records, err := svc.appointment.GetAppointmentRecord()
		if err != nil {
			log.Log.Errorf("venue_trace: get appointment record fail, err:%s", err)
		}

		ok, err = svc.venue.GetRecommendInfoById(fmt.Sprint(info.RecommendType))
		if !ok || err != nil {
			log.Log.Errorf("venue_trace: get recommend info by id fail, id:%d", info.RecommendType)
		}
		info.RecommendName = svc.venue.Recommend.Name

		info.ReservedUsers = make([]*mappointment.SeatInfo, 0)
		if len(records) > 0 {
			for _, val := range records {
				if val.Date != date || val.TimeNode != item.TimeNode {
					continue
				}

				seatInfo := &mappointment.SeatInfo{}
				if err := util.JsonFast.UnmarshalFromString(val.SeatInfo, seatInfo); err == nil {
					user := svc.user.FindUserByUserid(val.UserId)
					if user != nil {
						seatInfo.NickName = user.NickName
						seatInfo.Avatar = tencentCloud.BucketURI(user.Avatar)
					}

					info.ReservedUsers = append(info.ReservedUsers, seatInfo)
				} else {
					log.Log.Errorf("venue_trace: unmarshal seat info fail, err:%s", err)
				}

				//uinfo := &mappointment.SeatInfo{
				//	UserId: val.UserId,
				//}
				//
				//user := svc.user.FindUserByUserid(val.UserId)
				//if user != nil {
				//	uinfo.NickName = user.NickName
				//	uinfo.Avatar = user.Avatar
				//	uinfo.SeatNo = 1
				//}
				//
				//info.ReservedUsers = append(info.ReservedUsers, uinfo)
				//
				//if val.PurchasedNum > 1 {
				//	for i := 0; i < val.PurchasedNum-1; i++ {
				//		uinfo.SeatNo += 1
				//		info.ReservedUsers = append(info.ReservedUsers, uinfo)
				//	}
				//}

			}
		}

		res = append(res, info)
	}

	return errdef.SUCCESS, res
}

// 预约详情
func (svc *VenueAppointmentModule) AppointmentDetail() (int, interface{}) {

	return 4000, nil
}

// 场馆预约日期配置
func (svc *VenueAppointmentModule) AppointmentDate() (int, interface{}) {
	return errdef.SUCCESS, svc.AppointmentDateInfo(6, consts.APPOINTMENT_VENUE)
}

// 会员抵扣流程
func (svc *VenueAppointmentModule) VipDeductionProcess(userId string, list []*models.VenueAppointmentInfo) error {
	vip, err := svc.appointment.GetVenueVipInfo(userId)
	if err != nil {
		// 查询失败
		return err
	}

	if vip == nil {
		return nil
	}

	if vip.Duration <= 0 {
		return nil
	}

	// 查看会员是否过期 已过期会员无法抵扣
	if vip.EndTm < time.Now().Unix() {
		return nil
	}

	// 如果是会员 且 会员时长 > 0
	// 开始走抵扣流程 预约的时间节点[多个] 按价格从高至低 开始抵扣 每个时间节点最多只可抵扣一次
	for key, val := range list {
		if val.Duration <= 0 {
			continue
		}

		// 是否足够抵扣
		affected, err := svc.appointment.UpdateVenueVipInfo(val.Duration*-1, val.VenueId, userId)
		if err != nil {
			log.Log.Errorf("venue_trace: update vip duration fail, err:%s", err)
			return err
		}

		// 会员时长不够 查看下一个预约节点 是否可抵扣
		if affected == 0 {
			continue
		}

		// 足够抵扣 则记录抵扣的记录
		if affected == 1 {
			// 抵扣一个 则 减去一个的售价
			svc.recordMp[val.Id][0].DeductionNum = affected
			svc.recordMp[val.Id][0].DeductionTm = int64(val.Duration)
			svc.recordMp[val.Id][0].DeductionAmount = int64(val.CurAmount)
			svc.Extra.TotalDeductionTm += val.Duration
			// 订单总金额 = 商品总价 - 抵扣金额
			svc.Extra.TotalAmount = svc.Extra.TotalAmount - val.CurAmount
			svc.Extra.IsDeduct = true
			// 当前节点付款金额 = 当前节点总价 - 当前抵扣金额
			svc.orderMp[val.Id].Amount = svc.orderMp[val.Id].Amount - val.CurAmount
			if len(svc.Extra.TimeNodeInfo) <= key {
				svc.Extra.TimeNodeInfo[key].DeductionTm = svc.recordMp[val.Id][0].DeductionTm
			}
		}
	}

	return nil
}

// 获取标签信息
func (svc *VenueAppointmentModule) GetLabelInfo() (int, []*models.VenuePersonalLabelConf) {
	list, err := svc.appointment.GetUserLabelConf()
	if err != nil {
		return errdef.ERROR, []*models.VenuePersonalLabelConf{}
	}

	return errdef.SUCCESS, list
}
