package mstat

import (
	"github.com/go-xorm/xorm"
	"sports_service/server/models"
)

type StatModel struct {
	Engine    *xorm.Session
}

type Stat struct {
	Sum     int64   `json:"sum"`
	Count   int64   `json:"count"`
	Avg     int64   `json:"avg"`
	Dt      string  `json:"dt"`
	Subarea int64   `json:"subarea,omitempty"`
	Rate    string  `json:"rate,omitempty"`
}

// 管理后台首页统计数据
type HomePageInfo struct {
	TopInfo        map[string]interface{}      `json:"top_info"`         // 顶部统计数据
	DauList        []*Stat                      `json:"dau_list"`         // 日活数据
	NewUserList    []*Stat                      `json:"new_user_list"`    // 新增用户数据
	RetentionRate  []*RetentionRateInfo         `json:"retention_rate"`   // 留存率数据

	NextDayRetentionRate []*RetentionRateInfo   `json:"next_day_retention_rate"` // 次日留存率数据
}

func NewStatModel(engine *xorm.Session) *StatModel {
	return &StatModel{
		Engine: engine,
	}
}

// 获取总用户数
func (m *StatModel) GetTotalUser() (int64, error) {
	return m.Engine.Count(&models.User{})
}

const (
	GET_DAU_BY_DATE = "SELECT count(DISTINCT user_id) AS count FROM `user_activity_record` WHERE date(FROM_UNIXTIME(create_at))=?"
)
// 通过日期获取日活
func (m *StatModel) GetDAUByDate(date string) (Stat, error) {
	stat := Stat{}
	if ok, err := m.Engine.SQL(GET_DAU_BY_DATE, date).Get(&stat); !ok || err != nil {
		return stat, err
	}

	return stat, nil
}

const (
	GET_MAU_BY_MONTH = "SELECT count(DISTINCT user_id) AS count FROM `user_activity_record` WHERE " +
		"LEFT(date(FROM_UNIXTIME(create_at)), 7)=?"
)
// 通过年月获取月活
func (m *StatModel) GetMAUByMonth(month string) (Stat, error) {
	stat := Stat{}
	if ok, err := m.Engine.SQL(GET_MAU_BY_MONTH, month).Get(&stat); !ok || err != nil {
		return stat, err
	}

	return stat, nil
}


const (
	GET_NET_ADDITION_BY_DATE = "SELECT count(1) AS count FROM `user` WHERE date(FROM_UNIXTIME(create_at))=?"
)
// 通过日期[年月日] 获取新增用户数
func (m *StatModel) GetNetAdditionByDate(date string) (Stat, error) {
	stat := Stat{}
	if ok, err := m.Engine.SQL(GET_NET_ADDITION_BY_DATE, date).Get(&stat); !ok || err != nil {
		return stat, err
	}

	return stat, nil
}

const (
	GET_DAU_BY_DAYS = "SELECT count(DISTINCT user_id) AS count, date(FROM_UNIXTIME(create_at)) AS dt FROM " +
		"user_activity_record WHERE DATE_SUB(CURDATE(), INTERVAL ? DAY) <= date(FROM_UNIXTIME(create_at)) GROUP BY dt"
)
// 获取N天的日活数据
func (m *StatModel) GetDAUByDays(days int) ([]*Stat, error) {
	var dauList []*Stat
	if err := m.Engine.SQL(GET_DAU_BY_DAYS, days).Find(&dauList); err != nil {
		return dauList, err
	}

	return dauList, nil
}

const (
	GET_NET_ADDITION_BY_DAYS = "SELECT count(1) AS sum, date(FROM_UNIXTIME(create_at)) AS dt FROM user WHERE " +
		"DATE_SUB(CURDATE(), INTERVAL ? DAY) <= date(FROM_UNIXTIME(create_at)) GROUP BY dt"
)
// 获取N天的新增用户数据
func (m *StatModel) GetNetAdditionByDays(days int) ([]*Stat, error) {
	var statList []*Stat
	if err := m.Engine.SQL(GET_NET_ADDITION_BY_DAYS, days).Find(&statList); err != nil {
		return statList, err
	}

	return statList, nil
}


