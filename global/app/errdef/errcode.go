package errdef

const (
	PAY_SUCCESS        = 1
	CALLBACK_SUCCESS   = 0
	SUCCESS            = 200
	ERROR              = 500
	INVALID_PARAMS     = 400
	UNAUTHORIZED       = 401

	INVALID_TOKEN      = 888

	// 用户相关错误码 1001-2000
	USER_FREE_LOGIN_FAIL    = 1001
	USER_INVALID_MOBILE_NUM = 1002
	USER_ALREADY_EXISTS     = 1003
	USER_REPEAT_REG         = 1004
	USER_REGISTER_FAIL      = 1005
	USER_ADD_INFO_FAIL      = 1006
	USER_GET_INFO_FAIL      = 1007
	USER_NOT_EXISTS         = 1008
	USER_INVALID_NAME_LEN   = 1009
	USER_INVALID_SIGN_LEN   = 1010
	USER_NICK_NAME_EXISTS   = 1011
	USER_AVATAR_NOT_EXISTS  = 1012
	USER_COUNTRY_NOT_EXISTS = 1013
	USER_UPDATE_INFO_FAIL   = 1014
	USER_INVALID_NAME       = 1015
	USER_INVALID_SIGNATURE  = 1016
	USER_FEEDBACK_FAIL      = 1017
	USER_ADD_NOTIFY_SET_FAIL= 1018
	USER_NO_LOGIN           = 1019
	USER_INVALID_FEEDBACK   = 1020
	USER_BIND_DEVICE_TOKEN  = 1021
	USER_PACKAGE_NOT_EXISTS = 1022
	USER_LATEST_PACKAGE_FAIL= 1023
	USER_FORBID_STATUS      = 1024
	USER_GEN_IM_SIGN_FAIL   = 1025
	USER_UPDATE_IM_SIGN_FAIL= 1026
	USER_ADD_GUEST_FAIL     = 1027
	USER_GET_GUEST_SIGN_FAIL= 1028

	WX_USER_INFO_FAIL       = 1101
	WX_ACCESS_TOKEN_FAIL    = 1102
	WX_REGISTER_FAIL        = 1103
	WX_ADD_ACCOUNT_FAIL     = 1104

	WEIBO_USER_INFO_FAIL    = 1201
	WEIBO_ADD_ACCOUNT_FAIL  = 1202
	WEIBO_REGISTER_FAIL     = 1203

	QQ_UNIONID_FAIL         = 1301
	QQ_USER_INFO_FAIL       = 1302
	QQ_REGISTER_FAIL        = 1303
	QQ_ADD_ACCOUNT_FAIL     = 1304

	// 关注相关错误码 2001-3000
	ATTENTION_USER_NOT_EXISTS   = 2001
	ATTENTION_ALREADY_EXISTS    = 2002
	ATTENTION_USER_FAIL         = 2003
	ATTENTION_RECORD_NOT_EXISTS = 2004
	ATTENTION_REPEAT_CANCEL     = 2005
	ATTENTION_CANCEL_FAIL       = 2006
	ATTENTION_YOURSELF_FAIL     = 2007

	// 收藏相关错误码 3001-4000
	COLLECT_VIDEO_NOT_EXISTS    = 3001
	COLLECT_ALREADY_EXISTS      = 3002
	COLLECT_VIDEO_FAIL          = 3003
	COLLECT_RECORD_NOT_EXISTS   = 3004
	COLLECT_REPEAT_CANCEL       = 3005
	COLLECT_CANCEL_FAIL         = 3006
	COLLECT_DELETE_FAIL         = 3007

	// 视频相关错误码 4001-5000
	VIDEO_NOT_EXISTS            = 4001
	VIDEO_PUBLISH_FAIL          = 4002
	VIDEO_DELETE_HISTORY_FAIL   = 4003
	VIDEO_DELETE_PUBLISH_FAIL   = 4004
	VIDEO_DELETE_LABEL_FAIL     = 4005
	VIDEO_DELETE_STATISTIC_FAIL = 4006
	VIDEO_UPLOAD_GEN_SIGN_FAIL  = 4007
	VIDEO_INVALID_DESCRIBE      = 4008
	VIDEO_INVALID_TITLE         = 4009
	VIDEO_INVALID_CUSTOM_LABEL  = 4010
	VIDEO_REPORT_FAIL           = 4011
	VIDEO_INVALID_PLAY_DURATION = 4012
	VIDEO_RECORD_PLAY_DURATION  = 4013
	VIDEO_SUBAREA_FAIL          = 4014
	VIDEO_CREATE_ALBUM_FAIL     = 4015
	VIDEO_INVALID_ALBUM_NAME    = 4016
	VIDEO_ALBUM_NOT_EXISTS      = 4017
	VIDEO_ADD_TO_ALBUM_FAIL     = 4018
	VIDEO_LIST_BY_SUBAREA_FAIL  = 4019

	// 点赞相关错误码 5001-6000
	LIKE_VIDEO_NOT_EXISTS       = 5001
	LIKE_ALREADY_EXISTS         = 5002
	LIKE_VIDEO_FAIL             = 5003
	LIKE_RECORD_NOT_EXISTS      = 5004
	LIKE_REPEAT_CANCEL          = 5005
	LIKE_CANCEL_FAIL            = 5006
	LIKE_COMMENT_NOT_EXISTS     = 5007
	LIKE_COMMENT_FAIL           = 5008
	LIKE_POST_NOT_EXISTS        = 5009
	LIKE_POST_FAIL              = 5010
	LIKE_INFORMATION_NOT_EXISTS = 5011
	LIKE_INFORMATION_FAIL       = 5012

	// 通知相关错误码 6001-7000
	NOTIFY_SETTING_FAIL         = 6001

	// 评论相关错误码 7001-8000
	COMMENT_PUBLISH_FAIL        = 7001
	COMMENT_NOT_FOUND           = 7002
	COMMENT_REPLY_FAIL          = 7003
	COMMENT_REPLY_NOT_FOUND     = 7004
	COMMENT_INVALID_LEN         = 7005
	COMMENT_INVALID_CONTENT     = 7006
	COMMENT_INVALID_REPLY       = 7007
	COMMENT_REPORT_FAIL         = 7008
	COMMENT_DELETE_FAIL         = 7009
	COMMENT_USER_NOT_MATCH      = 7010

	// 短信验证码相关错误 8001-9000
	SMS_CODE_INTERVAL_ERROR     = 8001
	SMS_CODE_INTERVAL_SHORT     = 8002
	SMS_CODE_INVALID_SEND_TYPE  = 8003
	SMS_CODE_SEND_FAIL          = 8004
	SMS_INVALID_CODE            = 8005
	SMS_CODE_NOT_SEND           = 8006
	SMS_CODE_NOT_MATCH          = 8007

	// 弹幕相关错误码 9001 - 10000
	BARRAGE_VIDEO_SEND_FAIL     = 9001
	BARRAGE_INVALID_CONTENT     = 9002
	BARRAGE_VIDEO_LIST_FAIL     = 9003
	BARRAGE_LIVE_SEND_FAIL      = 9004

	// 搜索相关错误码 10001 - 11000
	SEARCH_CLEAN_HISTORY_FAIL   = 10001

	// 腾讯云相关错误码 11001 - 12000
	CLOUD_COS_ACCESS_FAIL       = 11001
	CLOUD_FILTER_FAIL           = 11002

	// 帖子相关错误码 12001 - 13000
	POST_INVALID_TITLE          = 12001
	POST_INVALID_CONTENT        = 12002
	POST_SECTION_NOT_EXISTS     = 12003
	POST_TOPIC_NOT_EXISTS       = 12004
	POST_PUBLISH_FAIL           = 12005
	POST_INVALID_CONTENT_LEN    = 12006
	POST_DETAIL_FAIL            = 12007
	POST_NOT_EXISTS             = 12008
	POST_DELETE_PUBLISH_FAIL    = 12009
	POST_DELETE_TOPIC_FAIL      = 12010
	POST_DELETE_STATISTIC_FAIL  = 12011
	POST_AUTHOR_NOT_MATCH       = 12012
	POST_APPLY_CREAM_FAIL       = 12013
	POST_APPLY_ALREADY_EXISTS   = 12014
	POST_REPORT_FAIL            = 12015
	POST_PARAMS_FAIL            = 12016

	// 分享/转发相关错误码 13001-14000
	SHARE_DATA_FAIL             = 13001


	// 社区相关错误码 14001-15000
	COMMUNITY_SECTION_NOT_EXISTS = 14001
	COMMUNITY_TOPIC_NOT_EXISTS   = 14002
	COMMUNITY_TOPICS_FAIL        = 14003
	COMMUNITY_SECTIONS_FAIL      = 14004
	COMMUNITY_TOPIC_FAIL         = 14005
	COMMUNITY_POSTS_BY_SECTION   = 14006
	COMMUNITY_POSTS_BY_TOPIC     = 14007

	// 私教相关错误码 15001-16000
	COACH_NOT_EXISTS             = 15001
	COACH_GET_LABEL_FAIL         = 15002
	COACH_PUB_EVALUATE_FAIL      = 15003
	COACH_ORDER_NOT_EXISTS       = 15004
	COACH_ORDER_NOT_SUCCESS      = 15005
	COACH_TYPE_FAIL              = 15006
	COACH_ID_NOT_MATCH           = 15007
	COACH_SCORE_INFO_FAIL        = 15008
	COACH_ALREADY_EVALUATE       = 15009

	// 预约错误 16001-17000
	APPOINTMENT_INVALID_INFO     = 16001
	APPOINTMENT_QUERY_NODE_FAIL  = 16002
	APPOINTMENT_INVALID_NODE_ID  = 16003
	APPOINTMENT_PROCESS_FAIL     = 16004
	APPOINTMENT_VIP_DEDUCTION    = 16005
	APPOINTMENT_NOT_ENOUGH_STOCK = 16006
	APPOINTMENT_ADD_RECORD_FAIL  = 16007
	APPOINTMENT_RECORD_ORDER_FAIL= 16008

	// 场馆错误 17001-18000
	VENUE_NOT_EXISTS             = 17001
	VENUE_PRODUCT_NOT_EXIST      = 17002
	VENUE_VIP_INFO_FAIL          = 17003
	VENUE_ADD_VIP_FAIL           = 17004
	VENUE_UPDATE_VIP_FAIL        = 17005
	VENUE_ACTION_RECORD_FAIL     = 17006

	// 订单错误 18001-19000
	ORDER_ADD_FAIL               = 18001
	ORDER_PRODUCT_ADD_FAIL       = 18002
	ORDER_NOT_EXISTS             = 18003
	ORDER_STATUS_FAIL            = 18004
	ORDER_UPDATE_FAIL            = 18005
	ORDER_ALREADY_DEL            = 18006
	ORDER_REFUND_FAIL            = 18007
	ORDER_NOT_ALLOW_REFUND       = 18008
	ORDER_DELETE_FAIL            = 18009
	ORDER_COUPON_CODE_FAIL       = 18010
	ORDER_CANCEL_FAIL            = 18011
	ORDER_NOT_ALLOW_CANCEL       = 18012
	ORDER_USER_NOT_MATCH         = 18013
	ORDER_PROCESS_FAIL           = 18014
	ORDER_ADD_REFUND_RECORD_FAIL = 18015
	ORDER_ADD_CARD_RECORD_FAIL   = 18016

	// 大课相关错误 19001-20000
	COURSE_NOT_EXISTS            = 19001
	COURSE_TYPE_FAIL             = 19002
	COURSE_ID_NOT_MATCH          = 19003

	// 支付相关错误 20001-30000
	PAY_INVALID_TYPE             = 20001
	PAY_ALI_PARAM_FAIL           = 20002
	PAY_WX_PARAM_FAIL            = 20003
	PAY_CHANNEL_NOT_EXISTS       = 20004

	// 资讯相关错误 30001-40000
	INFORMATION_LIST_FAIL        = 30001
	INFORMATION_DETAIL_FAIL      = 30002
	INFORMATION_NOT_EXISTS       = 30003

	// 赛事相关错误码 40001-50000
	CONTEST_GET_LIVE_FAIL        = 40001
	CONTEST_INFO_FAIL            = 40002
	CONTEST_SCHEDULE_FAIL        = 40003
	CONTEST_PROMOTION_INFO_FAIL  = 40004
	CONTEST_PLAYER_INFO_FAIL     = 40005
	CONTEST_RANKING_FAIL         = 40006
	CONTEST_LIVE_SCHEDULE_DATA   = 40007

	// 商城相关错误码 50001-60000
	SHOP_GET_ALL_SPU_FAIL        = 50001
	SHOP_GET_SPU_BY_CATEGORY_FAIL= 50002
	SHOP_RECOMMEND_FAIL          = 50003
	SHOP_PRODUCT_SKU_FAIL        = 50004
	SHOP_PRODUCT_SPU_FAIL        = 50005
	SHOP_ADD_USER_ADDR_FAIL      = 50006
	SHOP_UPDATE_USER_ADDR_FAIL   = 50007
	SHOP_USER_ADDR_NOT_FOUND     = 50008
	SHOP_GET_USER_ADDR_FAIL      = 50009
	SHOP_ADD_PRODUCT_CART_FAIL   = 50010
	SHOP_GET_PRODUCT_CART_FAIL   = 50011
	SHOP_UPDATE_PRODUCT_CART_FAIL= 50012
	SHOP_SKU_STOCK_NOT_ENOUGH    = 50013
	SHOP_PLACE_ORDER_FAIL        = 50014
	SHOP_ORDER_NOT_EXISTS        = 50015
	SHOP_ORDER_CANCEL_FAIL       = 50016
	SHOP_ORDER_NOT_ALLOW_CANCEL  = 50017
	SHOP_ORDER_UPDATE_FAIL       = 50018
	SHOP_ORDER_LIST_FAIL         = 50019
)

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	PAY_SUCCESS:    "success",
	INVALID_PARAMS: "请求参数错误",
	UNAUTHORIZED:   "未经授权",
	INVALID_TOKEN:  "鉴权失败，请重新登陆",

	USER_FREE_LOGIN_FAIL:    "一键登陆失败",
	USER_INVALID_MOBILE_NUM: "非法的手机号",
	USER_ALREADY_EXISTS:     "用户已存在",
	USER_REPEAT_REG:         "用户重复注册",
	USER_REGISTER_FAIL:      "用户注册失败",
	USER_ADD_INFO_FAIL:      "添加用户信息失败",
	USER_GET_INFO_FAIL:      "获取用户信息失败",
	USER_NOT_EXISTS:         "用户不存在",
	USER_INVALID_NAME_LEN:   "昵称长度最多30个字符（15个汉字），最少1个字符",
	USER_INVALID_SIGN_LEN:   "签名最多140个字符（70个汉字）",
	USER_NICK_NAME_EXISTS:   "昵称已存在",
	USER_AVATAR_NOT_EXISTS:  "系统头像不存在",
	USER_COUNTRY_NOT_EXISTS: "国家不存在",
	USER_UPDATE_INFO_FAIL:   "更新用户信息失败",
	USER_INVALID_NAME:       "昵称不合法",
	USER_INVALID_SIGNATURE:  "签名不合法",
	USER_FEEDBACK_FAIL:      "反馈提交失败",
	USER_ADD_NOTIFY_SET_FAIL:"系统设置初始化失败",
	USER_NO_LOGIN:           "用户未登录",
	USER_INVALID_FEEDBACK:   "反馈信息含有违规字段",
	USER_BIND_DEVICE_TOKEN:  "绑定设备token失败",
	USER_PACKAGE_NOT_EXISTS: "当前版本对应的下载包不存在",
	USER_LATEST_PACKAGE_FAIL:"获取最新包失败",
	USER_FORBID_STATUS:      "您的账号已被封禁",
	USER_GEN_IM_SIGN_FAIL:   "生成im签名失败",
	USER_UPDATE_IM_SIGN_FAIL:"更新im签名失败",
	USER_ADD_GUEST_FAIL:     "添加游客失败",
	USER_GET_GUEST_SIGN_FAIL:"获取游客签名失败",

	WX_USER_INFO_FAIL:    "获取微信用户信息失败",
	WX_ACCESS_TOKEN_FAIL: "获取微信授权token失败",
	WX_REGISTER_FAIL:     "微信注册帐号失败",
	WX_ADD_ACCOUNT_FAIL:  "微信帐号添加失败",

	WEIBO_USER_INFO_FAIL:   "获取微博用户信息失败",
	WEIBO_ADD_ACCOUNT_FAIL: "记录微博登陆信息失败",
	WEIBO_REGISTER_FAIL:    "微博用户注册失败",

	QQ_UNIONID_FAIL:       "获取QQ授权信息失败",
	QQ_USER_INFO_FAIL:     "获取QQ用户信息失败",
	QQ_REGISTER_FAIL:      "QQ注册账户失败",
	QQ_ADD_ACCOUNT_FAIL:   "QQ账户添加失败",

	ATTENTION_USER_NOT_EXISTS:   "关注的用户不存在",
	ATTENTION_ALREADY_EXISTS:    "已关注该用户",
	ATTENTION_USER_FAIL:         "关注失败",
	ATTENTION_RECORD_NOT_EXISTS: "未关注该用户",
	ATTENTION_REPEAT_CANCEL:     "已取消关注 请勿重复操作",
	ATTENTION_CANCEL_FAIL:       "取消关注失败",
	ATTENTION_YOURSELF_FAIL:     "不能关注自己",

	COLLECT_VIDEO_NOT_EXISTS:    "收藏的视频不存在",
	COLLECT_ALREADY_EXISTS:      "已收藏该视频",
	COLLECT_VIDEO_FAIL:          "收藏视频失败",
	COLLECT_RECORD_NOT_EXISTS:   "未收藏该视频",
	COLLECT_REPEAT_CANCEL:       "已取消收藏 请勿重复操作",
	COLLECT_CANCEL_FAIL:         "取消收藏失败",
	COLLECT_DELETE_FAIL:         "删除收藏的记录失败",

	VIDEO_NOT_EXISTS:            "视频不存在",
	VIDEO_PUBLISH_FAIL:          "视频发布失败",
	VIDEO_DELETE_HISTORY_FAIL:   "删除历史记录失败",
	VIDEO_DELETE_PUBLISH_FAIL:   "删除发布的视频失败",
	VIDEO_DELETE_LABEL_FAIL:     "删除视频标签失败",
	VIDEO_DELETE_STATISTIC_FAIL: "删除视频统计数据失败",
	VIDEO_UPLOAD_GEN_SIGN_FAIL:  "上传签名生成失败",
	VIDEO_INVALID_DESCRIBE:      "视频描述含有违规文字",
	VIDEO_INVALID_TITLE:         "视频标题含有违规文字",
	VIDEO_INVALID_CUSTOM_LABEL:  "自定义标签含有违规文字",
	VIDEO_REPORT_FAIL:           "举报视频失败",
	VIDEO_INVALID_PLAY_DURATION: "播放时长 > 视频时长！！wtf?",
	VIDEO_RECORD_PLAY_DURATION:  "记录用户播放的视频时长失败",
	VIDEO_SUBAREA_FAIL:          "获取视频分区失败",
	VIDEO_CREATE_ALBUM_FAIL:     "创建视频专辑失败",
	VIDEO_INVALID_ALBUM_NAME:    "专辑名称含有违规文字",
	VIDEO_ALBUM_NOT_EXISTS:      "视频专辑不存在",
	VIDEO_ADD_TO_ALBUM_FAIL:     "视频添加到专辑失败",
	VIDEO_LIST_BY_SUBAREA_FAIL:  "分区视频列表获取失败",

	LIKE_VIDEO_NOT_EXISTS:       "点赞的视频不存在",
	LIKE_ALREADY_EXISTS:         "已点过赞",
	LIKE_VIDEO_FAIL:             "视频点赞失败",
	LIKE_RECORD_NOT_EXISTS:      "未点赞该视频",
	LIKE_REPEAT_CANCEL:          "已取消点赞，请勿重复操作",
	LIKE_CANCEL_FAIL:            "取消点赞失败",
	LIKE_COMMENT_NOT_EXISTS:     "点赞的评论不存在",
	LIKE_COMMENT_FAIL:           "评论点赞失败",
	LIKE_POST_NOT_EXISTS:        "点赞的帖子不存在",
	LIKE_POST_FAIL:              "帖子点赞失败",
	LIKE_INFORMATION_NOT_EXISTS: "点赞的资讯不存在",
	LIKE_INFORMATION_FAIL:       "资讯点赞失败",

	NOTIFY_SETTING_FAIL:         "系统通知设置失败",

	COMMENT_PUBLISH_FAIL:    "发布评论失败",
	COMMENT_NOT_FOUND:       "评论不存在",
	COMMENT_REPLY_FAIL:      "评论回复失败",
	COMMENT_REPLY_NOT_FOUND: "未找到该评论相关回复",
	COMMENT_INVALID_LEN:     "评论最少10字符，最多1000字符",
	COMMENT_INVALID_CONTENT: "评论中含有违规文字",
	COMMENT_INVALID_REPLY:   "回复中含有违规文字",
	COMMENT_REPORT_FAIL:     "举报评论失败",
	COMMENT_DELETE_FAIL:     "删除评论失败",
	COMMENT_USER_NOT_MATCH:  "用户不匹配",

	SMS_CODE_INTERVAL_ERROR:     "一天内该手机获取验证码次数超限(最多10次)",
	SMS_CODE_INTERVAL_SHORT:     "获取短信验证间隔时间过短(间隔60秒)",
	SMS_CODE_INVALID_SEND_TYPE:  "无效的短信类型",
	SMS_CODE_SEND_FAIL:          "短信验证码发送失败",
	SMS_INVALID_CODE:            "无效的短信验证码",
	SMS_CODE_NOT_SEND:           "该手机未获取验证码",
	SMS_CODE_NOT_MATCH:          "短信验证码不正确",

	BARRAGE_VIDEO_SEND_FAIL:     "发送视频弹幕失败",
	BARRAGE_INVALID_CONTENT:     "弹幕内容含有违规文字",
	BARRAGE_VIDEO_LIST_FAIL:     "视频弹幕获取失败",
	BARRAGE_LIVE_SEND_FAIL:      "发送直播弹幕失败",

	SEARCH_CLEAN_HISTORY_FAIL:   "清空搜索记录失败",

	CLOUD_COS_ACCESS_FAIL:       "通行证获取失败",
	CLOUD_FILTER_FAIL:           "敏感词过滤失败",

	POST_INVALID_TITLE:          "帖子标题含有违规文字",
	POST_INVALID_CONTENT:        "帖子内容含有违规文字",
	POST_SECTION_NOT_EXISTS:     "板块不存在",
	POST_TOPIC_NOT_EXISTS:       "话题不存在",
	POST_PUBLISH_FAIL:           "发布帖子失败",
	POST_INVALID_CONTENT_LEN:    "内容长度超过限制",
	POST_DETAIL_FAIL:            "获取帖子详情失败",
	POST_NOT_EXISTS:             "帖子不存在",
	POST_DELETE_PUBLISH_FAIL:    "删除发布的帖子失败",
	POST_DELETE_TOPIC_FAIL:      "删除帖子标签失败",
	POST_DELETE_STATISTIC_FAIL:  "删除帖子统计数据失败",
	POST_AUTHOR_NOT_MATCH:       "帖子作者id与当前用户不匹对",
	POST_APPLY_CREAM_FAIL:       "申精失败",
	POST_APPLY_ALREADY_EXISTS:   "请勿重复申请",
	POST_REPORT_FAIL:            "举报失败",
	POST_PARAMS_FAIL:            "内容和图片都为空",

	SHARE_DATA_FAIL:             "分享失败",

	COMMUNITY_SECTION_NOT_EXISTS:"社区板块不存在",
	COMMUNITY_TOPIC_NOT_EXISTS:  "社区话题不存在",
	COMMUNITY_TOPICS_FAIL:       "获取社区话题列表失败",
	COMMUNITY_SECTIONS_FAIL:     "获取社区板块列表失败",
	COMMUNITY_TOPIC_FAIL:        "获取社区话题失败",
	COMMUNITY_POSTS_BY_SECTION:  "获取当前板块下的帖子失败",
	COMMUNITY_POSTS_BY_TOPIC:    "获取当前话题下的帖子失败",

	COACH_NOT_EXISTS:            "私教不存在",
	COACH_GET_LABEL_FAIL:        "获取私教标签失败",
	COACH_PUB_EVALUATE_FAIL:     "添加评价失败",
	COACH_ORDER_NOT_EXISTS:      "私教订单不存在",
	COACH_ORDER_NOT_SUCCESS:     "私教订单未成功",
	COACH_TYPE_FAIL:             "私教类型错误",
	COACH_ID_NOT_MATCH:          "私教id不匹配",
	COACH_SCORE_INFO_FAIL:       "获取私教评价信息失败",
	COACH_ALREADY_EVALUATE:      "已评价",

	APPOINTMENT_INVALID_INFO:    "预约信息错误", // 场次信息有误,请重新选择～
	APPOINTMENT_QUERY_NODE_FAIL: "查询时间节点配置错误",// 场次信息有误,请重新选择～
	APPOINTMENT_INVALID_NODE_ID: "错误的时间节点id",   // 场次信息有误,请重新选择～
	APPOINTMENT_PROCESS_FAIL:    "预约流程错误",       // 服务器开小差,请重试～ 需抛出code码
	APPOINTMENT_VIP_DEDUCTION:   "VIP抵扣时长错误",    // 预约信息有错,请重试～
	APPOINTMENT_NOT_ENOUGH_STOCK:"库存不足",
	APPOINTMENT_ADD_RECORD_FAIL: "添加预约流水失败",    // 服务器开小差,请重试～
	APPOINTMENT_RECORD_ORDER_FAIL: "记录预约订单号失败", // 服务器

	VENUE_NOT_EXISTS:            "场馆不存在",
	VENUE_PRODUCT_NOT_EXIST:     "商品不存在",
	VENUE_VIP_INFO_FAIL:         "获取场馆会员信息失败",
	VENUE_ADD_VIP_FAIL:          "添加会员失败",
	VENUE_UPDATE_VIP_FAIL:       "更新会员信息失败",
	VENUE_ACTION_RECORD_FAIL:    "获取进出场记录失败",

	ORDER_ADD_FAIL:              "添加订单失败",
	ORDER_PRODUCT_ADD_FAIL:      "添加商品订单失败",
	ORDER_NOT_EXISTS:            "订单不存在",
	ORDER_STATUS_FAIL:           "订单状态错误",
	ORDER_UPDATE_FAIL:           "更新订单信息失败",
	ORDER_ALREADY_DEL:           "订单已删除",
	ORDER_REFUND_FAIL:           "退款失败",
	ORDER_NOT_ALLOW_REFUND:      "订单不允许退款",
	ORDER_DELETE_FAIL:           "订单删除失败",
	ORDER_COUPON_CODE_FAIL:      "获取券码失败",
	ORDER_CANCEL_FAIL:           "订单取消失败",
	ORDER_NOT_ALLOW_CANCEL:      "订单不允许取消",
	ORDER_USER_NOT_MATCH:        "用户不配对",
	ORDER_PROCESS_FAIL:          "订单处理失败",
	ORDER_ADD_REFUND_RECORD_FAIL:"添加退款记录失败",
	ORDER_ADD_CARD_RECORD_FAIL:  "添加会员卡记录失败",

	COURSE_NOT_EXISTS:           "课程不存在",
	COURSE_TYPE_FAIL:            "课程类型错误",
	COURSE_ID_NOT_MATCH:         "课程id不匹配",

	PAY_INVALID_TYPE:            "无效的支付类型",
	PAY_ALI_PARAM_FAIL:          "获取支付宝请求参数失败",
	PAY_WX_PARAM_FAIL:           "获取微信请求参数失败",
	PAY_CHANNEL_NOT_EXISTS:      "支付渠道不存在",

	INFORMATION_LIST_FAIL:       "资讯列表获取失败",
	INFORMATION_DETAIL_FAIL:     "获取资讯详情失败",

	CONTEST_GET_LIVE_FAIL:       "获取赛事直播失败",
	CONTEST_INFO_FAIL:           "获取赛事信息失败",
	CONTEST_SCHEDULE_FAIL:       "获取赛程信息失败",
	CONTEST_PROMOTION_INFO_FAIL: "获取赛程晋级信息失败",
	CONTEST_PLAYER_INFO_FAIL:    "获取选手信息失败",
	CONTEST_RANKING_FAIL:        "获取选手积分排行失败",
	CONTEST_LIVE_SCHEDULE_DATA:  "获取赛程直播选手数据失败",

    SHOP_GET_ALL_SPU_FAIL:         "商品获取失败",
	SHOP_GET_SPU_BY_CATEGORY_FAIL: "分类商品获取失败",
	SHOP_RECOMMEND_FAIL:           "推荐商品获取失败",
	SHOP_PRODUCT_SKU_FAIL:         "商品sku获取失败",
	SHOP_PRODUCT_SPU_FAIL:         "商品spu获取失败",
	SHOP_ADD_USER_ADDR_FAIL:       "添加用户地址失败",
	SHOP_UPDATE_USER_ADDR_FAIL:    "更新用户地址失败",
	SHOP_USER_ADDR_NOT_FOUND:      "当前地址不存在",
	SHOP_GET_USER_ADDR_FAIL:       "地址信息获取失败",
	SHOP_ADD_PRODUCT_CART_FAIL:    "添加商品购物车失败",
	SHOP_GET_PRODUCT_CART_FAIL:    "商品购物车获取失败",
	SHOP_UPDATE_PRODUCT_CART_FAIL: "更新商品购物车失败",
	SHOP_SKU_STOCK_NOT_ENOUGH:     "库存不足",
	SHOP_PLACE_ORDER_FAIL:         "下单失败",
	SHOP_ORDER_NOT_EXISTS:         "订单不存在",
	SHOP_ORDER_CANCEL_FAIL:        "订单取消失败",
	SHOP_ORDER_NOT_ALLOW_CANCEL:   "订单不允许取消",
	SHOP_ORDER_UPDATE_FAIL:        "订单更新失败",
	SHOP_ORDER_LIST_FAIL:          "获取订单列表失败",
	
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}



