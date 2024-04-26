package router

import (
	"gin_Ranking/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {

	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/info", controller.UserController{}.GetInfo)

		user.GET("/list", controller.UserController{}.GetList)
	}
	order := r.Group("order")
	{
		order.GET("/list", controller.OrderController{}.GetList)
	}
	g := r.Group("/index")
	{
		g.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "get")
		})

		g.POST("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "post request")
		})

		g.PUT("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "put request")

		})

		g.DELETE("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "delete request")
		})

	}
	return r
}
