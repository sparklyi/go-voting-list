package dao

import (
	"gin_Ranking/config"
	"gin_Ranking/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	//连接数据库
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: config.MysqlDsn,
	}), &gorm.Config{})

	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect failed": err.Error()})
	}

}
