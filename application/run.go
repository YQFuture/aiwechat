package application

import (
	"aiwechat/application/init_application"
	"aiwechat/application/utils"
)

func RUN() {
	defer init_application.RouteInit()
	init_application.LogrusInit()
	utils.Logger.Infoln("aiwechat start")
}
