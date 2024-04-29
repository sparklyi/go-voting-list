package models

import (
	"gin_Ranking/dao"
	"gin_Ranking/pkg/logger"
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(255)"`

	Details string `gorm:"column:details;type:varchar(255)"`
}

func (Activity) TableName() string {
	return "activity"
}

func init() {

	//自动迁移
	err := dao.DB.AutoMigrate(&Activity{})
	if err != nil {
		logger.Error(map[string]interface{}{"error": "activity table autoMigrate failed"}, err.Error())
	}

}

//todo curd操作

// CreateAct 创建活动
func CreateAct(name string, details string) error {
	//创建记录对象
	record := &Activity{Name: name, Details: details}
	err := dao.DB.Create(&record).Error
	return err
}

// ReadActToName 通过活动名读取活动信息
func ReadActToName(name string) (*[]Activity, error) {

	var records []Activity
	err := dao.DB.Model(&Activity{}).Where("name = ?", name).Find(&records).Error
	return &records, err
}

// ReadActToID 通过ID读取活动信息
func ReadActToID(id int) (*Activity, error) {

	var record Activity
	err := dao.DB.Model(&Activity{}).Where("id = ?", id).First(&record).Error
	return &record, err
}

//
