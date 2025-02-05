package post

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sports_service/backend/controller/cpost"
	"sports_service/global/app/log"
	"sports_service/global/backend/errdef"
	"sports_service/models/mcommunity"
	"sports_service/models/mposting"
	"sports_service/util"
)

// 帖子审核
func AuditPost(c *gin.Context) {
	reply := errdef.New(c)
	param := &mposting.AudiPostParam{}
	if err := c.BindJSON(param); err != nil {
		log.Log.Errorf("post_trace: invalid param, err:%s", err)
		reply.Response(http.StatusBadRequest, errdef.INVALID_PARAMS)
		return
	}

	svc := cpost.New(c)
	syscode := svc.AudiPost(param)
	reply.Response(http.StatusOK, syscode)
}

// 帖子列表 todo：展示数据待确认
func PostList(c *gin.Context) {
	reply := errdef.New(c)
	page, size := util.PageInfo(c.Query("page"), c.Query("size"))
	status := c.DefaultQuery("status", "1")
	title := c.Query("title")
	svc := cpost.New(c)
	code, list := svc.GetPostList(page, size, status, title)
	reply.Data["list"] = list
	reply.Data["total"] = svc.GetTotalCountByPost(status, title)
	reply.Response(http.StatusOK, code)
}

func AddSection(c *gin.Context) {
	reply := errdef.New(c)
	param := &mcommunity.AddSection{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := cpost.New(c)
	reply.Response(http.StatusOK, svc.AddSection(param))
}

func DelSection(c *gin.Context) {
	reply := errdef.New(c)
	param := &mcommunity.DelSection{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := cpost.New(c)
	reply.Response(http.StatusOK, svc.DelSection(param))
}

func EditSection(c *gin.Context) {
	reply := errdef.New(c)
	param := &mcommunity.AddSection{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := cpost.New(c)
	reply.Response(http.StatusOK, svc.EditSection(param))
}

func AddTopic(c *gin.Context) {
	reply := errdef.New(c)
	param := &mcommunity.AddTopic{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := cpost.New(c)
	reply.Response(http.StatusOK, svc.AddTopic(param))
}

func DelTopic(c *gin.Context) {
	reply := errdef.New(c)
	param := &mcommunity.DelTopic{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := cpost.New(c)
	reply.Response(http.StatusOK, svc.DelTopic(param))
}

func EditTopic(c *gin.Context) {
	reply := errdef.New(c)
	param := &mcommunity.AddTopic{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := cpost.New(c)
	reply.Response(http.StatusOK, svc.UpdateTopic(param))
}

func PostSetting(c *gin.Context) {
	reply := errdef.New(c)
	param := &mposting.SettingParam{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := cpost.New(c)
	reply.Response(http.StatusOK, svc.PostSetting(param))
}

func ApplyCreamList(c *gin.Context) {
	page, size := util.PageInfo(c.Query("page"), c.Query("size"))
	reply := errdef.New(c)
	svc := cpost.New(c)
	code, list := svc.GetApplyCreamList(page, size)
	reply.Data["list"] = list
	reply.Data["total"] = svc.GetApplyCreamCount()
	reply.Response(http.StatusOK, code)
}

func SectionList(c *gin.Context) {
	reply := errdef.New(c)
	svc := cpost.New(c)
	code, list := svc.GetSectionList()
	reply.Data["list"] = list
	reply.Response(http.StatusOK, code)
}

func TopicList(c *gin.Context) {
	reply := errdef.New(c)
	svc := cpost.New(c)
	code, list := svc.GetTopicList()
	reply.Data["list"] = list
	reply.Response(http.StatusOK, code)
}

func BatchEditPostInfo(c *gin.Context) {
	reply := errdef.New(c)
	param := &mposting.BatchEditParam{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := cpost.New(c)
	reply.Response(http.StatusOK, svc.BatchEditPostInfo(param))
}

func DelPost(c *gin.Context) {
	reply := errdef.New(c)
	postId := c.Query("id")
	svc := cpost.New(c)
	reply.Response(http.StatusOK, svc.DelPost(postId))
}
