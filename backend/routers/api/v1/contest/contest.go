package contest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sports_service/backend/controller/contest"
	"sports_service/global/backend/errdef"
	"sports_service/global/backend/log"
	"sports_service/models"
	"sports_service/models/mcontest"
	"sports_service/util"
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
	code, list, count := svc.GetPlayerList(page, size)
	reply.Data["list"] = list
	reply.Data["total"] = count
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
	status := c.Query("status")

	svc := contest.New(c)
	code, list := svc.GetContestGroupList(page, size, scheduleId, contestId, status)
	reply.Data["list"] = list
	reply.Data["total"] = svc.GetContestGroupCount(scheduleId, contestId, status)
	reply.Response(http.StatusOK, code)
}

func ContestSchedule(c *gin.Context) {
	reply := errdef.New(c)
	svc := contest.New(c)
	code, list := svc.GetContestScheduleInfo()
	reply.Data["list"] = list
	reply.Response(http.StatusOK, code)
}

func AddContestScheduleDetail(c *gin.Context) {
	reply := errdef.New(c)
	param := &mcontest.AddScheduleDetail{}
	if err := c.BindJSON(param); err != nil {
		log.Log.Errorf("contest_trace: invalid param, param:%+v, err:%s", param, err)
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := contest.New(c)
	reply.Response(http.StatusOK, svc.AddContestScheduleDetail(param))
}

func ContestScheduleDetailList(c *gin.Context) {
	reply := errdef.New(c)
	scheduleId := c.Query("schedule_id")
	svc := contest.New(c)
	code, list := svc.GetContestScheduleDetailList(scheduleId)
	reply.Data["list"] = list
	reply.Response(http.StatusOK, code)
}

func DelScheduleDetail(c *gin.Context) {
	reply := errdef.New(c)
	ids := c.Query("ids")
	var infoIds []int
	if err := util.JsonFast.UnmarshalFromString(ids, &infoIds); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := contest.New(c)
	reply.Response(http.StatusOK, svc.DelScheduleDetail(infoIds))
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
	reply.Data["total"] = svc.GetIntegralRankingTotal()
	reply.Response(http.StatusOK, code)
}

func AddContestLive(c *gin.Context) {
	reply := errdef.New(c)
	param := &models.VideoLive{}
	if err := c.BindJSON(param); err != nil {
		log.Log.Errorf("contest_trace: invalid param, err:%s", err)
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := contest.New(c)
	reply.Response(http.StatusOK, svc.AddContestLive(param))
}

func UpdateContestLive(c *gin.Context) {
	reply := errdef.New(c)
	param := &models.VideoLive{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := contest.New(c)
	reply.Response(http.StatusOK, svc.UpdateContestLive(param))
}

func DelContestLive(c *gin.Context) {
	reply := errdef.New(c)
	id := c.Query("id")

	svc := contest.New(c)
	reply.Response(http.StatusOK, svc.DelContestLive(id))
}

func ContestLiveList(c *gin.Context) {
	reply := errdef.New(c)
	page, size := util.PageInfo(c.Query("page"), c.Query("size"))
	svc := contest.New(c)
	code, list := svc.GetContestLiveList(page, size)
	reply.Data["list"] = list
	reply.Data["total"] = svc.GetContestLiveCount()
	reply.Response(http.StatusOK, code)
}

func AddLiveData(c *gin.Context) {
	reply := errdef.New(c)
	param := &mcontest.AddLiveDataParam{}
	if err := c.BindJSON(param); err != nil {
		reply.Response(http.StatusOK, errdef.INVALID_PARAMS)
		return
	}

	svc := contest.New(c)
	reply.Response(http.StatusOK, svc.AddLiveData(param))
}

func LiveDataList(c *gin.Context) {
	reply := errdef.New(c)
	svc := contest.New(c)
	code, list := svc.GetContestLiveData(c.Query("live_id"))
	reply.Data["list"] = list
	reply.Response(http.StatusOK, code)
}
