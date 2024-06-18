package route

import (
	api2 "aiwechat/api"
	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {
	aiwechat := engine.Group("/aiwechat")
	{
		aiwechat.GET("/conn", api2.Conn)
	}
}
