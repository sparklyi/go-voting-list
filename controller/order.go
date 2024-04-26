package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController struct {
}

func (order OrderController) GetList(c *gin.Context) {
	Failed(c, http.StatusNotFound, "order list")
}
