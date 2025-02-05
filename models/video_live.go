package models

type VideoLive struct {
	Id            int64  `json:"id" xorm:"pk autoincr comment('主键') BIGINT(20)"`
	UserId        string `json:"user_id" xorm:"not null default '' comment('主播id') VARCHAR(60)"`
	RoomId        string `json:"room_id" xorm:"not null default '' comment('房间id') index VARCHAR(60)"`
	GroupId       string `json:"group_id" xorm:"not null default '' comment('im群组id') VARCHAR(60)"`
	Cover         string `json:"cover" xorm:"not null default '' comment('直播封面') VARCHAR(512)"`
	RtmpAddr      string `json:"rtmp_addr" xorm:"not null default '' comment('rtmp地址[拉流]') VARCHAR(512)"`
	FlvAddr       string `json:"flv_addr" xorm:"not null default '' comment('flv地址[拉流]') VARCHAR(512)"`
	HlsAddr       string `json:"hls_addr" xorm:"not null default '' comment('hls地址[拉流]') VARCHAR(512)"`
	PushStreamUrl string `json:"push_stream_url" xorm:"not null default '' comment('推流url') VARCHAR(255)"`
	PlayTime      int    `json:"play_time" xorm:"not null default 0 comment('后台设置的赛事开播时间') INT(11)"`
	EndTime       int    `json:"end_time" xorm:"not null default 0 comment('结束时间') INT(11)"`
	Status        int    `json:"status" xorm:"default 0 comment('状态 0未直播 1直播中 2 已结束') TINYINT(1)"`
	Title         string `json:"title" xorm:"not null default '' comment('标题') VARCHAR(255)"`
	HighLights    string `json:"high_lights" xorm:"not null default '' comment('亮点') VARCHAR(255)"`
	Describe      string `json:"describe" xorm:"not null default '' comment('描述') VARCHAR(512)"`
	Tags          string `json:"tags" xorm:"not null default '' comment('直播标签') VARCHAR(512)"`
	CreateAt      int    `json:"create_at" xorm:"not null default 0 comment('创建时间') INT(11)"`
	UpdateAt      int    `json:"update_at" xorm:"not null default 0 comment('更新时间') INT(11)"`
	Duration      int64  `json:"duration" xorm:"default 0 comment('时长（毫秒）') BIGINT(20)"`
	LiveType      int    `json:"live_type" xorm:"not null default 0 comment('直播类型（0：管理员[sys_user]，1：用户[user]）') TINYINT(1)"`
	ContestId     int    `json:"contest_id" xorm:"not null default 0 comment('赛事id') INT(11)"`
	ScheduleId    int    `json:"schedule_id" xorm:"not null default 0 comment('赛程id') INT(11)"`
	StartTime     int    `json:"start_time" xorm:"not null default 0 comment('真实开播时间') INT(11)"`
	PushStreamKey string `json:"push_stream_key" xorm:"not null default '' comment('推流密钥') VARCHAR(255)"`
	Subhead       string `json:"subhead" xorm:"not null default '' comment('副标题') VARCHAR(255)"`
}
