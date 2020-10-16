package mbarrage

import (
	"sports_service/server/global/app/log"
	"sports_service/server/models"
	"github.com/go-xorm/xorm"
)

type BarrageModel struct {
	Barrage      *models.VideoBarrage
	Engine       *xorm.Session
}

// 发送弹幕请求参数
type SendBarrageParams struct {
	Color            string `json:"color"`
	Content          string `binding:"required" json:"content"`
	Font             string `json:"font"`
	Location         int    `json:"location"`
	VideoCurDuration int    `binding:"required" json:"video_cur_duration"`
	VideoId          int64  `binding:"required" json:"video_id"`
}

// 实栗
func NewBarrageModel(engine *xorm.Session) *BarrageModel {
	return &BarrageModel{
		Engine: engine,
		Barrage: new(models.VideoBarrage),
	}
}

// 记录视频弹幕
func (m *BarrageModel) RecordVideoBarrage() error {
	if _, err := m.Engine.Insert(m.Barrage); err != nil {
		log.Log.Errorf("barrage_trace: record video barrage err:%s", err)
		return err
	}

	return nil
}

// 根据视频时长区间获取弹幕 todo:限制下 最多取最新的1000条 根据取的时间区间大小做调整
func (m *BarrageModel) GetBarrageByDuration(videoId, minDuration, maxDuration string, offset, limit int) []*models.VideoBarrage {
	var list []*models.VideoBarrage
	if err := m.Engine.Where(" video_id =? AND video_cur_duration >= ? AND video_cur_duration <= ?", videoId,
		minDuration, maxDuration).Desc("id").Limit(limit, offset).Find(&list); err != nil {
		return []*models.VideoBarrage{}
	}

	return list
}

// 获取用户视频总弹幕数
func (m *BarrageModel) GetUserTotalVideoBarrage(userId string) int64 {
  total, err := m.Engine.Where("user_id=?", userId).Count(m.Barrage)
  if err != nil {
    log.Log.Errorf("barrage_trace: get user total barrage err:%s, uid:%s", err, userId)
    return 0
  }

  return total
}

type VideoBarrageInfo struct {
  Id               int64  `json:"id"`
  VideoId          int64  `json:"video_id"`
  VideoCurDuration int    `json:"video_cur_duration"`
  Content          string `json:"content"`
  UserId           string `json:"user_id"`
  Color            string `json:"color"`
  Font             string `json:"font"`
  BarrageType      int    `json:"barrage_type"`
  Location         int    `json:"location"`
  SendTime         int64  `json:"send_time"`
  Title            string `json:"title"`
  VideoAddr        string `json:"video_addr"`
}

// 后台分页获取 视频弹幕 列表
func (m *BarrageModel) GetVideoBarrageList(offset, size int) []*VideoBarrageInfo {
  sql := "SELECT vb.*, v.title, v.video_addr FROM video_barrage AS vb LEFT JOIN videos AS v ON vb.video_id=v.video_id GROUP BY vb.id LIMIT ?, ?"
  var list []*VideoBarrageInfo
  if err := m.Engine.Table(&models.VideoComment{}).SQL(sql, offset, size).Find(&list); err != nil {
      log.Log.Errorf("barrage_trace: get video barrage list by sort, err:%s", err)
      return []*VideoBarrageInfo{}
  }

  return list
}

// 后台删除弹幕请求参数
type DelBarrageParam struct {
  Id      string     `binding:"required" json:"id"`       // 弹幕id
}
const (
  DELETE_VIDEO_BARRAGE = "DELETE FROM `video_barrage` WHERE `id`=?"
)
// 删除弹幕
func (m *BarrageModel) DelVideoBarrage(id string) error {
  if _, err := m.Engine.Exec(DELETE_VIDEO_BARRAGE, id); err != nil {
    return err
  }

  return nil
}

// 获取弹幕总数
func (m *BarrageModel) GetVideoBarrageTotal() int64 {
  count, err := m.Engine.Where("barrage_type=0").Count(&models.VideoBarrage{})
  if err != nil {
    log.Log.Errorf("comment_trace: get total barrage err:%s", err)
    return 0
  }

  return count
}



