package contest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sports_service/server/models"
	"sports_service/server/global/backend/errdef"
	"sports_service/server/global/backend/log"
	"sports_service/server/backend/controller/contest"
	"sports_service/server/util"
)

func AddPlayer(c *gin.Context) {
	reply := errdef.New(c)
	param := &models.FpvContestPlayerInformation{}
	if err := c.BindJSON(param); err != nil {
		log.Log.Errorf("contest_trace: add player param fail, param:%+v, err:%s", param, err)
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := contest.New(c)
	reply.Response(http.StatusOK, svc.AddPlayer(param))
}

func EditPlayer(c *gin.Context) {
	reply := errdef.New(c)
	param := &models.FpvContestPlayerInformation{}
	if err := c.BindJSON(param); err != nil {
		log.Log.Errorf("contest_trace: add player param fail, param:%+v, err:%s", param, err)
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := contest.New(c)
	reply.Response(http.StatusOK, svc.UpdatePlayer(param))
}

func PlayerList(c *gin.Context) {
	reply := errdef.New(c)
	page, size := util.PageInfo(c.Query("page"), c.Query("size"))

	svc := contest.New(c)
	code, list := svc.GetPlayerList(page, size)
	reply.Data["list"] = list
	reply.Response(http.StatusOK, code)
}

func AddContestGroup(c *gin.Context) {
	reply := errdef.New(c)
	param := &models.FpvContestScheduleGroup{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := contest.New(c)
	reply.Response(http.StatusOK, svc.AddContestGroup(param))
}

func EditContestGroup(c *gin.Context) {
	reply := errdef.New(c)
	param := &models.FpvContestScheduleGroup{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := contest.New(c)
	reply.Response(http.StatusOK, svc.EditContestGroup(param))
}

func ContestGroupList(c *gin.Context) {
	reply := errdef.New(c)
	page, size := util.PageInfo(c.Query("page"), c.Query("size"))
	scheduleId := c.Query("schedule_id")
	contestId := c.Query("contest_id")

	svc := contest.New(c)
	code, list := svc.GetContestGroupList(page, size, scheduleId, contestId)
	reply.Data["list"] = list
	reply.Response(http.StatusOK, code)
}

func ContestSchedule(c *gin.Context) {
	reply := errdef.New(c)
	svc := contest.New(c)
	code, list := svc.GetContestScheduleInfo()
	reply.Data["list"] = list
	reply.Response(http.StatusOK, code)
}

func SetIntegralRanking(c *gin.Context) {
	reply := errdef.New(c)
	param := &models.FpvContestPlayerIntegralRanking{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := contest.New(c)
	reply.Response(http.StatusOK, svc.SetIntegralRanking(param))
}

func EditIntegralRanking(c *gin.Context) {
	reply := errdef.New(c)
	param := &models.FpvContestPlayerIntegralRanking{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := contest.New(c)
	reply.Response(http.StatusOK, svc.UpdateIntegralRanking(param))
}

// 赛事积分排行
func IntegralRankingList(c *gin.Context) {
	reply := errdef.New(c)
	page, size := util.PageInfo(c.Query("page"), c.Query("size"))
	svc := contest.New(c)
	code, list := svc.GetIntegralRankingList(page, size)
	reply.Data["list"] = list
	reply.Response(http.StatusOK, code)
}
