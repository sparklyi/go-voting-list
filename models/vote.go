package models

import (
	"gin_Ranking/dao"
	"gin_Ranking/pkg/logger"
	"time"
)

type Vote struct {
	ID       int   `gorm:"column:id" json:"id"`
	UserID   int   `gorm:"column:user_id;"`
	PlayerID int   `gorm:"column:player_id"`
	VoteTime int64 `gorm:"vote_time" json:"voteTime"`
}

func (Vote) TableName() string {
	return "vote"
}

func init() {

	//自动迁移
	err := dao.DB.AutoMigrate(&Vote{})
	if err != nil {
		logger.Error(map[string]interface{}{"error": "vote table autoMigrate failed"}, err.Error())
	}

}

func AddVote(userID, playerID int) (*Vote, error) {

	record := &Vote{UserID: userID, PlayerID: playerID, VoteTime: time.Now().Unix()}
	err := dao.DB.Create(&record).Error
	return record, err
}
