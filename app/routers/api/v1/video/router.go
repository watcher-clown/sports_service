package video

import (
	"github.com/gin-gonic/gin"
	"sports_service/middleware/sign"
	"sports_service/middleware/token"
)

// 视频点播模块路由
func Router(engine *gin.Engine) {
	api := engine.Group("/api/v1")
	video := api.Group("/video")
	video.Use(sign.CheckSign())
	{
		// 用户发布视频
		video.POST("/publish", token.TokenAuth(), VideoPublish)
		// 用户视频浏览记录
		video.GET("/browse/history", token.TokenAuth(), BrowseHistory)
		// 用户发布的视频列表
		video.GET("/publish/list", VideoPublishList)
		// 其他用户发布的视频列表
		video.GET("/other/publish", OtherUserPublishList)
		// 删除浏览记录
		video.POST("/delete/history", token.TokenAuth(), DeleteHistory)
		// 删除发布的记录
		video.POST("/delete/publish", token.TokenAuth(), DeletePublish)
		// 首页推荐的视频列表
		video.GET("/recommend", RecommendVideos)
		// 首页推荐的banner列表
		video.GET("/homepage/banner", RecommendBanners)
		// 关注的人发布的视频列表
		video.GET("/attention", token.TokenAuth(), AttentionVideos)
		// 视频详情信息
		video.GET("/detail", VideoDetail)
		// 视频详情页推荐视频（同标签推荐）
		video.GET("/detail/recommend", DetailRecommend)
		// 获取上传签名（腾讯云）
		video.GET("/upload/sign", token.TokenAuth(), UploadSign)
		// 事件回调
		video.GET("/event/callback", EventCallback)
		// 用户自定义视频标签检测
		video.POST("/custom/labels", CheckCustomLabels)
		// 获取视频标签列表
		video.GET("/label/list", VideoLabelList)
		// 举报视频
		video.POST("/report", VideoReport)
		// 上传测试
		//video.GET("/test/upload", TestUpload)
		// 记录用户视频播放的时长
		video.POST("/record/play/duration", RecordPlayDuration)
		// 获取视频分区配置列表
		video.GET("/subarea", VideoSubarea)
		// 创建视频专辑
		video.POST("/create/album", token.TokenAuth(), CreateVideoAlbum)
		// 将视频添加到专辑中
		//video.POST("/add/album", token.TokenAuth(), AddVideoToAlbum)
		// 视频分区列表
		video.GET("/subarea/list", VideoListBySubarea)
		// 获取用户创建的视频专辑列表
		video.GET("/album/list", token.TokenAuth(), VideoAlbumList)
		// 首页板块信息
		video.GET("/homepage/section/info", HomePageSectionInfo)
		// 首页板块推荐的内容
		video.GET("/homepage/section/recommend", SectionRecommendInfo)
	}
}
