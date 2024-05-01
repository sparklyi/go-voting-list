package controller

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonStruct struct {
	Code  int `json:"code"`
	Msg   any `json:"msg"`
	Data  any `json:"data"`
	Count int `json:"count"`
}

func Success(c *gin.Context, code int, msg, data any, count int) {
	json := &JsonStruct{Code: code, Msg: msg, Data: data, Count: count}
	c.JSON(http.StatusOK, json)
}

func Failed(c *gin.Context, code int, msg any) {
	json := &JsonStruct{Code: code, Msg: msg}
	c.JSON(http.StatusOK, json)
}

// EncryptMD5 返回字符串的MD5哈希值
func EncryptMD5(s string) string {

	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}
