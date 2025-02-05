package clike

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"sports_service/dao"
	"sports_service/global/app/errdef"
	"sports_service/global/app/log"
	"sports_service/global/consts"
	"sports_service/models/mattention"
	"sports_service/models/mcollect"
	"sports_service/models/mcomment"
	"sports_service/models/mconfigure"
	"sports_service/models/minformation"
	"sports_service/models/mlike"
	"sports_service/models/mposting"
	"sports_service/models/muser"
	"sports_service/models/mvideo"
	redismq "sports_service/redismq/event"
	"sports_service/tools/tencentCloud"
	"sports_service/util"
	"strings"
	"time"
)

type LikeModule struct {
	context     *gin.Context
	engine      *xorm.Session
	user        *muser.UserModel
	like        *mlike.LikeModel
	video       *mvideo.VideoModel
	comment     *mcomment.CommentModel
	attention   *mattention.AttentionModel
	collect     *mcollect.CollectModel
	post        *mposting.PostingModel
	information *minformation.InformationModel
	config      *mconfigure.ConfigModel
}

func New(c *gin.Context) LikeModule {
	socket := dao.AppEngine.NewSession()
	defer socket.Close()
	return LikeModule{
		context:     c,
		user:        muser.NewUserModel(socket),
		video:       mvideo.NewVideoModel(socket),
		comment:     mcomment.NewCommentModel(socket),
		like:        mlike.NewLikeModel(socket),
		attention:   mattention.NewAttentionModel(socket),
		collect:     mcollect.NewCollectModel(socket),
		post:        mposting.NewPostingModel(socket),
		information: minformation.NewInformationModel(socket),
		config:      mconfigure.NewConfigModel(socket),
		engine:      socket,
	}
}

// 点赞视频
func (svc *LikeModule) GiveLikeForVideo(userId string, videoId int64) int {
	// 开启事务
	if err := svc.engine.Begin(); err != nil {
		log.Log.Errorf("like_trace: session begin err:%s", err)
		return errdef.ERROR
	}

	// 查询用户是否存在
	user := svc.user.FindUserByUserid(userId)
	if user == nil {
		log.Log.Errorf("like_trace: user not found, userId:%s", userId)
		svc.engine.Rollback()
		return errdef.USER_NOT_EXISTS
	}

	// 查找视频是否存在
	video := svc.video.FindVideoById(fmt.Sprint(videoId))
	if video == nil {
		log.Log.Errorf("like_trace: like video not found, videoId:%d", videoId)
		svc.engine.Rollback()
		return errdef.LIKE_VIDEO_NOT_EXISTS
	}

	if fmt.Sprint(video.Status) != consts.VIDEO_AUDIT_SUCCESS {
		log.Log.Errorf("like_trace: like video audit failure, videoId:%d", videoId)
		svc.engine.Rollback()
		return errdef.LIKE_VIDEO_NOT_EXISTS
	}

	// 获取点赞的视频信息
	info := svc.like.GetLikeInfo(userId, videoId, consts.TYPE_VIDEOS)
	// 是否已点赞
	// 已点赞
	if info != nil && info.Status == consts.ALREADY_GIVE_LIKE {
		log.Log.Errorf("like_trace: already give like, userId:%s, videoId:%d", userId, videoId)
		svc.engine.Rollback()
		return errdef.LIKE_ALREADY_EXISTS
	}

	now := int(time.Now().Unix())
	score := svc.config.GetActionScore(int(consts.WORK_TYPE_VIDEO), consts.ACTION_TYPE_FABULOUS)
	// 更新视频点赞总计 +1
	if err := svc.video.UpdateVideoLikeNum(videoId, now, consts.CONFIRM_OPERATE, score); err != nil {
		log.Log.Errorf("like_trace: update video like num err:%s", err)
		svc.engine.Rollback()
		return errdef.LIKE_VIDEO_FAIL
	}

	// 未点赞
	// 记录存在 且 状态为 未点赞 更新状态为 已点赞
	if info != nil && info.Status == consts.NOT_GIVE_LIKE {
		info.Status = consts.ALREADY_GIVE_LIKE
		info.CreateAt = now
		if err := svc.like.UpdateLikeStatus(); err != nil {
			log.Log.Errorf("like_trace: update like status err:%s", err)
			svc.engine.Rollback()
			return errdef.LIKE_VIDEO_FAIL
		}

	} else {
		// 添加点赞记录
		if err := svc.like.AddGiveLikeByType(userId, video.UserId, videoId, consts.ALREADY_GIVE_LIKE, consts.TYPE_VIDEOS); err != nil {
			log.Log.Errorf("like_trace: add like video record err:%s", err)
			svc.engine.Rollback()
			return errdef.LIKE_VIDEO_FAIL
		}
	}

	svc.engine.Commit()

	// 发送视频点赞推送
	//event.PushEventMsg(config.Global.AmqpDsn, video.UserId, user.NickName, video.Cover, "", consts.VIDEO_LIKE_MSG)
	redismq.PushEventMsg(redismq.NewEvent(video.UserId, fmt.Sprint(video.VideoId), user.NickName, video.Cover, "", consts.VIDEO_LIKE_MSG))
	// 视频置顶事件
	redismq.PushTopEventMsg(redismq.NewTopEvent(video.UserId, fmt.Sprint(video.VideoId), consts.EVENT_SET_TOP_VIDEO))
	return errdef.SUCCESS
}

