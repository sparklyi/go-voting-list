package dao

import (
	"fmt"
	"gin_Ranking/pkg/logger"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	//todo 从ini文件中读取mysql相关内容
	//加载my.ini文件
	cnf, iniErr := ini.Load("config.ini")
	if iniErr != nil {
		logger.Error(map[string]interface{}{"read ini file failed": iniErr.Error()})
	}
	//读取mysql区块的数据
	pwd := cnf.Section("mysql").Key("password").String()
	address := cnf.Section("mysql").Key("address").String()
	port := cnf.Section("mysql").Key("port").String()
	db := cnf.Section("mysql").Key("db").String()

	dsn := fmt.Sprintf("root:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", pwd, address, port, db)

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect failed": err.Error()})
	}

}
