package controller

import (
	"gin_Ranking/models"
	"gin_Ranking/pkg/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Account struct {
	Name string `json:"name"`
	Msg  any    `json:"msg"`
}
type UserController struct {
}

// Register 用户注册
func (u UserController) Register(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")

	//todo 不能为空
	if username == "" || password == "" || confirmPassword == "" {
		Failed(c, http.StatusBadRequest, "null data present")
		return
	}
	//todo 密码不一致
	if password != confirmPassword {
		Failed(c, http.StatusBadRequest, "twice password differ")
		return
	}
	//todo username存在
	record, _ := models.GetUserInfoByName(username)
	if record.ID != 0 {
		Failed(c, http.StatusBadRequest, "username is existed")
		return
	}
	err := models.CreateUserInfo(username, EncryptMD5(password))
	if err != nil {
		logger.Error(map[string]interface{}{"error": username + " register failed"}, err.Error())
		Failed(c, http.StatusInternalServerError, "fail to register")
		return
	}
	Success(c, http.StatusOK, "registered successfully", username, 1)
}

// Login 用户登录
func (u UserController) Login(c *gin.Context) {

	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	//todo 不为空
	if username == "" || password == "" {
		Failed(c, http.StatusBadRequest, "null data present")
		return
	}
	//todo 用户是否存在,存在时密码是否正确
	record, _ := models.GetUserInfoByName(username)
	if record.ID == 0 {
		Failed(c, http.StatusNotFound, "the user does not exist")
		return
	}
	//md5只能生成哈希字符串，所以利用当前输入密码加密后匹配record
	if record.Password != EncryptMD5(password) {
		Failed(c, http.StatusUnauthorized, "wrong password")
		return
	}
	//todo 保存session
	session := sessions.Default(c)
	session.Set("loginID:"+strconv.Itoa(record.ID), record.ID)
	err := session.Save()
	if err != nil {
		logger.Error(map[string]interface{}{"error": "session save failed"}, err.Error())
		return
	}

	data := models.UserAPI{ID: record.ID, Username: username}
	Success(c, http.StatusOK, "login successfully", data, 1)
}
