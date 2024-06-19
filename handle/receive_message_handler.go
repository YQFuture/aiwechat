package handle

import (
	"aiwechat/application/utils"
	"aiwechat/model"
	"bytes"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
	"log"
)

func ReceiveGroupMessage(ws *websocket.Conn, bot *openwechat.Bot, msg *openwechat.Message) {
}

func ReceiveFriendMessage(ws *websocket.Conn, bot *openwechat.Bot, msg *openwechat.Message) {
	switch msg.MsgType {
	case openwechat.MsgTypeText:
		ReceiveTextMessage(ws, msg, model.FriendType)
	case openwechat.MsgTypeImage:
		ReceiveImageMessage(ws, msg, model.FriendType)

	}
}

func ReceiveTextMessage(ws *websocket.Conn, msg *openwechat.Message, messageTargetType int) {
	sender, err := msg.Sender()
	if err != nil {
		log.Println("获取文本消息发送者失败:", err)
		return
	}
	text := msg.Content
	responseModel := &model.ResponseModel{
		Operation:         model.ReturnMessage,
		MessageType:       model.TextMessage,
		MessageTargetType: messageTargetType,
		MessageTarget:     sender.AvatarID(),
		Content:           text,
	}
	model.ReturnModel(ws, responseModel)
}

func ReceiveImageMessage(ws *websocket.Conn, msg *openwechat.Message, messageTargetType int) {
	sender, err := msg.Sender()
	if err != nil {
		log.Println("获取图片消息发送者失败:", err)
		return
	}
	picture, err := msg.GetPicture()
	if err != nil {
		log.Println("返回图片消息失败:", err)
		return
	}
	//获取图片字节流数组
	var buf bytes.Buffer
	utils.RespToBuf(picture, &buf)
	//构建消息并返回
	responseModel := &model.ResponseModel{
		Operation:         model.ReturnMessage,
		MessageType:       model.ImageMessage,
		MessageTargetType: messageTargetType,
		MessageTarget:     sender.AvatarID(),
		FileName:          msg.FileName,
		FileData:          buf.Bytes(),
	}
	model.ReturnModel(ws, responseModel)
}