// 取消点赞（视频）
func (svc *LikeModule) CancelLikeForVideo(userId string, videoId int64) int {
	// 开启事务
	if err := svc.engine.Begin(); err != nil {
		log.Log.Errorf("like_trace: session begin err:%s", err)
		return errdef.ERROR
	}

	// 查询用户是否存在
	if user := svc.user.FindUserByUserid(userId); user == nil {
		log.Log.Errorf("like_trace: user not found, userId:%s", userId)
		svc.engine.Rollback()
		return errdef.USER_NOT_EXISTS
	}

	// 查找视频是否存在
	if video := svc.video.FindVideoById(fmt.Sprint(videoId)); video == nil {
		log.Log.Errorf("like_trace: cancel like video not found, videoId:%d", videoId)
		svc.engine.Rollback()
		return errdef.LIKE_VIDEO_NOT_EXISTS
	}

	// 获取点赞的信息 判断是否已点赞 记录不存在 则 未点过赞
	info := svc.like.GetLikeInfo(userId, videoId, consts.TYPE_VIDEOS)
	if info == nil {
		log.Log.Errorf("like_trace: record not found, not give like, userId:%s, videoId:%d", userId, videoId)
		svc.engine.Rollback()
		return errdef.LIKE_RECORD_NOT_EXISTS
	}

	// 状态 ！= 已点赞 提示重复操作
	if info.Status != consts.ALREADY_GIVE_LIKE {
		log.Log.Errorf("like_trace: already cancel like, userId:%s, videoId:%d", userId, videoId)
		svc.engine.Rollback()
		return errdef.LIKE_REPEAT_CANCEL
	}

	now := int(time.Now().Unix())
	score := svc.config.GetActionScore(int(consts.WORK_TYPE_VIDEO), consts.ACTION_TYPE_FABULOUS)
	// 更新视频点赞总计 -1
	if err := svc.video.UpdateVideoLikeNum(videoId, now, consts.CANCEL_OPERATE, score*-1); err != nil {
		log.Log.Errorf("like_trace: update video like num err:%s", err)
		svc.engine.Rollback()
		return errdef.LIKE_CANCEL_FAIL
	}

	info.Status = consts.NOT_GIVE_LIKE
	info.CreateAt = now
	// 更新状态 未点赞
	if err := svc.like.UpdateLikeStatus(); err != nil {
		log.Log.Errorf("like_trace: update like status err:%s", err)
		svc.engine.Rollback()
		return errdef.LIKE_CANCEL_FAIL
	}

	svc.engine.Commit()

	return errdef.SUCCESS
}

