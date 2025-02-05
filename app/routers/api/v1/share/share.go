package share

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sports_service/app/controller/cshare"
	"sports_service/global/app/errdef"
	"sports_service/global/app/log"
	"sports_service/global/consts"
	"sports_service/models/mshare"
)

// @Summary 分享 (ok)
// @Tags 分享模块
// @Version 1.0
// @Description
// @Accept json
// @Produce  json
// @Param   AppId         header    string 	true  "AppId"
// @Param   Secret        header    string 	true  "调用/api/v1/client/init接口 服务端下发的secret"
// @Param   Timestamp     header    string 	true  "请求时间戳 单位：秒"
// @Param   Sign          header    string 	true  "签名 md5签名32位值"
// @Param   Version 	  header    string 	true  "版本" default(1.0.0)
// @Param   ForwardData   body  body string true "请求参数"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"success","tm":"1588888888"}"
// @Failure 500 {string} json "{"code":500,"data":{},"msg":"fail","tm":"1588888888"}"
// @Router /api/v1/share/social [post]
// 分享/转发 到社交平台
func ShareWithSocialPlatform(c *gin.Context) {
	reply := errdef.New(c)
	params := new(mshare.ShareParams)
	if err := c.BindJSON(params); err != nil {
		log.Log.Errorf("share_trace: share params err:%s, params:%+v", err, params)
		reply.Response(http.StatusBadRequest, errdef.INVALID_PARAMS)
		return
	}

	svc := cshare.New(c)
	code := svc.ShareData(params)
	reply.Response(http.StatusOK, code)
}

// 分享/转发 到社区
func ShareWithCommunity(c *gin.Context) {
	reply := errdef.New(c)
	userId, _ := c.Get(consts.USER_ID)
	params := new(mshare.ShareParams)
	if err := c.BindJSON(params); err != nil {
		log.Log.Errorf("share_trace: share params err:%s, params:%+v", err, params)
		reply.Response(http.StatusBadRequest, errdef.INVALID_PARAMS)
		return
	}

	svc := cshare.New(c)
	params.UserId = userId.(string)
	code := svc.ShareData(params)
	reply.Response(http.StatusOK, code)
}

// 获取分享链接
func GetShareUrl(c *gin.Context) {
	reply := errdef.New(c)
	userId := c.Query("user_id")
	shareType := c.Query("share_type")
	contentType := c.Query("content_type")
	contentId := c.Query("content_id")
	svc := cshare.New(c)
	shareUrl := svc.GenShareUrl(userId, contentType, contentId, shareType)
	reply.Data["share_url"] = shareUrl
	reply.Response(http.StatusOK, errdef.SUCCESS)
}
