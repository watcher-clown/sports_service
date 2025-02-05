package models

type User struct {
	Id            int64  `json:"id" xorm:"pk autoincr comment('主键id') BIGINT(20)"`
	NickName      string `json:"nick_name" xorm:"not null default '' comment('昵称') VARCHAR(45)"`
	MobileNum     int64  `json:"mobile_num" xorm:"not null comment('手机号码') unique BIGINT(20)"`
	Password      string `json:"password" xorm:"not null comment('用户密码') VARCHAR(128)"`
	UserId        string `json:"user_id" xorm:"not null comment('用户id') unique VARCHAR(60)"`
	Gender        int    `json:"gender" xorm:"not null default 0 comment('性别 0人妖 1男性 2女性') TINYINT(1)"`
	Born          string `json:"born" xorm:"not null default '' comment('出生日期') VARCHAR(128)"`
	Age           int    `json:"age" xorm:"not null default 0 comment('年龄') INT(3)"`
	Avatar        string `json:"avatar" xorm:"not null default '' comment('头像') VARCHAR(300)"`
	Status        int    `json:"status" xorm:"default 0 comment('1 正常 -1 封禁') TINYINT(1)"`
	LastLoginTime int    `json:"last_login_time" xorm:"comment('最后登录时间') INT(11)"`
	Signature     string `json:"signature" xorm:"not null default '' comment('签名') VARCHAR(200)"`
	DeviceType    int    `json:"device_type" xorm:"comment('设备类型 0 android 1 iOS 2 小程序 3 web') TINYINT(2)"`
	City          string `json:"city" xorm:"not null default '' comment('城市') VARCHAR(64)"`
	IsAnchor      int    `json:"is_anchor" xorm:"not null default 0 comment('0不是主播 1为主播') TINYINT(1)"`
	ChannelId     int    `json:"channel_id" xorm:"not null default 0 comment('渠道id') INT(11)"`
	BackgroundImg string `json:"background_img" xorm:"not null default '' comment('背景图') VARCHAR(255)"`
	Title         string `json:"title" xorm:"not null default '' comment('称号/特殊身份') VARCHAR(255)"`
	CreateAt      int    `json:"create_at" xorm:"not null default 0 comment('创建时间') INT(11)"`
	UpdateAt      int    `json:"update_at" xorm:"not null default 0 comment('更新时间') INT(11)"`
	UserType      int    `json:"user_type" xorm:"not null default 0 comment('用户类型 0 手机号 1 微信 2 QQ 3 微博') TINYINT(2)"`
	Country       int    `json:"country" xorm:"not null default 0 comment('国家') INT(3)"`
	RegIp         string `json:"reg_ip" xorm:"default ' ' comment('注册ip') VARCHAR(30)"`
	DeviceToken   string `json:"device_token" xorm:"not null default '' comment('设备token') VARCHAR(100)"`
	AccountType   int    `json:"account_type" xorm:"not null default 0 comment('账号类型 0 普通用户 1 官方账号') TINYINT(2)"`
	TxAccid       string `json:"tx_accid" xorm:"not null default '' comment('accid（腾讯im 唯一标识)') VARCHAR(60)"`
	TxToken       string `json:"tx_token" xorm:"not null default '' comment('token（腾讯im）') VARCHAR(300)"`
}
