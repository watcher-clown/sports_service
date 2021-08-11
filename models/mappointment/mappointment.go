package mappointment

import (
	"github.com/go-xorm/xorm"
	"sports_service/server/models"
)

type AppointmentModel struct {
	AppointmentInfo  *models.VenueAppointmentInfo
	Engine           *xorm.Session
	Stock            *models.VenueAppointmentStock
	Record           *models.AppointmentRecord
	Vip              *models.VenueVipInfo
}

type WeekInfo struct {
	Date      string    `json:"date"`
	Week      int       `json:"week"`
	Id        int32     `json:"id"`
	WeekCn    string    `json:"week_cn"`
	MinPrice  int       `json:"min_price"`
	PriceCn   string    `json:"price_cn"`
}

type OptionsInfo struct {
	Id              int64  `json:"id"`
	TimeNode        string `json:"time_node"`
	Duration        int    `json:"duration"`
	DurationCn      string `json:"duration_cn"`
	RealAmount      int    `json:"real_amount"`
	RealAmountCn    string `json:"real_amount_cn"`
	CurAmount       int    `json:"cur_amount"`
	DiscountRate    int    `json:"discount_rate"`
	RateCn          string `json:"rate_cn"`
	DiscountAmount  int    `json:"discount_amount"`
	QuotaNum        int    `json:"quota_num"`
	RelatedId       int64  `json:"related_id"`
	RecommendType   int    `json:"recommend_type"`
	AppointmentType int    `json:"appointment_type"`
	WeekNum         int    `json:"week_num"`

	HasDiscount     int    `json:"has_discount,omitempty"`       // 是否有优惠 0无 1有
	AmountCn        string `json:"amount_cn,omitempty"`          // 中文价格
	IsFull          int    `json:"is_full"`                      // 是否满场
	PurchasedNum    int    `json:"purchased_num,omitempty"`      // 已购买人数 包含[成功购买及已下单]
	Name            string `json:"name,omitempty"`               // 场馆名称
	Avatar          string `json:"avatar,omitempty"`             // 大课老师头像
	Address         string `json:"address,omitempty"`            // 上课地点
	Labels          []*LabelInfo     `json:"labels,omitempty"`   // 标签列表
	ReservedUsers   []*ReservedUsers `json:"reserved_users"`     // 已预约人数

}

// 已预约人数
type ReservedUsers struct {
	UserId      string       `json:"user_id"`
	NickName    string       `json:"nick_name"`
	Avatar      string       `json:"avatar"`
}

type LabelInfo struct {
	UserId      string       `json:"user_id"`
	//NickName    string       `json:"nick_name"`
	//Avatar      string       `json:"avatar"`
	LabelId     int64        `json:"label_id"`
	LabelName   string       `json:"label_name"`
}

type Options struct {
	Id                 int64        `json:"id"`
	Name               string       `json:"name"`
	Describe           string       `json:"describe,omitempty"`
	CostDescription    string       `json:"cost_description,omitempty"`     // 费用说明
	Instructions       string       `json:"instructions,omitempty"`         // 购买须知
}

type AppointmentReq struct {
	AppointmentType int                 `json:"appointment_type"` // 0场馆 1私教课 2大课
	UserId          string              `json:"user_id"`          // 用户id
	Infos           []*AppointmentInfo  `json:"infos"`            // 预约请求数据 1个时间点对应1条数据
	Ids             []interface{}       `json:"ids"`              // 时间配置ids
	LabelIds        []int64             `json:"label_ids"`        // 用户选择的标签id列表
}

// 预约请求数据
type AppointmentInfo struct {
	RelatedId int64      `json:"id"`         // 场馆id/私教课程id/大课老师id
	DateId    int        `json:"date_id"`    // 1为今天
	//TimeNode  string     `json:"time_node"`  // 预约时间节点
	Count     int        `json:"count"`            // 数量
	Id        int64      `json:"id"`               // 时间配置id
	SeatNo    int        `json:"seat_no"`          // 座位号
}

// 预约请求返回库存数据
type StockInfoResp struct {
	Id        int64      `json:"id"`               // 时间配置id
	Stock     int        `json:"stock"`            // 剩余数量
}

func NewAppointmentModel(engine *xorm.Session) *AppointmentModel {
	return &AppointmentModel{
		AppointmentInfo: new(models.VenueAppointmentInfo),
		Stock: new(models.VenueAppointmentStock),
		Record: new(models.AppointmentRecord),
		Vip: new(models.VenueVipInfo),
		Engine: engine,
	}
}

