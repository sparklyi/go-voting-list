package router

import (
	"gin_Ranking/config"
	"gin_Ranking/controller"
	"gin_Ranking/pkg/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp", config.RedisAddr, config.RedisPasswd, []byte("secret"))

	//挂载中间件
	r.Use(sessions.Sessions("loginStore", store))

	//todo 请求主机信息记录
	r.Use(gin.LoggerWithConfig(logger.ToFile()))
	//todo panic信息记录
	r.Use(logger.Recover)

	//用户请求组
	user := r.Group("/user")
	{
		user.POST("/register", controller.UserController{}.Register)
		user.POST("/login", controller.UserController{}.Login)

	}
	//参赛人员请求组
	player := r.Group("/player")
	{
		player.POST("/list", controller.PlayerController{}.GetPlayerInfo)
	}
	//投票请求组
	vote := r.Group("/vote")
	{
		vote.POST("/add", controller.VoteController{}.AddVote)
		vote.POST("/ranking", controller.VoteController{}.GetVoteRanking)

	}
	//活动请求组
	activity := r.Group("/activity")
	{
		activity.POST("/createAct", controller.ActController{}.CreateActivity)
	}

	return r
}
