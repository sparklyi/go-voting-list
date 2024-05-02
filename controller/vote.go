package controller

import (
	"gin_Ranking/models"
	"gin_Ranking/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VoteController struct {
}

func (VoteController) AddVote(c *gin.Context) {
	//获取当前用户，参赛者
	userID, _ := strconv.Atoi(c.DefaultPostForm("userId", "0"))
	playerID, _ := strconv.Atoi(c.DefaultPostForm("playerId", "0"))

	//不为空
	if userID == 0 || playerID == 0 {
		Failed(c, http.StatusBadRequest, "data uncompleted")
		return
	}
	//用户不存在
	userRecord, _ := models.GetUserInfoByID(userID)
	if userRecord.ID == 0 {
		Failed(c, http.StatusNotFound, "not found user")
		return
	}
	//选手不存在
	playerRecord, _ := models.GetPlayerInfoByID(playerID)
	if playerRecord.ID == 0 {
		Failed(c, http.StatusNotFound, "not found player")
		return
	}
	//查询剩余票数,此处默认只投一票
	voteRecord, _ := models.GetVoteInfo(userID, playerID)
	if voteRecord.ID != 0 {
		Failed(c, http.StatusBadRequest, "have voted")
		return
	}
	//创建投票记录
	voteMsg, err := models.AddVote(userID, playerID)
	if err != nil {
		logger.Error(map[string]interface{}{"error:": "create vote record failed"}, err.Error())
		Failed(c, http.StatusInternalServerError, "create vote record failed")
		return
	}
	//更新选手得票数
	err = models.UpdatePlayerPoll(playerID)
	if err != nil {
		logger.Error(map[string]interface{}{"error:": "update player poll  failed"}, err.Error())
		Failed(c, http.StatusInternalServerError, "update player poll  failed")
		return
	}
	Success(c, http.StatusOK, "add vote successfully", voteMsg, 1)
}
