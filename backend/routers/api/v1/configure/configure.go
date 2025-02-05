package configure

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sports_service/backend/controller/configure"
	"sports_service/global/backend/errdef"
	"sports_service/global/consts"
	"sports_service/models"
	"sports_service/models/mbanner"
	"sports_service/models/mconfigure"
	"sports_service/models/muser"
	"sports_service/models/mvideo"
	"sports_service/tools/tencentCloud"
	"sports_service/util"
	"strconv"
)

// 添加banner
func AddBanner(c *gin.Context) {
	reply := errdef.New(c)
	params := new(mbanner.AddBannerParams)
	if err := c.BindJSON(params); err != nil {
		reply.Response(http.StatusBadRequest, errdef.INVALID_PARAMS)
		return
	}

	svc := configure.New(c)
	syscode := svc.AddBanner(params)
	reply.Response(http.StatusOK, syscode)
}

// 更新banner
func UpdateBanner(c *gin.Context) {
	reply := errdef.New(c)
	params := new(mbanner.UpdateBannerParams)
	if err := c.BindJSON(params); err != nil {
		reply.Response(http.StatusBadRequest, errdef.INVALID_PARAMS)
		return
	}

	svc := configure.New(c)
	syscode := svc.UpdateBanner(params)
	reply.Response(http.StatusOK, syscode)
}

// 删除banner
func DelBanner(c *gin.Context) {
	reply := errdef.New(c)
	param := new(mbanner.DelBannerParam)
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusBadRequest, errdef.INVALID_PARAMS)
		return
	}

	svc := configure.New(c)
	syscode := svc.DelBanner(param)
	reply.Response(http.StatusOK, syscode)
}

// 获取banner列表
func GetBanners(c *gin.Context) {
	reply := errdef.New(c)
	page, size := util.PageInfo(c.Query("page"), c.Query("size"))

	svc := configure.New(c)
	list := svc.GetBannerList(page, size)

	reply.Data["list"] = list
	reply.Data["total"] = svc.GetBannerTotal()
	reply.Response(http.StatusOK, errdef.SUCCESS)
}

// 添加系统头像
func AddAvatar(c *gin.Context) {
	reply := errdef.New(c)
	params := new(muser.AddSystemAvatarParams)
	if err := c.BindJSON(params); err != nil {
		reply.Response(http.StatusBadRequest, errdef.INVALID_PARAMS)
		return
	}

	svc := configure.New(c)
	syscode := svc.AddSystemAvatar(params)
	reply.Response(http.StatusOK, syscode)
}

// 删除系统头像
func DelAvatar(c *gin.Context) {
	reply := errdef.New(c)
	param := new(muser.DelSystemAvatarParam)
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusBadRequest, errdef.INVALID_PARAMS)
		return
	}

	svc := configure.New(c)
	syscode := svc.DelSystemAvatar(param)
	reply.Response(http.StatusOK, syscode)
}

// 获取系统头像列表列表
func GetAvatarList(c *gin.Context) {
	reply := errdef.New(c)

	svc := configure.New(c)
	list := svc.GetSystemAvatars()
	reply.Data["list"] = list
	reply.Response(http.StatusOK, errdef.SUCCESS)
}

// 热搜配置
func GetHotSearch(c *gin.Context) {
	reply := errdef.New(c)

	svc := configure.New(c)
	list := svc.GetHotSearch()
	reply.Data["list"] = list
	reply.Response(http.StatusOK, errdef.SUCCESS)
}

// 添加热搜
func AddHotSearch(c *gin.Context) {
	reply := errdef.New(c)
	params := new(mvideo.AddHotSearchParams)
	if err := c.BindJSON(params); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := configure.New(c)
	syscode := svc.AddHotSearch(params)
	reply.Response(http.StatusOK, syscode)
}

// 删除热搜
func DelHotSearch(c *gin.Context) {
	reply := errdef.New(c)
	param := new(mvideo.DelHotSearchParams)
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := configure.New(c)
	syscode := svc.DelHotSearch(param)
	reply.Response(http.StatusOK, syscode)
}

