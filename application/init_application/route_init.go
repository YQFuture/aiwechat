package init_application

import (
	"aiwechat/api/route"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func RouteInit() {
	// 创建引擎
	engine := gin.Default()

	// 使用默认的CORS配置
	engine.Use(cors.Default())

	// 添加路由
	route.Route(engine)

	// 启动服务
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("Failed to listening and serving HTTP", err)
		return
	}
}
