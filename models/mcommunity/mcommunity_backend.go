package mcommunity

import "sports_service/server/models"

// 添加板块
type AddSection struct {
	Id          int    `json:"id"`
	SectionName string `json:"section_name"`
	Sortorder   int    `json:"sortorder"`
	Status      int    `json:"status"`
}

type DelSection struct {
	Id        int    `json:"id"`
}

type AddTopic struct {
	Name      string    `json:"name"`
	Sortorder int       `json:"sortorder"`
	Cover     string    `json:"cover"`
	Describe  string    `json:"describe"`
}

type DelTopic struct {
	Id       int    `json:"id"`
}

func (m *CommunityModel) TableName() string {
	return "community_section"
}

// 添加社区板块
func (m *CommunityModel) AddCommunitySection() (int64, error) {
	return m.Engine.InsertOne(m.CommunitySection)
}

// 删除社区板块
func (m *CommunityModel) DelCommunitySection(id int) (int64, error) {
	return m.Engine.Where("id=?", id).Delete(&models.CommunitySection{})
}

func (m *CommunityModel) UpdateSectionInfo(id int, mp map[string]interface{}) (int64, error) {
	return m.Engine.Table(m.TableName()).Where("id=?", id).Update(mp)
}

func (m *CommunityModel) AddTopic() (int64, error) {
	return m.Engine.InsertOne(m.CommunityTopic)
}

// 修改话题状态
func (m *CommunityModel) UpdateTopicStatus(id int) (int64, error) {
	return m.Engine.Where("id=?", id).Cols("status").Update(m.CommunityTopic)
}