// 获取用户点赞的视频列表
func (svc *LikeModule) GetUserLikeVideos(userId string, page, size int) []*mvideo.VideosInfoResp {
	if userId == "" {
		log.Log.Errorf("like_trace: userId not exists! userId:%s", userId)
		return []*mvideo.VideosInfoResp{}
	}

	offset := (page - 1) * size
	infos := svc.like.GetUserLikeVideos(userId, offset, size)
	if len(infos) == 0 {
		return []*mvideo.VideosInfoResp{}
	}

	// mp key videoId   value 用户视频点赞的时间
	mp := make(map[int64]int)
	// 当前页所有视频id
	videoIds := make([]string, len(infos))
	for index, like := range infos {
		mp[like.TypeId] = like.CreateAt
		videoIds[index] = fmt.Sprint(like.TypeId)
	}

	vids := strings.Join(videoIds, ",")
	// 获取点赞的视频列表信息
	videoList := svc.video.FindVideoListByIds(vids)
	if len(videoList) == 0 {
		log.Log.Errorf("like_trace: not found like video list info, len:%d, videoIds:%s", len(videoList), vids)
		return []*mvideo.VideosInfoResp{}
	}

	// 重新组装数据
	list := make([]*mvideo.VideosInfoResp, len(videoList))
	for index, video := range videoList {
		resp := new(mvideo.VideosInfoResp)
		resp.VideoId = video.VideoId
		resp.Title = util.TrimHtml(video.Title)
		resp.Describe = util.TrimHtml(video.Describe)
		resp.Cover = video.Cover
		resp.VideoAddr = svc.video.AntiStealingLink(video.VideoAddr)
		resp.IsRecommend = video.IsRecommend
		resp.IsTop = video.IsTop
		resp.VideoDuration = video.VideoDuration
		resp.VideoWidth = video.VideoWidth
		resp.VideoHeight = video.VideoHeight
		resp.CreateAt = video.CreateAt
		resp.UserId = video.UserId
		// 获取用户信息
		if user := svc.user.FindUserByUserid(video.UserId); user != nil {
			resp.Avatar = tencentCloud.BucketURI(user.Avatar)
			resp.Nickname = user.NickName
			// 是否关注
			attentionInfo := svc.attention.GetAttentionInfo(userId, video.UserId)
			if attentionInfo != nil {
				resp.IsAttention = attentionInfo.Status
			}
		}

		likeAt, ok := mp[video.VideoId]
		if ok {
			// 用户给视频点赞的时间
			resp.OpTime = likeAt
		}

		list[index] = resp
	}

	return list
}

// 评论点赞 包含视频评论、帖子评论
func (svc *LikeModule) GiveLikeForComment(userId string, commentId, commentType int64) int {
	switch commentType {
	case consts.COMMENT_TYPE_VIDEO:
		return svc.GiveLikeForVideoComment(userId, commentId)

	case consts.COMMENT_TYPE_POST:
		return svc.GiveLikeForPostComment(userId, commentId)

	case consts.COMMENT_TYPE_INFORMATION:
		return svc.GiveLikeForInformationComment(userId, commentId)

	default:
		log.Log.Errorf("comment_trace: invalid commentType:%d", commentType)
		return errdef.INVALID_PARAMS
	}
}

// 点赞资讯评论
func (svc *LikeModule) GiveLikeForInformationComment(userId string, commentId int64) int {
	// 开启事务
	if err := svc.engine.Begin(); err != nil {
		log.Log.Errorf("like_trace: session begin err:%s", err)
		return errdef.ERROR
	}

	// 查询用户是否存在
	user := svc.user.FindUserByUserid(userId)
	if user == nil {
		log.Log.Errorf("like_trace: user not found, userId:%s", userId)
		svc.engine.Rollback()
		return errdef.USER_NOT_EXISTS
	}

	// 查找评论是否存在
	comment := svc.comment.GetInformationCommentById(fmt.Sprint(commentId))
	if comment == nil {
		log.Log.Errorf("like_trace: like comment not found, commentId:%d", commentId)
		svc.engine.Rollback()
		return errdef.LIKE_COMMENT_NOT_EXISTS
	}

	if code := svc.GiveLike(userId, comment.UserId, commentId, consts.TYPE_INFORMATION_COMMENT); code != errdef.SUCCESS {
		svc.engine.Rollback()
		return code
	}

	svc.engine.Commit()

	return errdef.SUCCESS
}

