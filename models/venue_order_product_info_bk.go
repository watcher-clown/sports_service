package models

type VenueOrderProductInfoBk struct {
	Id          int64  `json:"id" xorm:"pk autoincr comment('自增ID') BIGINT(20)"`
	PayOrderId  string `json:"pay_order_id" xorm:"not null comment('订单号') index VARCHAR(150)"`
	ProductId   int64  `json:"product_id" xorm:"not null comment('商品id') BIGINT(20)"`
	ProductType int    `json:"product_type" xorm:"not null comment('1001 场馆预约 
2101 临时卡 2201 次卡 2311 月卡 2321 季卡 2331 半年卡 2341 年卡 
3001 私教（教练）订单 3002 课程订单 4001 充值订单  5101 线下实体商品') INT(8)"`
	Count           int   `json:"count" xorm:"not null comment('购买数量') INT(11)"`
	RealAmount      int   `json:"real_amount" xorm:"not null comment('[单个商品]定价（单位：分）') INT(11)"`
	CurAmount       int   `json:"cur_amount" xorm:"not null comment('[单个商品]当前价格 [售价](包含真实价格、 折扣价格（单位：分）') INT(11)"`
	DiscountRate    int   `json:"discount_rate" xorm:"not null default 0 comment('折扣率') INT(11)"`
	DiscountAmount  int   `json:"discount_amount" xorm:"not null default 0 comment('[单个商品]优惠的金额') INT(11)"`
	Amount          int   `json:"amount" xorm:"not null comment('商品总价') INT(11)"`
	ReceiveAmount   int   `json:"receive_amount" xorm:"not null default 0 comment('充值金额（钱包）') INT(11)"`
	Duration        int   `json:"duration" xorm:"not null default 0 comment('购买相关服务总时长') INT(11)"`
	CreateAt        int   `json:"create_at" xorm:"not null default 0 comment('创建时间') INT(11)"`
	UpdateAt        int   `json:"update_at" xorm:"not null default 0 comment('更新时间') INT(11)"`
	Status          int   `json:"status" xorm:"not null default 0 comment('0 待支付 1 订单超时/未支付/已取消 2 已支付 3 已完成  4 退款中 5 已退款 6 已过期') TINYINT(2)"`
	RelatedId       int64 `json:"related_id" xorm:"not null default 0 comment('场馆id/私教课id/大课id') BIGINT(20)"`
	DeductionTm     int64 `json:"deduction_tm" xorm:"not null default 0 comment('抵扣会员时长') BIGINT(20)"`
	DeductionAmount int64 `json:"deduction_amount" xorm:"not null default 0 comment('抵扣的价格') BIGINT(20)"`
	DeductionNum    int64 `json:"deduction_num" xorm:"not null default 0 comment('抵扣的数量') BIGINT(20)"`
	ExpireDuration  int   `json:"expire_duration" xorm:"not null default 0 comment('过期时长[单个商品]') INT(11)"`
	CoachId         int64 `json:"coach_id" xorm:"not null default 0 comment('教练id 包含私教老师/大课老师') BIGINT(20)"`
	SingleDuration  int   `json:"single_duration" xorm:"not null default 0 comment('单个时长') INT(11)"`
}
