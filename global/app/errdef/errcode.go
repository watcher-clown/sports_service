package errdef

const (
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

	// 收藏相关错误码 3001-4000
	COLLECT_VIDEO_NOT_EXISTS    = 3001
	COLLECT_ALREADY_EXISTS      = 3002
	COLLECT_VIDEO_FAIL          = 3003
	COLLECT_RECORD_NOT_EXISTS   = 3004
	COLLECT_REPEAT_CANCEL       = 3005
	COLLECT_CANCEL_FAIL         = 3006
	COLLECT_DELETE_FAIL         = 3007

	// 视频相关错误码 4001-5000
	VIDEO_NOT_EXISTS          = 4001
	VIDEO_PUBLISH_FAIL        = 4002
	VIDEO_DELETE_HISTORY_FAIL = 4003
	VIDEO_DELETE_PUBLISH_FAIL = 4004

	// 点赞相关错误码 5001-6000
	LIKE_VIDEO_NOT_EXISTS       = 5001
	LIKE_ALREADY_EXISTS         = 5002
	LIKE_VIDEO_FAIL             = 5003
	LIKE_RECORD_NOT_EXISTS      = 5004
	LIKE_REPEAT_CANCEL          = 5005
	LIKE_CANCEL_FAIL            = 5006

	// 通知相关错误吗 6001-7000
	NOTIFY_SETTING_FAIL         = 6001
)

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
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

	COLLECT_VIDEO_NOT_EXISTS:    "收藏的视频不存在",
	COLLECT_ALREADY_EXISTS:      "已收藏该视频",
	COLLECT_VIDEO_FAIL:          "收藏视频失败",
	COLLECT_RECORD_NOT_EXISTS:   "未收藏该视频",
	COLLECT_REPEAT_CANCEL:       "已取消收藏 请勿重复操作",
	COLLECT_CANCEL_FAIL:         "取消收藏失败",
	COLLECT_DELETE_FAIL:         "删除收藏的记录失败",

	VIDEO_NOT_EXISTS:          "视频不存在",
	VIDEO_PUBLISH_FAIL:        "视频发布失败",
	VIDEO_DELETE_HISTORY_FAIL: "删除历史记录失败",
	VIDEO_DELETE_PUBLISH_FAIL: "删除发布的视频失败",

	LIKE_VIDEO_NOT_EXISTS:       "点赞的视频不存在",
	LIKE_ALREADY_EXISTS:         "已点过赞",
	LIKE_VIDEO_FAIL:             "视频点赞失败",
	LIKE_RECORD_NOT_EXISTS:      "未点赞该视频",
	LIKE_REPEAT_CANCEL:          "已取消收藏，请勿重复操作",
	LIKE_CANCEL_FAIL:            "取消点赞失败",

	NOTIFY_SETTING_FAIL:         "系统通知设置失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}