// 点赞视频评论
func (svc *LikeModule) GiveLikeForVideoComment(userId string, commentId int64) int {
	// 开启事务
	if err := svc.engine.Begin(); err != nil {
		log.Log.Errorf("like_trace: session begin err:%s", err)
		return errdef.ERROR
	}

	// 查询用户是否存在
	user := svc.user.FindUserByUserid(userId)
	if user == nil {
		log.Log.Errorf("like_trace: user not found, userId:%s", userId)
		svc.engine.Rollback()
		return errdef.USER_NOT_EXISTS
	}

	// 查找评论是否存在
	comment := svc.comment.GetVideoCommentById(fmt.Sprint(commentId))
	if comment == nil {
		log.Log.Errorf("like_trace: like comment not found, commentId:%d", commentId)
		svc.engine.Rollback()
		return errdef.LIKE_COMMENT_NOT_EXISTS
	}

	if code := svc.GiveLike(userId, comment.UserId, commentId, consts.TYPE_VIDEO_COMMENT); code != errdef.SUCCESS {
		svc.engine.Rollback()
		return code
	}

	svc.engine.Commit()

	// 发送评论点赞推送
	//event.PushEventMsg(config.Global.AmqpDsn, comment.UserId, user.NickName, "", comment.Content, consts.VIDEO_COMMENT_LIKE_MSG)
	redismq.PushEventMsg(redismq.NewEvent(comment.UserId, fmt.Sprint(comment.VideoId), user.NickName, "", comment.Content, consts.VIDEO_COMMENT_LIKE_MSG))

	return errdef.SUCCESS
}

func (svc *LikeModule) GiveLike(userId, toUserId string, commentId int64, commentType int) int {
	// 获取点赞的评论信息
	info := svc.like.GetLikeInfo(userId, commentId, commentType)
	// 是否已点赞
	// 已点赞
	if info != nil && info.Status == consts.ALREADY_GIVE_LIKE {
		log.Log.Errorf("like_trace: already give like, userId:%s, commentId:%d", userId, commentId)
		svc.engine.Rollback()
		return errdef.LIKE_ALREADY_EXISTS
	}

	now := int(time.Now().Unix())
	// 未点赞
	// 记录存在 且 状态为 未点赞 更新状态为 已点赞
	if info != nil && info.Status == consts.NOT_GIVE_LIKE {
		info.Status = consts.ALREADY_GIVE_LIKE
		info.CreateAt = now
		if err := svc.like.UpdateLikeStatus(); err != nil {
			log.Log.Errorf("like_trace: update like comment status err:%s", err)
			return errdef.LIKE_COMMENT_FAIL
		}

	} else {
		// 添加点赞记录
		if err := svc.like.AddGiveLikeByType(userId, toUserId, commentId, consts.ALREADY_GIVE_LIKE, commentType); err != nil {
			log.Log.Errorf("like_trace: add like comment record err:%s", err)
			return errdef.LIKE_COMMENT_FAIL
		}
	}

	return errdef.SUCCESS

}

// 取消点赞（包含视频评论、帖子评论）
func (svc *LikeModule) CancelLikeForComment(userId string, commentId, commentType int64) int {
	switch commentType {
	case consts.COMMENT_TYPE_VIDEO:
		return svc.CancelLikeForVideoComment(userId, commentId)

	case consts.COMMENT_TYPE_POST:
		return svc.CancelLikeForPostComment(userId, commentId)

	case consts.COMMENT_TYPE_INFORMATION:
		return svc.CancelLikeForInformationComment(userId, commentId)
	default:
		log.Log.Errorf("comment_trace: invalid commentType:%d", commentType)
		return errdef.INVALID_PARAMS
	}
}

// 取消点赞（视频评论）
func (svc *LikeModule) CancelLikeForVideoComment(userId string, commentId int64) int {
	// 开启事务
	if err := svc.engine.Begin(); err != nil {
		log.Log.Errorf("like_trace: session begin err:%s", err)
		return errdef.ERROR
	}

	// 查询用户是否存在
	if user := svc.user.FindUserByUserid(userId); user == nil {
		log.Log.Errorf("like_trace: user not found, userId:%s", userId)
		svc.engine.Rollback()
		return errdef.USER_NOT_EXISTS
	}

	// 查找评论是否存在
	if comment := svc.comment.GetVideoCommentById(fmt.Sprint(commentId)); comment == nil {
		log.Log.Errorf("like_trace: cancel like comment not found, commentId:%d", commentId)
		svc.engine.Rollback()
		return errdef.LIKE_COMMENT_NOT_EXISTS
	}

	if code := svc.CancelLike(userId, commentId, consts.TYPE_VIDEO_COMMENT); code != errdef.SUCCESS {
		svc.engine.Rollback()
		return code
	}

	svc.engine.Commit()

	return errdef.SUCCESS
}

