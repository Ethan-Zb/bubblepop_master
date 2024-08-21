package main

import (
	"bubblepop_master/core"
	"bubblepop_master/global"
	"bubblepop_master/router"
	"fmt"
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
	addr := global.Config.System.Addr()
	r.Run(addr)

	//gormInit := core.Gorm()
	//fmt.Println(gormInit.Config)
}
