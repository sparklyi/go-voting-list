package models

import (
	"gin_Ranking/dao"
	"gin_Ranking/pkg/logger"
)

type Player struct {
	ID          int    `gorm:"column:id" json:"id"`
	ActID       int    `gorm:"column:act_id" json:"actID"`
	Serial      string `gorm:"column:serial;type:varchar(255)" json:"serial"`
	Nickname    string `gorm:"column:nickname;type:varchar(255)" json:"nickname"`
	Declaration string `gorm:"column:declaration;type:varchar(255)" json:"declaration"`
	Avatar      string `gorm:"column:avatar;type:varchar(255)" json:"avatar"`
	Poll        int    `gorm:"column:poll" json:"poll"`
	CreateTime  int64  `gorm:"column:create_time" json:"createTime"`
	UpdateTime  int64  `gorm:"column:update_time" json:"updateTime"`
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
