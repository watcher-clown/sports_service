package errdef

const (
	SUCCESS            = 200
	ERROR              = 500
	INVALID_PARAMS     = 400
	UNAUTHORIZED       = 401

	INVALID_TOKEN      = 888

	// 用户相关错误码 1000-2000
	USER_FREE_LOGIN_FAIL    = 1000
	USER_INVALID_MOBILE_NUM = 1001
	USER_ALREADY_EXISTS     = 1002
	USER_REPEAT_REG         = 1003
	USER_REGISTER_FAIL      = 1004
	USER_ADD_INFO_FAIL      = 1005
	USER_GET_INFO_FAIL      = 1006
	USER_NOT_EXISTS         = 1007
	USER_INVALID_NAME_LEN   = 1008
	USER_INVALID_SIGN_LEN   = 1009
	USER_NICK_NAME_EXISTS   = 1010
	USER_AVATAR_NOT_EXISTS  = 1011
	USER_COUNTRY_NOT_EXISTS = 1012
	USER_UPDATE_INFO_FAIL   = 1013
	USER_INVALID_NAME       = 1014
	USER_INVALID_SIGNATURE  = 1015

	WX_USER_INFO_FAIL       = 1100
	WX_ACCESS_TOKEN_FAIL    = 1101
	WX_REGISTER_FAIL        = 1102
	WX_ADD_ACCOUNT_FAIL     = 1103

	WEIBO_USER_INFO_FAIL    = 1200
	WEIBO_ADD_ACCOUNT_FAIL  = 1201
	WEIBO_REGISTER_FAIL     = 1202

	QQ_UNIONID_FAIL         = 1301
	QQ_USER_INFO_FAIL       = 1302
	QQ_REGISTER_FAIL        = 1303
	QQ_ADD_ACCOUNT_FAIL     = 1304

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
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}



