package handle

import (
	"aiwechat/model"
	"encoding/json"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
	"log"
)

func ConnHandler(ws *websocket.Conn, bot *openwechat.Bot) {
	for {
		// 读取消息
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Read:", err)
			break
		}

		// 解析成自定义的消息格式
		var messageModel model.RequestModel
		err = json.Unmarshal(message, &messageModel)

		//分发处理
		MessageDispatcher(ws, bot, &messageModel)

		// 发送消息回客户端
		err = ws.WriteMessage(websocket.TextMessage, append([]byte("Server Received: "), message...))
		if err != nil {
			log.Println("Write:", err)
			break
		}
	}
}
