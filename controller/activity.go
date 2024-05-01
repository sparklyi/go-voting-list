package controller

import (
	"fmt"
	"gin_Ranking/models"
	"gin_Ranking/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ActController struct {
}

func (act ActController) CreateActivity(c *gin.Context) {
	//获取form
	name := c.PostForm("name")
	details := c.PostForm("details")

	err := models.CreateAct(name, details)
	if err != nil {
		logger.Error(map[string]interface{}{"error": "create activity record failed"}, err.Error())
		Failed(c, http.StatusInternalServerError, "创建失败")
		return
	}
	Success(c, http.StatusOK, "create", "OK", 1)

	acts, err := models.ReadActToName(name)
	if err != nil {
		logger.Error(map[string]interface{}{"error": "read act failed"}, err.Error())
		Failed(c, http.StatusInternalServerError, "读取失败")
		return
	}
	for _, v := range acts {
		fmt.Println(v)
	}
	Success(c, http.StatusOK, "read", "OK", 1)

}
