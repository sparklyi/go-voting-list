package models

import (
	"gin_Ranking/dao"
	"gin_Ranking/pkg/logger"
)

type Vote struct {
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
