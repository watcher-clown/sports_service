package models

type VenueRecommendConf struct {
	Id       int    `json:"id" xorm:"not null pk autoincr comment('自增ID') INT(10)"`
	Name     string `json:"name" xorm:"not null comment('推荐名称') VARCHAR(60)"`
	Status   int    `json:"status" xorm:"not null default 0 comment('0 有效 1 废弃') TINYINT(1)"`
	CreateAt int    `json:"create_at" xorm:"not null default 0 comment('创建时间') INT(11)"`
	UpdateAt int    `json:"update_at" xorm:"not null default 0 comment('更新时间') INT(11)"`
	VenueId  int64  `json:"venue_id" xorm:"not null default 0 comment('场馆id') BIGINT(20)"`
}
