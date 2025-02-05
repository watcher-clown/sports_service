package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"sports_service/app/controller/cvideo"
	"sports_service/dao"
	"sports_service/global/app/log"
	"sports_service/models"
	"sports_service/models/mlabel"
	"sports_service/models/mvideo"
	"sports_service/util"
	"strings"
	"time"
)

func init() {
	dao.AppEngine = dao.InitXorm("root:bluetrans888@tcp(192.168.5.12:3306)/sports_service?charset=utf8mb4", []string{"root:bluetrans888@tcp(192.168.5.12:3306)/sports_service?charset=utf8mb4"})
}

func main() {
	AddVideoLabels()
}

func AddVideoLabels() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	svc := cvideo.New(c)
	svc.GetVideoLabelList()

	session := dao.AppEngine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		log.Log.Errorf("job_trace: session begin err:%s", err)
		return
	}

	vmodel := mvideo.NewVideoModel(session)
	vlist := vmodel.GetVideoList(0, 1000)
	if vlist == nil {
		return
	}

	for _, video := range vlist {
		lmodel := mlabel.NewLabelModel(session)
		labelIds := strings.Split(fmt.Sprintf("%d,%d", util.GenerateRandnum(1, 10), util.GenerateRandnum(10, 17)), ",")
		// 组装多条记录 写入视频标签表
		labelInfos := make([]*models.VideoLabels, 0)
		for _, labelId := range labelIds {
			if lmodel.GetLabelInfoByMem(labelId) == nil {
				log.Log.Errorf("job_trace: label not found, labelId:%s", labelId)
				continue
			}

			info := new(models.VideoLabels)
			info.VideoId = video.VideoId
			info.LabelId = labelId
			info.LabelName = lmodel.GetLabelNameByMem(labelId)
			info.CreateAt = int(time.Now().Unix())
			info.Status = 1
			labelInfos = append(labelInfos, info)
		}

		if len(labelInfos) > 0 {
			// 添加视频标签（多条）
			affected, err := vmodel.AddVideoLabels(labelInfos)
			if err != nil || int(affected) != len(labelInfos) {
				log.Log.Errorf("job_trace: add video labels err:%s", err)
				session.Rollback()
				return
			}
		}

		session.Commit()

	}

}
