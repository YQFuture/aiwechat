package application

import "aiwechat/application/init_application"

func RUN() {
	defer init_application.RouteInit()
}
