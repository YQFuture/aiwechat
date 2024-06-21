package utils

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var UPGRADER = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许任何源的连接
	},
}
