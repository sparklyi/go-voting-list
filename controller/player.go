package controller

import (
	"gin_Ranking/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PlayerController struct {
}

func (PlayerController) GetPlayerInfo(c *gin.Context) {
	//获取当前活动的id
	actID, _ := strconv.Atoi(c.DefaultPostForm("actId", "0"))
	//查询此活动的参赛人员
	players, err := models.GetPlayerInfoByActID(actID, "id asc")
	if err != nil {
		Failed(c, http.StatusNotFound, "not found player of the activity")
		return
	}

	Success(c, http.StatusOK, "successfully", players, 1)
}
