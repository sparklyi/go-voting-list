package models

import (
	"gin_Ranking/dao"
	"gin_Ranking/pkg/logger"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"username"`
}

// TableName 实现gorm中接口方法，数据库中为user表
func (User) TableName() string {
	return "user"
}

func init() {
	//自动迁移
	//表不存在时创建表结构
	err := dao.DB.AutoMigrate(&User{})
	if err != nil {
		logger.Error(map[string]interface{}{"error": "user table autoMigrate error"}, err.Error())
	}

}

// GetUserInfo 获取指定名字的用户信息
func GetUserInfo(username string) (User, error) {
	var recode User
	err := dao.DB.Model(&User{}).Where("username = ?", username).First(&recode).Error
	return recode, err
}

func CreateUserInfo(name string) error {
	//创建记录结构体
	record := &User{Model: gorm.Model{}, Username: name}
	//插入记录
	err := dao.DB.Create(&record).Error
	return err
}