const (
	QUERY_MIN_PRICE = "SELECT min(cur_amount) as cur_amount FROM venue_appointment_info WHERE week_num=? AND appointment_type=? AND status=0"
)
// 根据星期 及 预约类型 获取最低价格
func (m *AppointmentModel) GetMinPriceByWeek() error {
	ok, err := m.Engine.SQL(QUERY_MIN_PRICE, m.AppointmentInfo.WeekNum, m.AppointmentInfo.AppointmentType).Get(m.AppointmentInfo)
	if !ok || err != nil {
		return err
	}

	return nil
}

// 通过id获取预约配置
func (m *AppointmentModel) GetAppointmentConfById() (bool, error) {
	return m.Engine.Get(m.AppointmentInfo)
}

// 通过ids[多个id]获取预约配置列表
func (m *AppointmentModel) GetAppointmentConfByIds(ids []interface{}) ([]*models.VenueAppointmentInfo, error) {
	var list []*models.VenueAppointmentInfo
	if err := m.Engine.In("id", ids...).OrderBy("cur_amount DESC").Find(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// 通过场馆id、课程id、星期 及 预约类型 获取可预约选项
func (m *AppointmentModel) GetOptionsByWeek() ([]*models.VenueAppointmentInfo, error) {
	var list []*models.VenueAppointmentInfo
	if err := m.Engine.Where("related_id=? AND week_num=? AND appointment_type=? AND status=0", m.AppointmentInfo.RelatedId,
		m.AppointmentInfo.WeekNum, m.AppointmentInfo.AppointmentType).Asc("id").Find(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// 查库存表 获取某时间点 场馆预约人数 包含已成功及已下单且订单未超时
func (m *AppointmentModel) GetPurchaseNum() (bool, error) {
	return m.Engine.Get(m.Stock)
}

// 是否存在库存信息 todo: 前期并发不高 可以读取快照
func (m *AppointmentModel) HasExistsStockInfo() (bool, error) {
	return m.Engine.Exist(m.AppointmentInfo)
}

// 添加预约库存数据[多条]
func (m *AppointmentModel) AddMultiStockInfo(stock []*models.VenueAppointmentStock) (int64, error) {
	return m.Engine.InsertMulti(stock)
}

// 添加预约库存数据[单条]
func (m *AppointmentModel) AddStockInfo(stock *models.VenueAppointmentStock) (int64, error) {
	return m.Engine.Insert(stock)
}

const (
	UPDATE_STOCK_INFO = "UPDATE `venue_appointment_stock` SET `purchased_num`= `purchased_num`+ ?, `update_at`=? WHERE date=? AND " +
		"time_node=? AND appointment_type=? AND related_id=? AND quota_num >= `purchased_num`+ ? LIMIT 1"
)
func (m *AppointmentModel) UpdateStockInfo(timeNode, date string, count, now, appointmentType, relatedId int) (int64, error) {
	res, err := m.Engine.Exec(UPDATE_STOCK_INFO, count, now, date, timeNode, appointmentType, relatedId, count)
	if err != nil {
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, nil
}

// 获取成功预约的记录[包含已付款和支付中]
func (m *AppointmentModel) GetAppointmentRecord() ([]*models.AppointmentRecord, error) {
	var list []*models.AppointmentRecord
	sql := "SELECT * FROM appointment_record WHERE status in(0, 2) AND appointment_type=? AND related_id=? AND time_node=? " +
		"AND date=? ORDER BY id ASC"
	if err := m.Engine.SQL(sql, m.Record.AppointmentType, m.Record.RelatedId, m.Record.TimeNode, m.Record.Date).Find(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// 获取场馆会员数据
func (m *AppointmentModel) GetVenueVipInfo(userId string) (*models.VenueVipInfo, error) {
	ok, err := m.Engine.Where("user_id=?", userId).Get(m.Vip)
	if !ok || err != nil {
		return nil, err
	}

	return m.Vip, nil
}

const (
	UPDATE_VENUE_VIP_INFO = "UPDATE `venue_vip_info` SET `duration` = `duration` + ?, `update_at`=? WHERE `user_id`=? " +
		"AND `duration` + ? >= 0 LIMIT 1"
)
// 更新场馆会员数据 duration < 0 减少 duration > 0 增加
func (m *AppointmentModel) UpdateVenueVipInfo(duration int, userId string) (int64, error) {
	res, err := m.Engine.Exec(UPDATE_VENUE_VIP_INFO, duration, userId)
	if err != nil {
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, nil
}

// 批量添加预约记录
func (m *AppointmentModel) AddMultiAppointmentRecord(list []*models.AppointmentRecord) (int64, error) {
	return m.Engine.InsertMulti(list)
}