// 设置热搜权重
func SetSortByHotSearch(c *gin.Context) {
	reply := errdef.New(c)
	param := new(mvideo.SetSortParams)
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := configure.New(c)
	syscode := svc.SetSortByHotSearch(param)
	reply.Response(http.StatusOK, syscode)
}

// 设置热搜状态
func SetStatusByHotSearch(c *gin.Context) {
	reply := errdef.New(c)
	param := new(mvideo.SetStatusParams)
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := configure.New(c)
	syscode := svc.SetStatusByHotSearch(param)
	reply.Response(http.StatusOK, syscode)
}

// 获取腾讯cos临时通行证
func CosTempAccess(c *gin.Context) {
	reply := errdef.New(c)
	// private 私读私写 / public 公读私写
	uploadType := c.DefaultQuery("upload_type", "private")
	client := tencentCloud.New(consts.TX_CLOUD_COS_SECRET_ID, consts.TX_CLOUD_COS_SECRET_KEY, consts.TMS_API_DOMAIN)
	info, err := client.GetCosTempAccess("ap-shanghai", uploadType)
	if err != nil {
		reply.Response(http.StatusOK, errdef.CONFIG_COS_ACCESS_FAIL)
		return
	}

	reply.Data["access_info"] = info
	reply.Response(http.StatusOK, errdef.SUCCESS)
}

// 添加新包
func AddPackage(c *gin.Context) {
	reply := errdef.New(c)
	param := &mconfigure.AddPackageParams{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := configure.New(c)
	syscode := svc.AddNewPackage(param)
	reply.Response(http.StatusOK, syscode)
}

// 更新包数据
func UpdatePackage(c *gin.Context) {
	reply := errdef.New(c)
	param := &mconfigure.UpdatePackageParams{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := configure.New(c)
	syscode := svc.UpdatePackageInfo(param)
	reply.Response(http.StatusOK, syscode)
}

// 删除更新包
func DelPackage(c *gin.Context) {
	reply := errdef.New(c)
	param := &mconfigure.DelPackageParam{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := configure.New(c)
	syscode := svc.DelPackage(param.Id)
	reply.Response(http.StatusOK, syscode)
}

// 包列表
func PackageList(c *gin.Context) {
	reply := errdef.New(c)
	page, size := util.PageInfo(c.Query("page"), c.Query("size"))
	svc := configure.New(c)
	list := svc.GetPackageList(page, size)
	reply.Data["list"] = list
	reply.Response(http.StatusOK, errdef.SUCCESS)
}

// 包详情
func PackageDetail(c *gin.Context) {
	reply := errdef.New(c)
	id := c.Query("id")

	svc := configure.New(c)
	detail := svc.GetPackageDetail(id)
	reply.Data["detail"] = detail
	reply.Response(http.StatusOK, errdef.SUCCESS)
}

func AddActionScoreConf(c *gin.Context) {
	reply := errdef.New(c)
	param := &models.ActionScoreConfig{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := configure.New(c)
	reply.Response(http.StatusOK, svc.AddActionScoreConf(param))
}

func UpdateActionScoreConf(c *gin.Context) {
	reply := errdef.New(c)
	param := &models.ActionScoreConfig{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := configure.New(c)
	reply.Response(http.StatusOK, svc.UpdateActionScoreConf(param))
}

func GetActionScoreConf(c *gin.Context) {
	reply := errdef.New(c)
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := configure.New(c)
	code, detail := svc.GetActionScoreConfById(int64(id))
	reply.Data["detail"] = detail
	reply.Response(http.StatusOK, code)
}

func GetActionScoreConfList(c *gin.Context) {
	reply := errdef.New(c)
	svc := configure.New(c)
	code, list := svc.GetActionScoreConfList()
	reply.Data["list"] = list
	reply.Response(http.StatusOK, code)
}
