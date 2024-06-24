package handle

import (
	"aiwechat/application/utils"
	"aiwechat/handle/dispatcher"
	"aiwechat/model"
	"encoding/json"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
	"io"
)

func ConnHandler(ws *websocket.Conn, bot *openwechat.Bot) {
	autoLogout(ws, bot)
	for {
		// 读取消息
		_, message, err := ws.ReadMessage()
		if err != nil {
			utils.Logger.Errorln("读取请求失败", err)
			break
		}
		utils.Logger.Infoln("收到请求", string(message))

		// 解析成自定义的消息格式
		var messageModel model.RequestModel
		err = json.Unmarshal(message, &messageModel)
		if err != nil {
			utils.Logger.Errorln("解析请求失败", err)
			continue
		}

		utils.Logger.Infoln("解析请求结果", messageModel)

		//分发处理
		go dispatcher.MessageDispatcher(ws, bot, &messageModel)
	}
}

// 自动登出
func autoLogout(ws *websocket.Conn, bot *openwechat.Bot) {
	ws.SetCloseHandler(func(code int, text string) error {
		utils.Logger.Infoln("客户端断开ws连接")
		//bot登出
		err := bot.Logout()
		if err != nil && err != io.EOF {
			utils.Logger.Errorln("bot登出失败", err)
			return err
		}
		return nil
	})
}

func ReceiveMessageHandle(ws *websocket.Conn, msg *openwechat.Message) {
	go dispatcher.ReceiveMessageDispatcher(ws, msg)
}
