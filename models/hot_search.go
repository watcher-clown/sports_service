package models

type HotSearch struct {
	CreateAt         int    `json:"create_at" xorm:"not null comment('创建时间') INT(11)"`
	HotSearchContent string `json:"hot_search_content" xorm:"not null default '' comment('热门搜索内容 如：FPV、电竞') VARCHAR(256)"`
	Id               int    `json:"id" xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	Sortorder        int    `json:"sortorder" xorm:"not null default 0 comment('排序权重') INT(11)"`
	Status           int    `json:"status" xorm:"default 0 comment('0 展示 1 隐藏') TINYINT(1)"`
	UpdateAt         int    `json:"update_at" xorm:"not null comment('更新时间') INT(11)"`
}
