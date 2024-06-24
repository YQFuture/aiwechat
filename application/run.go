package application

import (
	"aiwechat/application/init_application"
	"aiwechat/application/utils"
)

func RUN() {
	// 最后启动路由监听
	defer init_application.RouteInit()
	// 初始化websocket
	init_application.WSInit()
	// 初始化日志框架
	init_application.LogrusInit()
	utils.Logger.Infoln("aiwechat start")
}
