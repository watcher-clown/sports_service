package event

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	nsqsvc "github.com/nsqio/go-nsq"
	"sports_service/app/config"
	"sports_service/dao"
	"sports_service/global/app/log"
	"sports_service/global/consts"
	"sports_service/models"
	"sports_service/models/mcomment"
	"sports_service/models/mlike"
	"sports_service/models/mnotify"
	"sports_service/models/muser"
	"sports_service/models/umeng"
	"sports_service/nsqlx/protocol"
	"sports_service/util"
	"time"
)

func connectConsumer(channel string) (*nsqsvc.Consumer, error) {
	nsqConfig := nsqsvc.NewConfig()
	fmt.Printf("初始化 topic: %v,channel:%v", consts.EVENT_TOPIC, channel)
	consumer, err := nsqsvc.NewConsumer(consts.EVENT_TOPIC, channel, nsqConfig)
	if err != nil {
		log.Log.Errorf("new consumer err:%v", err)
		return consumer, err
	}

	return consumer, err
}

func EventConsumer(channel string) (consumer *nsqsvc.Consumer) {
	consumer, err := connectConsumer(channel)
	if err != nil {
		panic(fmt.Sprintf("consumer conn err:%s", err))
	}

	consumer.AddHandler(nsqsvc.HandlerFunc(NsqHandler))

	err = consumer.ConnectToNSQD(config.Global.NsqAddr)
	if err != nil {
		log.Log.Errorf("ConnectToNSQD err:%s", err)
	}

	return
}

func NsqHandler(msg *nsqsvc.Message) error {
	event := new(protocol.Event)
	if err := util.JsonFast.Unmarshal(msg.Body, event); err != nil {
		log.Log.Errorf("event_trace: proto unmarshal event err:%s", err)
		return err
	}

	if err := handleEvent(event); err != nil {
		msg.RequeueWithoutBackoff(time.Second * 3)
		log.Log.Errorf("handleEvent err:%s", err)
		return err
	}

	return nil
}

func handleEvent(event *protocol.Event) error {
	info := &protocol.Data{}
	if err := util.JsonFast.Unmarshal(event.Data, info); err != nil {
		log.Log.Errorf("event_trace: proto unmarshal data err:%s", err)
		return nil
	}

	session := dao.AppEngine.NewSession()
	defer session.Close()
	umodel := muser.NewUserModel(session)
	user := umodel.FindUserByUserid(event.UserId)
	if user == nil {
		log.Log.Errorf("event_trace: user not found, userId:%s", event.UserId)
		return nil
	}

	nmodel := mnotify.NewNotifyModel(session)
	// 系统消息未读数
	sysNum := nmodel.GetUnreadSystemMsgNum(event.UserId)

	// likedNum 未读的被点赞的数量 atNum 未读的被@的数量
	var (
		likedNum, atNum int64
	)
	// 获取用户上次读取被点赞列表的时间
	readTm, err := nmodel.GetReadBeLikedTime(event.UserId)
	if err == nil || err == redis.ErrNil {
		if readTm == "" {
			readTm = "0"
		}

		lmodel := mlike.NewLikeModel(session)
		// 获取未读的被点赞的数量
		likedNum = lmodel.GetUnreadBeLikedCount(event.UserId, readTm)
	}

	// 获取用户上次读取被@列表数据的时间
	readAt, err := nmodel.GetReadAtTime(event.UserId)
	if err == nil || err == redis.ErrNil {
		if readAt == "" {
			readAt = "0"
		}

		cmodel := mcomment.NewCommentModel(session)
		// 获取未读的被@的数量
		atNum = cmodel.GetUnreadAtCount(event.UserId, readAt)
	}

	// 总未读数
	unReadNum := sysNum + likedNum + atNum

	var (
		content string
		msgType int32
	)

	log.Log.Errorf("event_trace: event:%+v", event)
	setting := nmodel.GetUserNotifySetting(event.UserId)

	var pushSet int
	switch event.EventType {
	// 系统类
	case consts.SYSTEM_MSG:

	// 活动类
	case consts.ACTIVITY_MSG:

	// 视频点赞
	case consts.VIDEO_LIKE_MSG:
		content = fmt.Sprintf("%s 赞了你的作品", info.NickName)
		msgType = int32(consts.MSG_TYPE_VIDEO_LIKE_NOTIFY)
		pushSet = setting.ThumbUpPushSet
	// 评论/回复 点赞
	case consts.VIDEO_COMMENT_LIKE_MSG:
		content = fmt.Sprintf("%s 赞了你的评论 @%s", info.NickName, info.Content)
		msgType = int32(consts.MSG_TYPE_VIDEO_COMMENT_LIKE_NOTIFY)
		pushSet = setting.ThumbUpPushSet
	// 收藏视频
	case consts.COLLECT_VIDEO_MSG:
		content = fmt.Sprintf("%s 收藏了你的作品", info.NickName)
		msgType = int32(consts.MSG_TYPE_VIDEO_COLLECT_NOTIFY)
	// 关注用户
	case consts.FOCUS_USER_MSG:
		content = fmt.Sprintf("%s 关注了你", info.NickName)
		msgType = int32(consts.MSG_TYPE_FOCUS_NOTIFY)
		pushSet = setting.AttentionPushSet
	// 关注的用户发布视频
	case consts.FOCUS_USER_PUBLISH_VIDEO_MSG:
		content = fmt.Sprintf("你关注的 %s 发布了新视频", info.NickName)
		msgType = int32(consts.MSG_TYPE_FOCUS_PUBLISH_VIDEO_NOTIFY)
		pushSet = setting.AttentionPushSet
	// 视频评论
	case consts.VIDEO_COMMENT_MSG:
		content = fmt.Sprintf("%s 评论了你的作品", info.NickName)
		msgType = int32(consts.MSG_TYPE_VIDEO_COMMENT_NOTIFY)
		pushSet = setting.CommentPushSet
	// 视频回复
	case consts.VIDEO_REPLY_MSG:
		// @ 用户发布的评论
		content = fmt.Sprintf("%s 回复了你的评论 @%s", info.NickName, info.Content)
		msgType = int32(consts.MSG_TYPE_VIDEO_REPLY_NOTIFY)
		pushSet = setting.CommentPushSet
	default:
		log.Log.Errorf("event_trace: unsupported eventType, eventType:%d", event.EventType)
		return nil

	}

	// 0为接收推送 1为拒绝接收
	if pushSet == 0 {
		// 推送通知
		PushNotify(user, "", content, info.Cover, msgType, unReadNum)
	}

	return nil
}

// 推送通知
func PushNotify(user *models.User, title, content, cover string, msgType int32, unreadNum int64) {
	extra := make(map[string]interface{}, 0)
	extra["unread_num"] = unreadNum
	title = "X-FLY官方"
	// android推送
	if user.DeviceType == int(consts.ANDROID_PLATFORM) && user.DeviceToken != "" {
		client := umeng.New()
		if err := client.PushUnicastNotify(msgType, umeng.FPV_ANDROID, user.DeviceToken, title, content, cover, extra, nil); err != nil {
			log.Log.Errorf("event_trace: push notify by user err:%s", err)
		}
	}

	// iOS推送
	if user.DeviceType == int(consts.IOS_PLATFORM) && user.DeviceToken != "" {
		client := umeng.New()
		if err := client.PushUnicastNotify(msgType, umeng.FPV_IOS, user.DeviceToken, title, content, cover, extra, nil); err != nil {
			log.Log.Errorf("event_trace: push notify by user err:%s", err)
		}
	}
}
