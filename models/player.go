package models

import (
	"gin_Ranking/dao"
	"gin_Ranking/pkg/logger"
)

type Player struct {
}

func (Player) TableName() string {
	return "player"
}

func init() {

	//自动迁移
	err := dao.DB.AutoMigrate(&Player{})
	if err != nil {
		logger.Error(map[string]interface{}{"error": "player table autoMigrate failed"}, err.Error())
	}
}
