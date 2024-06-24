package init_application

import (
	"aiwechat/application/utils"
	"github.com/gorilla/websocket"
	"net/http"
)

func WSInit() {
	utils.UPGRADER = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许任何源的连接
		},
	}
}
