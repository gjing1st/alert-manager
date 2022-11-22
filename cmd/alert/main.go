package main

import (
	"alert-manager/internal/apiserver"
	"alert-manager/internal/apiserver/config"
	"alert-manager/internal/apiserver/service"
	"alert-manager/internal/apiserver/store"
	"alert-manager/internal/apiserver/store/database"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	//加载配置文件
	config.InitConfig()
	//加载数据驱动并初始化
	store.DB = database.GetDB()
	if store.DB != nil {
		db, _ := store.DB.DB()
		// 程序结束前关闭数据库链接
		defer db.Close()
	}
	//开启定时任务
	service.AddCron()
	//开启http服务
	apiserver.HttpStart()

}