// 评论取消点赞
func (svc *LikeModule) CancelLike(userId string, commentId int64, commentType int) int {
	// 获取点赞的信息 判断是否已点赞 记录不存在 则 未点过赞
	info := svc.like.GetLikeInfo(userId, commentId, commentType)
	if info == nil {
		log.Log.Errorf("like_trace: record not found, not give like, userId:%s, commentId:%d", userId, commentId)
		return errdef.LIKE_RECORD_NOT_EXISTS
	}

	// 状态 ！= 已点赞 提示重复操作
	if info.Status != consts.ALREADY_GIVE_LIKE {
		log.Log.Errorf("like_trace: already cancel like, userId:%s, commentId:%d", userId, commentId)
		return errdef.LIKE_REPEAT_CANCEL
	}

	now := int(time.Now().Unix())
	info.Status = consts.NOT_GIVE_LIKE
	info.CreateAt = now
	// 更新状态 未点赞
	if err := svc.like.UpdateLikeStatus(); err != nil {
		log.Log.Errorf("like_trace: update like status err:%s", err)
		return errdef.LIKE_CANCEL_FAIL
	}

	return errdef.SUCCESS
}

// 点赞帖子
func (svc *LikeModule) GiveLikeForPost(userId string, postId int64) int {
	// 开启事务
	if err := svc.engine.Begin(); err != nil {
		log.Log.Errorf("like_trace: session begin err:%s", err)
		return errdef.ERROR
	}

	// 查询用户是否存在
	user := svc.user.FindUserByUserid(userId)
	if user == nil {
		log.Log.Errorf("like_trace: user not found, userId:%s", userId)
		svc.engine.Rollback()
		return errdef.USER_NOT_EXISTS
	}

	// 查找帖子是否存在
	post, err := svc.post.GetPostById(fmt.Sprint(postId))
	if err != nil || post == nil {
		svc.engine.Rollback()
		return errdef.LIKE_POST_NOT_EXISTS
	}

	if fmt.Sprint(post.Status) != consts.POST_AUDIT_SUCCESS {
		log.Log.Errorf("like_trace: post not found, postId:%d", postId)
		svc.engine.Rollback()
		return errdef.LIKE_POST_FAIL
	}

	// 获取点赞的帖子信息
	info := svc.like.GetLikeInfo(userId, postId, consts.TYPE_POSTS)
	// 是否已点赞
	// 已点赞
	if info != nil && info.Status == consts.ALREADY_GIVE_LIKE {
		log.Log.Errorf("like_trace: already give like, userId:%s, postId:%d", userId, postId)
		svc.engine.Rollback()
		return errdef.LIKE_ALREADY_EXISTS
	}

	now := int(time.Now().Unix())
	score := svc.config.GetActionScore(int(consts.WORK_TYPE_POST), consts.ACTION_TYPE_FABULOUS)
	// 更新帖子点赞总计 +1
	if err := svc.post.UpdatePostLikeNum(postId, now, consts.CONFIRM_OPERATE, score); err != nil {
		log.Log.Errorf("like_trace: update post like num err:%s", err)
		svc.engine.Rollback()
		return errdef.LIKE_POST_FAIL
	}

	// 未点赞
	// 记录存在 且 状态为 未点赞 更新状态为 已点赞
	if info != nil && info.Status == consts.NOT_GIVE_LIKE {
		info.Status = consts.ALREADY_GIVE_LIKE
		info.CreateAt = now
		if err := svc.like.UpdateLikeStatus(); err != nil {
			log.Log.Errorf("like_trace: update like status err:%s", err)
			svc.engine.Rollback()
			return errdef.LIKE_POST_FAIL
		}

	} else {
		// 添加点赞记录
		if err := svc.like.AddGiveLikeByType(userId, post.UserId, postId, consts.ALREADY_GIVE_LIKE, consts.TYPE_POSTS); err != nil {
			log.Log.Errorf("like_trace: add like post record err:%s", err)
			svc.engine.Rollback()
			return errdef.LIKE_POST_FAIL
		}
	}

	svc.engine.Commit()

	// todo: 帖子点赞 推送内容
	// 发送帖子点赞推送
	redismq.PushEventMsg(redismq.NewEvent(post.UserId, fmt.Sprint(post.Id), user.NickName, "", "", consts.POST_LIKE_MSG))
	// 帖子置顶事件
	redismq.PushTopEventMsg(redismq.NewTopEvent(post.UserId, fmt.Sprint(post.Id), consts.EVENT_SET_TOP_POST))

	return errdef.SUCCESS
}

