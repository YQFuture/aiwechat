package handle

import (
	"aiwechat/model"
	"encoding/json"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
	"io"
	"log"
)

func ConnHandler(ws *websocket.Conn, bot *openwechat.Bot) {
	autoLogout(ws, bot)
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
		go MessageDispatcher(ws, bot, &messageModel)
	}
}

// 自动登出
func autoLogout(ws *websocket.Conn, bot *openwechat.Bot) {
	ws.SetCloseHandler(func(code int, text string) error {
		log.Println("客户端断开ws连接")
		//bot登出
		err := bot.Logout()
		if err != nil && err != io.EOF {
			log.Println("bot登出失败", err)
			return err
		}
		return nil
	})
}
