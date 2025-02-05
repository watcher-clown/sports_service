package event

import (
	"github.com/garyburd/redigo/redis"
	"sports_service/dao"
	"sports_service/global/app/log"
	"sports_service/global/rdskey"
	"sports_service/redismq/protocol"
	"sports_service/util"
	"time"
)

// 事件消息
func PushEventMsg(msg []byte) {
	log.Log.Infof("event_trace: 事件推送, msg:%s", string(msg))
	//body := newEvent(userId, nickname, cover, content, eventType)
	conn := dao.RedisPool().Get()
	defer conn.Close()

	if _, err := redis.Int(conn.Do("LPUSH", rdskey.MSG_PUSH_EVENT_KEY, msg)); err != nil {
		log.Log.Infof("event_trace: msg push fail, err:%s", err)
	}

	return
}

// toUserId 接收者id
// nickname 发送者昵称
func NewEvent(toUserId, composeId, nickname, cover, content string, eventType int32) []byte {
	event := new(protocol.Event)
	event.UserId = toUserId
	event.EventType = eventType
	event.Ts = time.Now().Unix()

	data := new(protocol.PushData)
	data.NickName = nickname
	data.Cover = cover
	data.Content = content
	data.ComposeId = composeId

	msg, _ := util.JsonFast.Marshal(data)
	event.Data = msg
	b, err := util.JsonFast.Marshal(event)
	if err != nil {
		log.Log.Errorf("event_trace: marshal err:%s", err)
	}

	return b
}