// 取消点赞（帖子）
func (svc *LikeModule) CancelLikeForPost(userId string, postId int64) int {
	// 开启事务
	if err := svc.engine.Begin(); err != nil {
		log.Log.Errorf("like_trace: session begin err:%s", err)
		return errdef.ERROR
	}

	// 查询用户是否存在
	if user := svc.user.FindUserByUserid(userId); user == nil {
		log.Log.Errorf("like_trace: user not found, userId:%s", userId)
		svc.engine.Rollback()
		return errdef.USER_NOT_EXISTS
	}

	// 查找帖子是否存在
	post, err := svc.post.GetPostById(fmt.Sprint(postId))
	if err != nil || post == nil {
		log.Log.Errorf("like_trace: cancel like post not found, postId:%d", postId)
		svc.engine.Rollback()
		return errdef.LIKE_POST_NOT_EXISTS
	}

	// 获取点赞的信息 判断是否已点赞 记录不存在 则 未点过赞
	info := svc.like.GetLikeInfo(userId, postId, consts.TYPE_POSTS)
	if info == nil {
		log.Log.Errorf("like_trace: record not found, not give like, userId:%s, postId:%d", userId, postId)
		svc.engine.Rollback()
		return errdef.LIKE_RECORD_NOT_EXISTS
	}

	// 状态 ！= 已点赞 提示重复操作
	if info.Status != consts.ALREADY_GIVE_LIKE {
		log.Log.Errorf("like_trace: already cancel like, userId:%s, postId:%d", userId, postId)
		svc.engine.Rollback()
		return errdef.LIKE_REPEAT_CANCEL
	}

	now := int(time.Now().Unix())
	score := svc.config.GetActionScore(int(consts.WORK_TYPE_POST), consts.ACTION_TYPE_FABULOUS)
	// 更新帖子点赞总计 -1
	if err := svc.post.UpdatePostLikeNum(postId, now, consts.CANCEL_OPERATE, score*-1); err != nil {
		log.Log.Errorf("like_trace: update post like num err:%s", err)
		svc.engine.Rollback()
		return errdef.LIKE_CANCEL_FAIL
	}

	info.Status = consts.NOT_GIVE_LIKE
	info.CreateAt = now
	// 更新状态 未点赞
	if err := svc.like.UpdateLikeStatus(); err != nil {
		log.Log.Errorf("like_trace: update like status err:%s", err)
		svc.engine.Rollback()
		return errdef.LIKE_CANCEL_FAIL
	}

	svc.engine.Commit()

	return errdef.SUCCESS
}

// 点赞帖子评论
func (svc *LikeModule) GiveLikeForPostComment(userId string, commentId int64) int {
	// 开启事务
	if err := svc.engine.Begin(); err != nil {
		log.Log.Errorf("like_trace: session begin err:%s", err)
		return errdef.ERROR
	}

	// 查询用户是否存在
	user := svc.user.FindUserByUserid(userId)
	if user == nil {
		log.Log.Errorf("like_trace: user not found, userId:%s", userId)
		svc.engine.Rollback()
		return errdef.USER_NOT_EXISTS
	}

	// 查找帖子评论是否存在
	comment := svc.comment.GetPostCommentById(fmt.Sprint(commentId))
	if comment == nil {
		log.Log.Errorf("like_trace: like comment not found, commentId:%d", commentId)
		svc.engine.Rollback()
		return errdef.LIKE_COMMENT_NOT_EXISTS
	}

	if code := svc.GiveLike(userId, comment.UserId, commentId, consts.TYPE_POST_COMMENT); code != errdef.SUCCESS {
		svc.engine.Rollback()
		return code
	}

	svc.engine.Commit()

	// 发送帖子评论点赞推送
	redismq.PushEventMsg(redismq.NewEvent(comment.UserId, fmt.Sprint(comment.PostId), user.NickName, "", comment.Content, consts.POST_COMMENT_LIKE_MSG))

	return errdef.SUCCESS
}

