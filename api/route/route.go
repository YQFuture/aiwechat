package route

import (
	api2 "aiwechat/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route(engine *gin.Engine) {
	aiwechat := engine.Group("/aiwechat")
	{
		//websocket接口
		aiwechat.GET("/conn", api2.Conn)
		//前端页面
		aiwechat.StaticFS("/page", http.Dir("./dist"))
	}
}
