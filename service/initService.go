package service

import (
	"blog/db/mysql"
	"blog/model"
	"blog/util"
	"gorm.io/gorm"
)

// 初始化表
func InitTable(db *gorm.DB) {
	//自动创建表
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
}

// 加载配置文件，初始化数据库
func loadConfGetDb() *gorm.DB {
	conf, err := util.LoadConfig()
	if err == nil {
		db, err := mysql.InitDB(conf)
		if err == nil {
			return db
		}
	}
	return nil
}