// 取消点赞（帖子评论）
func (svc *LikeModule) CancelLikeForPostComment(userId string, commentId int64) int {
	// 开启事务
	if err := svc.engine.Begin(); err != nil {
		log.Log.Errorf("like_trace: session begin err:%s", err)
		return errdef.ERROR
	}

	// 查询用户是否存在
	if user := svc.user.FindUserByUserid(userId); user == nil {
		log.Log.Errorf("like_trace: user not found, userId:%s", userId)
		svc.engine.Rollback()
		return errdef.USER_NOT_EXISTS
	}

	// 查找评论是否存在
	if comment := svc.comment.GetPostCommentById(fmt.Sprint(commentId)); comment == nil {
		log.Log.Errorf("like_trace: cancel like comment not found, commentId:%d", commentId)
		svc.engine.Rollback()
		return errdef.LIKE_COMMENT_NOT_EXISTS
	}

	if code := svc.CancelLike(userId, commentId, consts.TYPE_POST_COMMENT); code != errdef.SUCCESS {
		svc.engine.Rollback()
		return code
	}

	svc.engine.Commit()

	return errdef.SUCCESS
}

// 取消点赞（资讯评论）
func (svc *LikeModule) CancelLikeForInformationComment(userId string, commentId int64) int {
	// 开启事务
	if err := svc.engine.Begin(); err != nil {
		log.Log.Errorf("like_trace: session begin err:%s", err)
		return errdef.ERROR
	}

	// 查询用户是否存在
	if user := svc.user.FindUserByUserid(userId); user == nil {
		log.Log.Errorf("like_trace: user not found, userId:%s", userId)
		svc.engine.Rollback()
		return errdef.USER_NOT_EXISTS
	}

	// 查找评论是否存在
	if comment := svc.comment.GetInformationCommentById(fmt.Sprint(commentId)); comment == nil {
		log.Log.Errorf("like_trace: cancel like comment not found, commentId:%d", commentId)
		svc.engine.Rollback()
		return errdef.LIKE_COMMENT_NOT_EXISTS
	}

	if code := svc.CancelLike(userId, commentId, consts.TYPE_INFORMATION_COMMENT); code != errdef.SUCCESS {
		svc.engine.Rollback()
		return code
	}

	svc.engine.Commit()

	return errdef.SUCCESS
}

