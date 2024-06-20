package handle

import (
	"aiwechat/application/utils"
	"aiwechat/model"
	"bytes"
	"encoding/json"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

// MsgMap 保存好友请求等消息
var MsgMap sync.Map

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
	MsgMap.Store(msg.MsgId, msg)
	model.ReturnModel(ws, responseModel)
}

func ReceiveGroupMessage(ws *websocket.Conn, msg *openwechat.Message) {
	switch msg.MsgType {
	case openwechat.MsgTypeText:
		ReceiveTextMessage(ws, msg, model.GroupType)
	case openwechat.MsgTypeImage:
		ReceiveImageMessage(ws, msg, model.GroupType)
	case openwechat.MsgTypeVideo:
		ReceiveVideoMessage(ws, msg, model.GroupType)
	default:
		ReceiveUnknownMessage(ws, msg, model.GroupType)
	}
}

func ReceiveFriendMessage(ws *websocket.Conn, msg *openwechat.Message) {
	switch msg.MsgType {
	case openwechat.MsgTypeText:
		ReceiveTextMessage(ws, msg, model.FriendType)
	case openwechat.MsgTypeImage:
		ReceiveImageMessage(ws, msg, model.FriendType)
	case openwechat.MsgTypeVideo:
		ReceiveVideoMessage(ws, msg, model.FriendType)
	default:
		ReceiveUnknownMessage(ws, msg, model.FriendType)
	}
}

func ReceiveTextMessage(ws *websocket.Conn, msg *openwechat.Message, messageTargetType int) {
	sender, err := msg.Sender()
	if err != nil {
		log.Println("获取文本消息发送者失败:", err)
		return
	}
	var userModel *model.UserModel
	var groupModel *model.GroupModel
	if messageTargetType == model.GroupType {
		userModel = getGroupMsgUserModel(msg)
		groupModel = getGroupMsgGroupModel(msg)
	}
	text := msg.Content
	responseModel := &model.ResponseModel{
		Operation:         model.ReturnMessage,
		MessageType:       model.TextMessage,
		MessageTargetType: messageTargetType,
		MessageTarget:     sender.AvatarID(),
		Content:           text,
		Timestamp:         time.Now(),
		MsgUserModel:      userModel,
		MsgGroupModel:     groupModel,
	}
	model.ReturnModel(ws, responseModel)
}

func ReceiveImageMessage(ws *websocket.Conn, msg *openwechat.Message, messageTargetType int) {
	sender, err := msg.Sender()
	if err != nil {
		log.Println("获取图片消息发送者失败:", err)
		return
	}
	var userModel *model.UserModel
	var groupModel *model.GroupModel
	if messageTargetType == model.GroupType {
		userModel = getGroupMsgUserModel(msg)
		groupModel = getGroupMsgGroupModel(msg)
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
		MsgUserModel:      userModel,
		MsgGroupModel:     groupModel,
	}
	model.ReturnModel(ws, responseModel)
}

func ReceiveVideoMessage(ws *websocket.Conn, msg *openwechat.Message, messageTargetType int) {
	sender, err := msg.Sender()
	if err != nil {
		log.Println("获取文件消息发送者失败:", err)
		return
	}
	var userModel *model.UserModel
	var groupModel *model.GroupModel
	if messageTargetType == model.GroupType {
		userModel = getGroupMsgUserModel(msg)
		groupModel = getGroupMsgGroupModel(msg)
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
		MessageType:       model.VideoMessage,
		MessageTargetType: messageTargetType,
		MessageTarget:     sender.AvatarID(),
		FileName:          msg.FileName,
		FileData:          buf.Bytes(),
		Timestamp:         time.Now(),
		MsgUserModel:      userModel,
		MsgGroupModel:     groupModel,
	}
	model.ReturnModel(ws, responseModel)
}

func ReceiveUnknownMessage(ws *websocket.Conn, msg *openwechat.Message, messageTargetType int) {
	sender, err := msg.Sender()
	if err != nil {
		log.Println("获取未知消息发送者失败:", err)
		return
	}
	responseModel := &model.ResponseModel{
		Operation:         model.ReturnMessage,
		MessageType:       model.UnknownMessage,
		MessageTargetType: messageTargetType,
		MessageTarget:     sender.AvatarID(),
		Timestamp:         time.Now(),
	}
	model.ReturnModel(ws, responseModel)
}

func getGroupMsgUserModel(msg *openwechat.Message) (userModel *model.UserModel) {
	groupMsgUser, err := msg.SenderInGroup()
	if err != nil {
		log.Println("获取群消息发送者失败:", err)
		return
	}
	var buf bytes.Buffer
	resp, err := groupMsgUser.GetAvatarResponse()
	if err != nil {
		return
	}
	utils.RespToBuf(resp, &buf)

	userModel = &model.UserModel{
		UserName: groupMsgUser.UserName,
		NickName: groupMsgUser.NickName,
		AvatarID: groupMsgUser.AvatarID(),
		FileData: buf.Bytes(),
	}

	return userModel
}

func getGroupMsgGroupModel(msg *openwechat.Message) (groupModel *model.GroupModel) {
	sender, err := msg.Sender()
	if err != nil {
		log.Println("获取群消息发送群失败:", err)
		return
	}
	var buf bytes.Buffer
	resp, err := sender.GetAvatarResponse()
	if err != nil {
		return
	}
	utils.RespToBuf(resp, &buf)

	groupModel = &model.GroupModel{
		GroupName: sender.NickName,
		AvatarID:  sender.AvatarID(),
		FileData:  buf.Bytes(),
	}

	return groupModel
}
