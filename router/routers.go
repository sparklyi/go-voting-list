package router

import (
	"gin_Ranking/controller"
	"gin_Ranking/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()
	//挂载中间件
	//todo 请求主机信息记录
	r.Use(gin.LoggerWithConfig(logger.ToFile()))

	//todo panic信息记录
	r.Use(logger.Recover)
	//用户请求组
	user := r.Group("/user")
	{
		user.POST("/info", controller.UserController{}.GetInfo)

		user.GET("/list", controller.UserController{}.GetList)
	}
	order := r.Group("order")
	{
		order.GET("/list", controller.OrderController{}.GetList)
	}

	//活动请求组
	activity := r.Group("/activity")
	{
		activity.POST("/createAct", controller.ActController{}.CreateActivity)
	}
	return r
}
