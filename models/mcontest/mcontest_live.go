package mcontest

import (
	"sports_service/models"
	"sports_service/tools/tencentCloud"
)

type baseInfo struct {
	EventType   int    `json:"event_type"` // 1推流  0断流  100直播录制
	Appid       int    `json:"appid"`
	App         string `json:"app"`
	Appname     string `json:"appname"`
	StreamID    string `json:"stream_id"`
	ChannelID   string `json:"channel_id"`
	StreamParam string `json:"stream_param"`
	Sign        string `json:"sign"`
	T           int    `json:"t"`
}

type AddLiveDataParam struct {
	List []*models.FpvContestScheduleLiveData `json:"list"`
}

// 推流/断流/流录制 回调信息
type StreamCallbackInfo struct {
	baseInfo
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	EventTime    int    `json:"event_time"`
	SetID        int    `json:"set_id"`
	Node         string `json:"node"`
	Sequence     string `json:"sequence"`
	UserIP       string `json:"user_ip"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	PushDuration string `json:"push_duration"`
	VideoURL     string `json:"video_url"`
	FileID       string `json:"file_id"`
	FileFormat   string `json:"file_format"`
	TaskID       string `json:"task_id"`
	StartTime    int    `json:"start_time"` // 录制文件起始时间戳
	EndTime      int    `json:"end_time"`   // 录制文件结束时间戳
	Duration     int    `json:"duration"`   // 录制文件时长，单位秒
	FileSize     int    `json:"file_size"`  // 录制文件大小，单位字节
	Status       int    // 0未直播 1直播中 2 已结束
}

// 赛事直播信息
type ContestLiveInfo struct {
	Date    string `json:"date"`
	Week    string `json:"week"`
	Index   int    `json:"index"`
	IsToday bool   `json:"is_today"`

	LiveInfo []*LiveInfo `json:"live_info"`
}

type LiveInfo struct {
	Id         int64                  `json:"id"`
	UserId     string                 `json:"user_id"`
	RoomId     string                 `json:"room_id"`
	GroupId    string                 `json:"group_id"`
	Cover      tencentCloud.BucketURI `json:"cover"`
	RtmpAddr   string                 `json:"rtmp_addr"`
	FlvAddr    string                 `json:"flv_addr"`
	HlsAddr    string                 `json:"hls_addr"`
	PlayTime   int                    `json:"play_time"`
	Title      string                 `json:"title"`
	Subhead    string                 `json:"subhead"`
	HighLights string                 `json:"high_lights"`
	Describe   string                 `json:"describe"`
	Tags       string                 `json:"tags"`
	LiveType   int                    `json:"live_type"`
	NickName   string                 `json:"nick_name"`
	Avatar     tencentCloud.BucketURI `json:"avatar"`
	Status     int                    `json:"status"`     // 状态 0未直播 1直播中 2 已结束
	HasReplay  int                    `json:"has_replay"` // 是否有回放 1 有 2 无

	LiveReplayInfo *LiveReplayInfo `json:"live_replay_info"` // 直播回放信息
}

// 直播回放信息
type LiveReplayInfo struct {
	Id          int64  `json:"id"`
	UserId      string `json:"user_id"`
	LiveId      int64  `json:"live_id"`
	HistoryAddr string `json:"history_addr"`
	Title       string `json:"title"`
	Subhead     string `json:"subhead"`
	PlayNum     int64  `json:"play_num"`
	Duration    int    `json:"duration"`
	Size        int64  `json:"size"`
	Cover       string `json:"cover"`
	CreateAt    int    `json:"create_at"`
	Describe    string `json:"describe"`

	PlayInfo []*PlayInfo `json:"play_info"`
}

// 转码信息
type PlayInfo struct {
	Type     string `json:"type" example:"1 流畅（FLU） 2 标清（SD）3 高清（HD）4 全高清（FHD）5 2K 6 4K"` // 1 流畅（FLU） 2 标清（SD）3 高清（HD）4 全高清（FHD）5 2K 6 4K
	Url      string `json:"url" example:"对应类型的视频地址"`
	Size     int64  `json:"size" example:"1000000000"`
	Duration int64  `json:"duration" example:"1000000000"`
}

type VideoLive struct {
	Id            int64                  `json:"id" xorm:"pk autoincr comment('主键') BIGINT(20)"`
	UserId        string                 `json:"user_id" xorm:"not null default '' comment('主播id') VARCHAR(60)"`
	RoomId        string                 `json:"room_id" xorm:"not null default '' comment('房间id') index VARCHAR(60)"`
	GroupId       string                 `json:"group_id" xorm:"not null default '' comment('im群组id') VARCHAR(60)"`
	Cover         tencentCloud.BucketURI `json:"cover" xorm:"not null default '' comment('直播封面') VARCHAR(512)"`
	RtmpAddr      string                 `json:"rtmp_addr" xorm:"not null default '' comment('rtmp地址[拉流]') VARCHAR(512)"`
	FlvAddr       string                 `json:"flv_addr" xorm:"not null default '' comment('flv地址[拉流]') VARCHAR(512)"`
	HlsAddr       string                 `json:"hls_addr" xorm:"not null default '' comment('hls地址[拉流]') VARCHAR(512)"`
	PushStreamUrl string                 `json:"push_stream_url" xorm:"not null default '' comment('推流url') VARCHAR(255)"`
	PlayTime      int                    `json:"play_time" xorm:"not null default 0 comment('后台设置的赛事开播时间') INT(11)"`
	EndTime       int                    `json:"end_time" xorm:"not null default 0 comment('结束时间') INT(11)"`
	Status        int                    `json:"status" xorm:"default 0 comment('状态 0未直播 1直播中 2 已结束') TINYINT(1)"`
	Title         string                 `json:"title" xorm:"not null default '' comment('标题') VARCHAR(255)"`
	HighLights    string                 `json:"high_lights" xorm:"not null default '' comment('亮点') VARCHAR(255)"`
	Describe      string                 `json:"describe" xorm:"not null default '' comment('描述') VARCHAR(512)"`
	Tags          string                 `json:"tags" xorm:"not null default '' comment('直播标签') VARCHAR(512)"`
	CreateAt      int                    `json:"create_at" xorm:"not null default 0 comment('创建时间') INT(11)"`
	UpdateAt      int                    `json:"update_at" xorm:"not null default 0 comment('更新时间') INT(11)"`
	Duration      int64                  `json:"duration" xorm:"default 0 comment('时长（毫秒）') BIGINT(20)"`
	LiveType      int                    `json:"live_type" xorm:"not null default 0 comment('直播类型（0：管理员[sys_user]，1：用户[user]）') TINYINT(1)"`
	ContestId     int                    `json:"contest_id" xorm:"not null default 0 comment('赛事id') INT(11)"`
	ScheduleId    int                    `json:"schedule_id" xorm:"not null default 0 comment('赛程id') INT(11)"`
	StartTime     int                    `json:"start_time" xorm:"not null default 0 comment('真实开播时间') INT(11)"`
	PushStreamKey string                 `json:"push_stream_key" xorm:"not null default '' comment('推流密钥') VARCHAR(255)"`
	Subhead       string                 `json:"subhead" xorm:"not null default '' comment('副标题') VARCHAR(255)"`
}

// 获取直播列表
// queryType 1 首页 [只看最近同一天内的 未开播/直播中的数据] 2 赛程 [最近同一天内 所有状态的直播数据]
// pullType 拉取类型 1 上拉加载 今天及未来赛事数据 [通过开播时间作为查询条件进行拉取] 2 下拉加载 历史赛事数据 [通过开播时间作为查询条件进行拉取] 默认上拉加载
func (m *ContestModel) GetLiveList(offset, size int, contestId, tm, queryType, pullType string) ([]*models.VideoLive, error) {
	var list []*models.VideoLive
	m.Engine.Where("contest_id=?", contestId)
	if queryType == "1" {
		m.Engine.Where("play_time> ? AND status in(0, 1)", tm).Asc("play_time")
	}

	switch pullType {
	case "1":
		m.Engine.Where("play_time > ?", tm).Asc("play_time")
	case "2":
		m.Engine.Where("play_time < ?", tm).Desc("play_time")
	}

	if err := m.Engine.Limit(size, offset).Find(&list); err != nil {
		return []*models.VideoLive{}, err
	}

	return list, nil
}

// 通过房间id获取直播信息
func (m *ContestModel) GetLiveInfoByRoomId(roomId string) (bool, error) {
	m.VideoLive = new(models.VideoLive)
	return m.Engine.Where("room_id=?", roomId).Get(m.VideoLive)
}

// 更新直播数据
func (m *ContestModel) UpdateLiveInfo(cols string) (int64, error) {
	return m.Engine.ID(m.VideoLive.Id).Cols(cols).Update(m.VideoLive)
}

// 添加直播回放
func (m *ContestModel) AddVideoLiveReply() (int64, error) {
	return m.Engine.InsertOne(m.VideoLiveReplay)
}

// 获取直播回放[已上架未删除]
func (m *ContestModel) GetVideoLiveReply(liveId string) (bool, error) {
	m.VideoLiveReplay = new(models.VideoLiveReplay)
	return m.Engine.Where("live_id=? AND labeltype=1 AND is_del=0", liveId).Get(m.VideoLiveReplay)
}

// 获取赛事直播总数量
func (m *ContestModel) GetVideoLiveCount() (int64, error) {
	return m.Engine.Count(&models.VideoLive{})
}

// 获取直播信息
func (m *ContestModel) GetLiveInfoByCondition(condition string) (bool, error) {
	m.VideoLive = new(models.VideoLive)
	return m.Engine.Where(condition).Get(m.VideoLive)
}

// 通过腾讯云文件id 获取直播回放
func (m *ContestModel) GetVideoLiveReplyByFileId(fileId string) (bool, error) {
	m.VideoLiveReplay = new(models.VideoLiveReplay)
	return m.Engine.Where("file_id=?", fileId).Get(m.VideoLiveReplay)
}

// 更新直播回放数据
func (m *ContestModel) UpdateVideoLiveReplayInfo(condition, cols string) error {
	if _, err := m.Engine.Where(condition).Cols(cols).
		Update(m.VideoLiveReplay); err != nil {
		return err
	}

	return nil
}

// 添加赛事直播
func (m *ContestModel) AddContestLive(live *models.VideoLive) (int64, error) {
	return m.Engine.InsertOne(live)
}

// 更新赛事直播信息
func (m *ContestModel) UpdateContestLive(live *models.VideoLive) (int64, error) {
	return m.Engine.Where("id=?", live.Id).Update(live)
}

// 删除赛事直播
func (m *ContestModel) DelContestLive(id string) (int64, error) {
	return m.Engine.Where("id=?", id).Delete(&models.VideoLive{})
}

// 获取赛事直播列表
func (m *ContestModel) GetContestLiveList(offset, size int) ([]*VideoLive, error) {
	var list []*VideoLive
	if err := m.Engine.Limit(size, offset).Find(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// 获取直播数据总数
func (m *ContestModel) GetLiveCount() int64 {
	count, err := m.Engine.Count(&models.VideoLive{})
	if err != nil {
		return 0
	}

	return count
}

func (m *ContestModel) DelLiveDataById(liveId string) (int64, error) {
	return m.Engine.Where("live_id=?", liveId).Delete(&models.FpvContestScheduleLiveData{})
}

func (m *ContestModel) AddLiveData(list []*models.FpvContestScheduleLiveData) (int64, error) {
	return m.Engine.InsertMulti(list)
}

func (m *ContestModel) GetLiveDataById(liveId string) ([]*models.FpvContestScheduleLiveData, error) {
	var list []*models.FpvContestScheduleLiveData
	if err := m.Engine.Where("status=0 AND live_id=?", liveId).Asc("ranking").Find(&list); err != nil {
		return list, err
	}

	return list, nil

}