// 留存率信息
type RetentionRateInfo struct {
	Dt                string      `json:"dt"`
	NewUsers          int64       `json:"new_users"`
	NextDayRate       string      `json:"next_day_rate"`
	TwoDayRate        string      `json:"two_day_rate"`
	ThreeDayRate      string      `json:"three_day_rate"`
	FourDayRate       string      `json:"four_day_rate"`
	FiveDayRate       string      `json:"five_day_rate"`
	SixDayRate        string      `json:"six_day_rate"`
	OneWeekRate       string      `json:"one_week_rate"`
	TwoWeekRate       string      `json:"two_week_rate"`
	ThirtyDayRate     string      `json:"thirty_day_rate"`
	NinetyDayRate     string      `json:"ninety_day_rate"`
	HalfYearRate      string      `json:"half_year_rate"`
}
// 获取用户留存率 queryType != 1 只查次日留存率
func (m *StatModel) GetUserRetentionRate(queryType, minDate, maxDate string) ([]*RetentionRateInfo, error) {
	var rateList []*RetentionRateInfo
	sql := "SELECT " +
	    "date(FROM_UNIXTIME(u.create_at)) dt," +
		"count(DISTINCT u.user_id) new_users," +
		"concat(round(sum(DISTINCT datediff(from_unixtime(uar.create_at), from_unixtime(u.create_at))=1) / count(DISTINCT u.user_id) * 100, 2), '%') next_day_rate"

	if queryType == "1" {
		sql += ",concat(round(sum(DISTINCT datediff(from_unixtime(uar.create_at), from_unixtime(u.create_at))=2) / count(DISTINCT u.user_id) * 100, 2), '%') two_day_rate," +
			"concat(round(sum(DISTINCT datediff(from_unixtime(uar.create_at), from_unixtime(u.create_at))=3) / count(DISTINCT u.user_id) * 100, 2), '%') three_day_rate," +
			"concat(round(sum(DISTINCT datediff(from_unixtime(uar.create_at), from_unixtime(u.create_at))=4) / count(DISTINCT u.user_id) * 100, 2), '%') four_day_rate," +
			"concat(round(sum(DISTINCT datediff(from_unixtime(uar.create_at), from_unixtime(u.create_at))=5) / count(DISTINCT u.user_id) * 100, 2), '%') five_day_rate," +
			"concat(round(sum(DISTINCT datediff(from_unixtime(uar.create_at), from_unixtime(u.create_at))=6) / count(DISTINCT u.user_id) * 100, 2), '%') six_day_rate," +
			"concat(round(sum(DISTINCT datediff(from_unixtime(uar.create_at), from_unixtime(u.create_at))=7) / count(DISTINCT u.user_id) * 100, 2), '%') one_week_rate," +
			"concat(round(sum(DISTINCT datediff(from_unixtime(uar.create_at), from_unixtime(u.create_at))=14) / count(DISTINCT u.user_id) * 100, 2), '%') two_week_rate," +
			"concat(round(sum(DISTINCT datediff(from_unixtime(uar.create_at), from_unixtime(u.create_at))=30) / count(DISTINCT u.user_id) * 100, 2), '%') thirty_day_rate," +
			"concat(round(sum(DISTINCT datediff(from_unixtime(uar.create_at), from_unixtime(u.create_at))=90) / count(DISTINCT u.user_id) * 100, 2), '%') ninety_day_rate," +
			"concat(round(sum(DISTINCT datediff(from_unixtime(uar.create_at), from_unixtime(u.create_at))=180) / count(DISTINCT u.user_id) * 100, 2), '%') half_year_rate "
	}

	sql += " FROM user AS u LEFT JOIN user_activity_record AS uar ON u.user_id=uar.user_id"
	if minDate != "" && maxDate != "" {
		sql += " WHERE date(FROM_UNIXTIME(u.create_at)) >= ? AND date(FROM_UNIXTIME(u.create_at)) <= ? "
	}

	sql += "GROUP BY date(FROM_UNIXTIME(u.create_at))"

	if err := m.Engine.SQL(sql, minDate, maxDate).Find(&rateList); err != nil {
		return rateList, err
	}

	return rateList, nil
}

const (
	GET_VIDEO_SUBAREA_STAT = "SELECT subarea, count(1) AS count FROM videos WHERE status=1 GROUP BY subarea"
)
// 视频分区统计 [发布占比]
func (m *StatModel) GetVideoSubareaStat() ([]Stat, error) {
	var stat []Stat
	if err := m.Engine.SQL(GET_VIDEO_SUBAREA_STAT).Find(&stat); err != nil {
		return stat, err
	}

	return stat, nil
}

// 获取视频总数 [已审核的视频]
func (m *StatModel) GetVideoTotal() (int64, error) {
	return m.Engine.Where("status=1").Count(&models.Videos{})
}

