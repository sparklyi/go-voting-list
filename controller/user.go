package controller

import (
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

	//通过postform获得
	//acc.Name = c.PostForm("name")
	//acc.Msg = c.PostForm("msg")
	if err != nil {
		Failed(c, http.StatusNotFound, "no info")
	}
	Success(c, http.StatusOK, "OK", acc, 1)

}

func (user UserController) GetList(c *gin.Context) {
	Failed(c, http.StatusBadRequest, "no info")
}
