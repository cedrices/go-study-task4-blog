package main

import (
	"blog/auth"
	"blog/db/mysql"
	"blog/handler"
	"blog/service"
	"blog/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config, err := util.LoadConfig()
	port := ":8081"

	router.Use(auth.AuthFunc())
	if err == nil {
		port = fmt.Sprintf(":%d", config.Server.Port)
		db, errs := mysql.InitDB(config)
		if errs == nil {
			api := router.Group("v1")
			//初始化表
			service.InitTable(db)
			//初始化controller
			handler.InitController(api)
		}
	}
	router.Run(port)
}
