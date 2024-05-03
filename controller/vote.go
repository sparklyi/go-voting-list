package controller

import (
	"gin_Ranking/cache"
	"gin_Ranking/models"
	"gin_Ranking/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
	"time"
)

type VoteController struct {
}

func (VoteController) AddVote(c *gin.Context) {
	//获取当前活动,用户，参赛者
	actID, _ := strconv.Atoi(c.DefaultPostForm("actId", "0"))
	userID, _ := strconv.Atoi(c.DefaultPostForm("userId", "0"))
	playerID, _ := strconv.Atoi(c.DefaultPostForm("playerId", "0"))

	//不为空
	if userID == 0 || playerID == 0 || actID == 0 {
		Failed(c, http.StatusBadRequest, "data uncompleted")
		return
	}
	//活动必然存在，逻辑是前端点击了活动才会有投票操作
	//actRecord, _ := models.GetActToID(actID)
	//用户不存在
	userRecord, _ := models.GetUserInfoByID(userID)
	if userRecord.ID == 0 {
		Failed(c, http.StatusNotFound, "not found user")
		return
	}
	//选手不存在
	playerRecord, _ := models.GetPlayerInfoByID(playerID)
	if playerRecord.ID == 0 {
		Failed(c, http.StatusNotFound, "not found player")
		return
	}
	//查询剩余票数,此处默认只投一票
	voteRecord, _ := models.GetVoteInfo(actID, userID, playerID)
	if voteRecord.ID != 0 {
		Failed(c, http.StatusBadRequest, "have voted")
		return
	}
	//创建投票记录
	voteMsg, voteErr := models.AddVote(actID, userID, playerID)
	if voteErr != nil {
		logger.Error(map[string]interface{}{"error": "create vote record failed"}, voteErr.Error())
		Failed(c, http.StatusInternalServerError, "create vote record failed")
		return
	}
	//更新选手得票数,同时更新redis中的数据(活动在redis存在)
	pollErr := models.UpdatePlayerPoll(playerID)
	//判断活动存在于redis中
	if pollErr != nil {
		logger.Error(map[string]interface{}{"error:": "update player poll  failed"}, pollErr.Error())
		Failed(c, http.StatusInternalServerError, "update player poll  failed")
		return
	}
	redisKey := "ranking:" + strconv.Itoa(actID)
	isKey, keyErr := cache.Rdb.Exists(cache.Rctx, redisKey).Result()
	if keyErr != nil {
		logger.Error(map[string]interface{}{"error:": "/controller/vote.go:64"}, pollErr.Error())
		Failed(c, http.StatusInternalServerError, "internal server error ")
		return
	}
	//返回的是键的数量
	if isKey > 0 {
		//更新当前键值的playerID的投票数
		cache.Rdb.ZIncrBy(cache.Rctx, redisKey, 1, strconv.Itoa(playerID))
	}
	Success(c, http.StatusOK, "add vote successfully", voteMsg, 1)
}

// GetVoteRanking 获取活动排行榜
func (VoteController) GetVoteRanking(c *gin.Context) {
	//获取当前活动id
	actId, _ := strconv.Atoi(c.DefaultPostForm("actId", "0"))
	if actId == 0 {
		Failed(c, http.StatusNotFound, "activity does not exist")
		return
	}
	//redis中存储的键值
	redisKey := "ranking:" + strconv.Itoa(actId)
	//获取有序集合(zset)中键值为redisKey的集合，ZRevRange 当键值不存在时会报错，ZRange不会
	zSet, zErr := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 0, -1).Result()
	//集合存在且不会空,利用有序集合直接获取数据，不使用sql的排序
	if zErr == nil && len(zSet) > 0 {
		var players []models.Player
		for _, v := range zSet {
			tid, _ := strconv.Atoi(v)
			tRecord, _ := models.GetPlayerInfoByID(tid)
			if tRecord.ID > 0 {
				players = append(players, tRecord)
			}
		}
		Success(c, http.StatusOK, "successfully", players, 1)
		return
	}

	//查询活动id并排序分数
	players, err := models.GetPlayerInfoByActID(actId, "poll desc")
	if err != nil {
		logger.Error(map[string]interface{}{"error": "sql error"}, err.Error())
		Failed(c, http.StatusInternalServerError, "query error")
		return
	}
	//将数据存入redis 下次查询直接使用redis

	for _, v := range players {
		addErr := cache.Rdb.ZAdd(cache.Rctx, redisKey, redis.Z{Member: v.ID, Score: float64(v.Poll)}).Err()
		if addErr != nil {
			logger.Error(map[string]interface{}{"error": "zset add error"}, addErr.Error())
			Failed(c, http.StatusInternalServerError, "internal server error")
			return
		}
		cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Hour)

	}
	Success(c, http.StatusOK, "successfully", players, 1)

}
