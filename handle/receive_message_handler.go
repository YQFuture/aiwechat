package handle

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
)

func ReceiveMessageAdapter(ws *websocket.Conn, bot *openwechat.Bot, msg *openwechat.Message) {
	//获取消息发送
	sender, err := msg.Sender()
	if err != nil {
		return
	}

	sender.AvatarID()

	switch msg.MsgType {
	case openwechat.MsgTypeText:
		//文本消息
		text := msg.Content
		fmt.Println(text)
	}

}
