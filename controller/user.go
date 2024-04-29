package controller

import (
	"gin_Ranking/models"
	"gin_Ranking/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Account struct {
	Name string `json:"name"`
	Msg  any    `json:"msg"`
}
type UserController struct {
}

func (user UserController) GetInfo(c *gin.Context) {
	//通过shouldBind绑定字段,此时设置了json 发送时只能解析json格式
	acc := Account{}
	err := c.ShouldBind(&acc)
	if err != nil {
		Failed(c, http.StatusNotFound, "no info")
	}
	//通过postform获得
	//acc.Name = c.PostForm("name")
	//acc.Msg = c.PostForm("msg")
	err = models.CreateUserInfo(acc.Name)

	if err != nil {
		logger.Error(map[string]interface{}{"error": "create field failed"}, err.Error())
	}
	record, err := models.GetUserInfo(acc.Name)
	if err != nil {
		logger.Error(map[string]interface{}{"getUserTest failed": err.Error()})
	}
	Success(c, http.StatusOK, acc.Name, record, 1)

}

func (user UserController) GetList(c *gin.Context) {
	//logger.Info(logrus.Fields{"test": 1}, "test")
}
