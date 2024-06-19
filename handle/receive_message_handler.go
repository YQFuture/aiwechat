package handle

import (
	"aiwechat/application/utils"
	"aiwechat/model"
	"bytes"
	"encoding/json"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func ReceiveFriendAdd(ws *websocket.Conn, msg *openwechat.Message) {
	sender, err := msg.Sender()
	if err != nil {
		log.Println("获取好友请求消息发送者失败:", err)
		return
	}
	friendAddMessageContent, err := msg.FriendAddMessageContent()
	if err != nil {
		log.Println("获取好友请求消息消息内容失败:", err)
		return
	}
	marshal, err := json.Marshal(friendAddMessageContent)
	//构建消息并返回
	responseModel := &model.ResponseModel{
		Operation:     model.ReturnMessage,
		MessageType:   model.FriendAddMessage,
		MessageTarget: sender.AvatarID(),
		FileName:      msg.FileName,
		FileData:      marshal,
		Timestamp:     time.Now(),
		Msg:           msg,
	}
	model.ReturnModel(ws, responseModel)
}

func ReceiveGroupMessage(ws *websocket.Conn, msg *openwechat.Message) {
	switch msg.MsgType {
	case openwechat.MsgTypeText:
		ReceiveTextMessage(ws, msg, model.GroupType)
	case openwechat.MsgTypeImage:
		ReceiveImageMessage(ws, msg, model.GroupType)
	case openwechat.MsgTypeVideo:
		ReceiveFileMessage(ws, msg, model.GroupType)
	}
}

func ReceiveFriendMessage(ws *websocket.Conn, msg *openwechat.Message) {
	switch msg.MsgType {
	case openwechat.MsgTypeText:
		ReceiveTextMessage(ws, msg, model.FriendType)
	case openwechat.MsgTypeImage:
		ReceiveImageMessage(ws, msg, model.FriendType)
	case openwechat.MsgTypeVideo:
		ReceiveFileMessage(ws, msg, model.FriendType)
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
		Timestamp:         time.Now(),
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
		Timestamp:         time.Now(),
	}
	model.ReturnModel(ws, responseModel)
}

func ReceiveFileMessage(ws *websocket.Conn, msg *openwechat.Message, messageTargetType int) {
	sender, err := msg.Sender()
	if err != nil {
		log.Println("获取文件消息发送者失败:", err)
		return
	}
	file, err := msg.GetVideo()
	if err != nil {
		log.Println("返回文件消息失败:", err)
		return
	}
	//获取文件字节流数组
	var buf bytes.Buffer
	utils.RespToBuf(file, &buf)
	//构建消息并返回
	responseModel := &model.ResponseModel{
		Operation:         model.ReturnMessage,
		MessageType:       model.ImageMessage,
		MessageTargetType: messageTargetType,
		MessageTarget:     sender.AvatarID(),
		FileName:          msg.FileName,
		FileData:          buf.Bytes(),
		Timestamp:         time.Now(),
	}
	model.ReturnModel(ws, responseModel)
}
