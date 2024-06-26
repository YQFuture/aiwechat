package route

import (
	api2 "aiwechat/api"
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed frontend
var staticFiles embed.FS

func Route(engine *gin.Engine) {

	// 将 embed.FS 类型转换为 http.FileSystem 接口
	fs := http.FS(staticFiles)

	// 提供 assets 目录中的静态文件
	engine.GET("/assets/*filepath", func(c *gin.Context) {
		c.FileFromFS("frontend/dist/assets/"+c.Param("filepath"), fs)
	})

	// 提供 favicon.ico
	engine.GET("/favicon.ico", func(c *gin.Context) {
		c.FileFromFS("frontend/dist/favicon.ico", fs)
	})

	// 静态web服务
	engine.GET("/frontend/*filepath", func(c *gin.Context) {
		c.FileFromFS("frontend/dist/"+c.Param("filepath"), fs)
	})

	aiwechat := engine.Group("/aiwechat")
	{
		//websocket接口
		aiwechat.GET("/conn", api2.Conn)
		//前端页面
		// aiwechat.StaticFS("/page", http.Dir("./dist"))
	}
}
