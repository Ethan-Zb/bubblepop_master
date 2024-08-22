package main

import (
	"bubblepop_master/core"
	"bubblepop_master/global"
	"bubblepop_master/helper"
	"bubblepop_master/router"
	"fmt"
	"log"
)

func main() {
	// 读取配置文件
	core.InitConf()

	// 日志初始化
	global.Log = core.InitLogger()

	// db初始化
	global.DB = core.Gorm()
	fmt.Println(global.DB)

	// 路由初始化
	r := router.SetupRouters()

	// 鉴权
	r.Use(helper.AuthMiddleware(false, false))

	// 运行服务
	addr := global.Config.System.Addr()
	err := r.Run(addr)
	if err != nil {
		log.Fatalf("项目启动失败:%s", err)
	}

	//gormInit := core.Gorm()
	//fmt.Println(gormInit.Config)
}
