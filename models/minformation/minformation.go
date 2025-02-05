package minformation

import (
	"github.com/go-xorm/xorm"
	"sports_service/models"
	"sports_service/tools/tencentCloud"
)

type InformationModel struct {
	Engine      *xorm.Session
	Information *models.Information
	Statistic   *models.InformationStatistic
	Browse      *models.UserBrowseRecord
}

type InformationResp struct {
	Id        int64                  `json:"id"`
	Cover     tencentCloud.BucketURI `json:"cover"`
	Title     string                 `json:"title"`
	Content   string                 `json:"content"`
	Describe  string                 `json:"describe,omitempty"`
	PubType   int                    `json:"pub_type,omitempty"` // 1. 发布至赛事模块 2. 发布至视频首页板块',
	RelatedId int64                  `json:"related_id"`
	Name      string                 `json:"name"`
	//JumpUrl     string `json:"jump_url"`
	CreateAt    int                    `json:"create_at"`
	UserId      string                 `json:"user_id"`
	Avatar      tencentCloud.BucketURI `json:"avatar"`
	NickName    string                 `json:"nick_name"`
	CommentNum  int                    `json:"comment_num"`
	FabulousNum int                    `json:"fabulous_num"`
	IsAttention int                    `json:"is_attention"` // 是否关注 1 关注 0 未关注
	BrowseNum   int                    `json:"browse_num"`   // 浏览数
	IsLike      int                    `json:"is_like"`      // 是否点赞 1 点赞 0 未点赞
	ShareNum    int                    `json:"share_num"`    // 分享数
	Status      int                    `json:"status"`       // 0：审核中，1：审核通过 2：审核不通过 3：逻辑删除
	//IsCollect   int    `json:"is_collect"`               // 是否收藏 1 收藏 0 未收藏
}

func NewInformationModel(engine *xorm.Session) *InformationModel {
	return &InformationModel{
		Information: new(models.Information),
		Statistic:   new(models.InformationStatistic),
		Engine:      engine,
	}
}

// 获取资讯列表
func (m *InformationModel) GetInformationList(condition string, offset, size int) ([]*models.Information, error) {
	var list []*models.Information
	if err := m.Engine.Where(condition).Limit(size, offset).OrderBy("sortorder DESC, id DESC").Find(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// 获取资讯总数
func (m *InformationModel) GetTotalNum() (int64, error) {
	return m.Engine.Where("status != 3").Count(&models.Information{})
}

func (m *InformationModel) AddInformation(information *models.Information) (int64, error) {
	return m.Engine.InsertOne(information)
}

// 资讯统计初始化
func (m *InformationModel) AddInformationStatistic() (int64, error) {
	return m.Engine.InsertOne(m.Statistic)
}

// 获取资讯统计数据
func (m *InformationModel) GetInformationStatistic(newsId string) (bool, error) {
	m.Statistic = new(models.InformationStatistic)
	return m.Engine.Where("news_id=?", newsId).Get(m.Statistic)
}

// id获取资讯信息
func (m *InformationModel) GetInformationById(id string) (bool, error) {
	m.Information = new(models.Information)
	return m.Engine.Where("id=?", id).Get(m.Information)
}

// 获取用户浏览过的资讯
func (m *InformationModel) GetUserBrowseInformation(userId string, composeType int, composeId int64) *models.UserBrowseRecord {
	m.Browse = new(models.UserBrowseRecord)
	ok, err := m.Engine.Where("user_id=? AND compose_type=? AND compose_id=?", userId, composeType, composeId).Get(m.Browse)
	if !ok || err != nil {
		return nil
	}

	return m.Browse
}

// 记录用户浏览的资讯记录
func (m *InformationModel) RecordUserBrowseRecord() error {
	if _, err := m.Engine.InsertOne(m.Browse); err != nil {
		return err
	}

	return nil
}

// 之前有浏览记录 更新浏览时间
func (m *InformationModel) UpdateUserBrowseInformation(userId string, composeType int, composeId int64) error {
	if _, err := m.Engine.Where("user_id=? AND compose_id=? AND compose_type=?", userId, composeId, composeType).Cols("create_at, update_at").Update(m.Browse); err != nil {
		return err
	}

	return nil
}

const (
	UPDATE_INFORMATION_LIKE_NUM = "UPDATE `information_statistic` SET `fabulous_num` = `fabulous_num` + ?, " +
		"`heat_num` = `heat_num` + ?, `update_at`=? WHERE `news_id`=? AND `fabulous_num` + ? >= 0 LIMIT 1"
)

// 更新资讯点赞数
func (m *InformationModel) UpdateInformationLikeNum(newsId int64, now, num, score int) error {
	if _, err := m.Engine.Exec(UPDATE_INFORMATION_LIKE_NUM, num, score, now, newsId, num); err != nil {
		return err
	}

	return nil
}

const (
	UPDATE_INFORMATION_COLLECT_NUM = "UPDATE `information_statistic` SET `collect_num` = `collect_num` + ?, `update_at`=? WHERE `news_id`=? AND `collect_num` + ? >= 0 LIMIT 1"
)

// 更新资讯收藏数
func (m *InformationModel) UpdateInformationCollectNum(newsId int64, now, num int) error {
	if _, err := m.Engine.Exec(UPDATE_INFORMATION_COLLECT_NUM, num, now, newsId, num); err != nil {
		return err
	}

	return nil
}

const (
	UPDATE_INFORMATION_COMMENT_NUM = "UPDATE `information_statistic` SET `comment_num` = `comment_num` + ?, " +
		"`heat_num` = `heat_num` + ?, `update_at`=? WHERE `news_id`=? AND `comment_num` + ? >= 0 LIMIT 1"
)

// 更新资讯评论数
func (m *InformationModel) UpdateInformationCommentNum(newsId int64, now, num, score int) error {
	if _, err := m.Engine.Exec(UPDATE_INFORMATION_COMMENT_NUM, num, score, now, newsId, num); err != nil {
		return err
	}

	return nil
}

const (
	UPDATE_INFORMATION_BROWSE_NUM = "UPDATE `information_statistic` SET `browse_num` = `browse_num` + ?, `heat_num` = `heat_num` + ?," +
		" `update_at`=? WHERE `news_id`=? AND `browse_num` + ? >= 0 LIMIT 1"
)

// 更新资讯浏览数
func (m *InformationModel) UpdateInformationBrowseNum(newsId int64, now, num, score int) error {
	if _, err := m.Engine.Exec(UPDATE_INFORMATION_BROWSE_NUM, num, score, now, newsId, num); err != nil {
		return err
	}

	return nil
}

const (
	UPDATE_INFORMATION_SHARE_NUM = "UPDATE `information_statistic` SET `share_num` = `share_num` + ?, `heat_num` = `heat_num` + ?, " +
		"`update_at`=? WHERE `news_id`=? AND `share_num` + ? >= 0 LIMIT 1"
)

// 更新资讯分享数
func (m *InformationModel) UpdateInformationShareNum(newsId int64, now, num, score int) error {
	if _, err := m.Engine.Exec(UPDATE_INFORMATION_SHARE_NUM, num, score, now, newsId, num); err != nil {
		return err
	}

	return nil
}

// 获取用户发布的资讯数
func (m *InformationModel) GetTotalPublish(userId string) int64 {
	total, err := m.Engine.Where("user_id=? AND status=1", userId).Count(&models.Information{})
	if err != nil {
		return 0
	}

	return total
}
