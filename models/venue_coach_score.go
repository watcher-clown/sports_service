package models

type VenueCoachScore struct {
	CoachId    int64 `json:"coach_id" xorm:"not null pk comment('私教id') BIGINT(20)"`
	TotalScore int   `json:"total_score" xorm:"not null default 0 comment('1颗星1分') INT(11)"`
	TotalNum   int   `json:"total_num" xorm:"not null default 0 comment('评分总人数') INT(11)"`
	Total5Star int   `json:"total5_star" xorm:"not null default 0 comment('五星评价总人数') INT(11)"`
	Total4Star int   `json:"total4_star" xorm:"not null default 0 comment('四星评价总人数') INT(11)"`
	Total3Star int   `json:"total3_star" xorm:"not null default 0 comment('三星评价总人数') INT(11)"`
	Total2Star int   `json:"total2_star" xorm:"not null default 0 comment('二星评价总人数') INT(11)"`
	Total1Star int   `json:"total1_star" xorm:"not null default 0 comment('一星评价总人数') INT(11)"`
	CreateAt   int   `json:"create_at" xorm:"not null comment('记录创建时间') INT(11)"`
	UpdateAt   int   `json:"update_at" xorm:"not null comment('记录更新时间') INT(11)"`
	Status     int   `json:"status" xorm:"not null default 0 comment('0 正常 1 废弃') TINYINT(1)"`
}