// 点赞资讯
func (svc *LikeModule) GiveLikeForInformation(userId string, newsId int64) int {
	// 开启事务
	if err := svc.engine.Begin(); err != nil {
		log.Log.Errorf("like_trace: session begin err:%s", err)
		return errdef.ERROR
	}

	// 查询用户是否存在
	user := svc.user.FindUserByUserid(userId)
	if user == nil {
		log.Log.Errorf("like_trace: user not found, userId:%s", userId)
		svc.engine.Rollback()
		return errdef.USER_NOT_EXISTS
	}

	// 查找资讯是否存在
	ok, err := svc.information.GetInformationById(fmt.Sprint(newsId))
	if !ok || err != nil {
		svc.engine.Rollback()
		return errdef.LIKE_INFORMATION_NOT_EXISTS
	}

	// 获取点赞的资讯信息
	info := svc.like.GetLikeInfo(userId, newsId, consts.LIKE_TYPE_INFORMATION)
	// 是否已点赞
	// 已点赞
	if info != nil && info.Status == consts.ALREADY_GIVE_LIKE {
		log.Log.Errorf("like_trace: already give like, userId:%s, postId:%d", userId, newsId)
		svc.engine.Rollback()
		return errdef.LIKE_ALREADY_EXISTS
	}

	now := int(time.Now().Unix())
	score := svc.config.GetActionScore(int(consts.WORK_TYPE_INFO), consts.ACTION_TYPE_FABULOUS)
	// 更新资讯点赞总计 +1
	if err := svc.information.UpdateInformationLikeNum(newsId, now, consts.CONFIRM_OPERATE, score); err != nil {
		log.Log.Errorf("like_trace: update information like num err:%s", err)
		svc.engine.Rollback()
		return errdef.LIKE_INFORMATION_FAIL
	}

	// 未点赞
	// 记录存在 且 状态为 未点赞 更新状态为 已点赞
	if info != nil && info.Status == consts.NOT_GIVE_LIKE {
		info.Status = consts.ALREADY_GIVE_LIKE
		info.CreateAt = now
		if err := svc.like.UpdateLikeStatus(); err != nil {
			log.Log.Errorf("like_trace: update like status err:%s", err)
			svc.engine.Rollback()
			return errdef.LIKE_INFORMATION_FAIL
		}

	} else {
		// 添加点赞记录
		if err := svc.like.AddGiveLikeByType(userId, svc.information.Information.UserId, newsId, consts.ALREADY_GIVE_LIKE, consts.LIKE_TYPE_INFORMATION); err != nil {
			log.Log.Errorf("like_trace: add like information record err:%s", err)
			svc.engine.Rollback()
			return errdef.LIKE_INFORMATION_FAIL
		}
	}

	svc.engine.Commit()

	// 资讯置顶事件
	redismq.PushTopEventMsg(redismq.NewTopEvent(svc.information.Information.UserId,
		fmt.Sprint(svc.information.Information.Id), consts.EVENT_SET_TOP_INFO))

	return errdef.SUCCESS
}

// 取消点赞（资讯）
func (svc *LikeModule) CancelLikeForInformation(userId string, newsId int64) int {
	// 开启事务
	if err := svc.engine.Begin(); err != nil {
		log.Log.Errorf("like_trace: session begin err:%s", err)
		return errdef.ERROR
	}

	// 查询用户是否存在
	if user := svc.user.FindUserByUserid(userId); user == nil {
		log.Log.Errorf("like_trace: user not found, userId:%s", userId)
		svc.engine.Rollback()
		return errdef.USER_NOT_EXISTS
	}

	// 查找帖子是否存在
	ok, err := svc.information.GetInformationById(fmt.Sprint(newsId))
	if !ok || err != nil {
		log.Log.Errorf("like_trace: cancel like information not found, newsId:%d", newsId)
		svc.engine.Rollback()
		return errdef.LIKE_INFORMATION_NOT_EXISTS
	}

	// 获取点赞的信息 判断是否已点赞 记录不存在 则 未点过赞
	info := svc.like.GetLikeInfo(userId, newsId, consts.LIKE_TYPE_INFORMATION)
	if info == nil {
		log.Log.Errorf("like_trace: record not found, not give like, userId:%s, postId:%d", userId, newsId)
		svc.engine.Rollback()
		return errdef.LIKE_RECORD_NOT_EXISTS
	}

	// 状态 ！= 已点赞 提示重复操作
	if info.Status != consts.ALREADY_GIVE_LIKE {
		log.Log.Errorf("like_trace: already cancel like, userId:%s, postId:%d", userId, newsId)
		svc.engine.Rollback()
		return errdef.LIKE_REPEAT_CANCEL
	}

	now := int(time.Now().Unix())
	score := svc.config.GetActionScore(int(consts.WORK_TYPE_INFO), consts.ACTION_TYPE_FABULOUS)
	// 更新资讯点赞总计 -1
	if err := svc.information.UpdateInformationLikeNum(newsId, now, consts.CANCEL_OPERATE, score*-1); err != nil {
		log.Log.Errorf("like_trace: update information like num err:%s", err)
		svc.engine.Rollback()
		return errdef.LIKE_CANCEL_FAIL
	}

	info.Status = consts.NOT_GIVE_LIKE
	info.CreateAt = now
	// 更新状态 未点赞
	if err := svc.like.UpdateLikeStatus(); err != nil {
		log.Log.Errorf("like_trace: update like status err:%s", err)
		svc.engine.Rollback()
		return errdef.LIKE_CANCEL_FAIL
	}

	svc.engine.Commit()

	return errdef.SUCCESS
}
