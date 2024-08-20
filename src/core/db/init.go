package db

import (
	"core/config"
	"fmt"

	coreLogger "core/logger"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	//"gorm.io/plugin/dbresolver"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

// InitDb 初始化数据库连接
func InitDb(config config.Datasource) *gorm.DB {
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.UserName,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.Charset)
	if db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		coreLogger.SugarLogger.Errorf(">>> 连接数据库错误:[%v:%v],错误原因: %s <<<", config.Host, config.Port, err.Error())
		panic("数据连接错误 info => " + err.Error())
	} else {
		coreLogger.SugarLogger.Infof(">>> 连接数据库:[%v:%v] <<<", config.Host, config.Port)
		DB = db
	}
	return DB
}

// GetDB 返回数据库实例
func GetDB() *gorm.DB {
	return DB
}
