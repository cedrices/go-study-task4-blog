package mysql

import (
	"blog/util"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
)

var (
	db   *gorm.DB
	err  error
	once sync.Once
)

func InitDB(conf *util.Config) (*gorm.DB, error) {
	once.Do(func() {
		if conf == nil {
			log.Printf("configuration is nil")
		}

		// Validate required fields
		if conf.DB.Username == "" || conf.DB.Password == "" || conf.DB.Host == "" || conf.DB.Port == 0 || conf.DB.Dbname == "" {
			log.Printf("missing required database configuration fields")
		}

		// Database initialization logic here
		str := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
		dsn := fmt.Sprintf(str, conf.DB.Username, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.Dbname)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Printf("failed to connect database")
		}
		log.Printf("数据库连接成功", db)
	})
	return db, err
}
