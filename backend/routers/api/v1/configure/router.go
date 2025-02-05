package configure

import (
	"github.com/gin-gonic/gin"
	"sports_service/middleware/jwt"
)

func Router(engine *gin.Engine) {
	api := engine.Group("/backend/v1")
	configure := api.Group("/config")
	configure.Use(jwt.JwtAuth())
	{
		// 添加banner
		configure.POST("/add/banner", AddBanner)
		// 更新banner
		configure.POST("/update/banner", UpdateBanner)
		// 删除banner
		configure.POST("/del/banner", DelBanner)
		// 获取banner列表
		configure.GET("/banners", GetBanners)
		// 添加系统头像
		configure.POST("/add/avatar", AddAvatar)
		// 删除系统头像
		configure.POST("/del/avatar", DelAvatar)
		// 获取系统头像列表
		configure.GET("/avatar/list", GetAvatarList)
		// 热搜列表
		configure.GET("/hot/search", GetHotSearch)
		// 添加热搜
		configure.POST("/add/hot/search", AddHotSearch)
		// 删除热搜
		configure.POST("/del/hot/search", DelHotSearch)
		// 设置热搜权重
		configure.POST("/set/hot/sort", SetSortByHotSearch)
		// 设置热搜状态
		configure.POST("/set/hot/status", SetStatusByHotSearch)
		// 获取腾讯cos通行证
		configure.GET("/cos/access", CosTempAccess)
		// 添加新包
		configure.POST("/package/add", AddPackage)
		// 更新包信息
		configure.POST("/package/update", UpdatePackage)
		// 删除包
		configure.POST("/package/del", DelPackage)
		// 获取包信息列表
		configure.GET("/package/list", PackageList)
		// 获取包详情信息
		configure.GET("/package/detail", PackageDetail)
		// 添加行为得分配置
		configure.POST("/add/action/score", AddActionScoreConf)
		// 更新行为得分配置
		configure.POST("/update/action/score", UpdateActionScoreConf)
		// 获取行为得分配置
		configure.GET("/action/score", GetActionScoreConf)
		// 获取行为得分配置列表
		configure.GET("/action/score/list", GetActionScoreConfList)
	}
}
