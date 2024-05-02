package models

import (
	"gin_Ranking/dao"
	"gin_Ranking/pkg/logger"
	"gorm.io/gorm"
)

type Player struct {
	ID          int    `gorm:"column:id" json:"id"`
	ActID       int    `gorm:"column:act_id" json:"actId"`
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

// GetPlayerInfoByActID 通过活动id查询选手
func GetPlayerInfoByActID(actID int) ([]Player, error) {
	var records []Player
	err := dao.DB.Model(&Player{}).Where("act_id = ?", actID).Find(&records).Error

	return records, err
}

// GetPlayerInfoByID 通过选手id查询选手
func GetPlayerInfoByID(id int) (Player, error) {

	var record Player
	err := dao.DB.Model(&Player{}).Where("id = ?", id).First(&record).Error
	return record, err

}

// UpdatePlayerPoll 更新选手的票数
func UpdatePlayerPoll(id int) error {
	//gorm.Expr 将表达式和参数拼接成子串返回
	//UPDATE `player` SET `poll`=poll + 1 WHERE id = 1
	err := dao.DB.Debug().Model(&Player{}).Where("id = ?", id).Update("poll", gorm.Expr("poll + ?", 1)).Error
	return err
}
