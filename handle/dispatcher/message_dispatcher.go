package dispatcher

import (
	"aiwechat/application/utils"
	"aiwechat/logic"
	"aiwechat/model"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
)

// MessageDispatcher 分发处理前端发送的请求
func MessageDispatcher(ws *websocket.Conn, bot *openwechat.Bot, messageModel *model.RequestModel) {
	defer func() {
		if err := recover(); err != nil {
			utils.Logger.Errorln("请求分发处理失败", err)
		}
	}()
	//分发处理不同类型的消息
	switch messageModel.Operation {
	case model.SendMessage:
		logic.SendMessage(bot, messageModel)
	case model.SendGroupMessage:
		logic.SendGroupMessage(bot, messageModel)
	case model.GetFriendList:
		logic.GetFriendList(ws, bot)
		logic.GetFriendHeadImgList(ws, bot)
	case model.GetGroupList:
		logic.GetGroupList(ws, bot)
		logic.GetGroupHeadImgList(ws, bot)
	case model.GetHeadImg:
		logic.GetHeadImg(ws, bot, messageModel)
	case model.GetGroupHeadImg:
		logic.GetGroupHeadImg(ws, bot, messageModel)
	case model.AcceptFriendRequest:
		logic.AcceptFriendRequest(ws, bot, messageModel)
	default:
		utils.Logger.Errorln("收到未定义的请求类型", messageModel)
	}
}

// ReceiveMessageDispatcher 分发处理接收到的消息
func ReceiveMessageDispatcher(ws *websocket.Conn, msg *openwechat.Message) {
	defer func() {
		if err := recover(); err != nil {
			utils.Logger.Errorln("消息接收分发处理失败", err)
		}
	}()
	if msg.IsSendByFriend() {
		logic.ReceiveFriendMessage(ws, msg)
	} else if msg.IsSendByGroup() {
		logic.ReceiveGroupMessage(ws, msg)
	} else if msg.IsFriendAdd() {
		utils.Logger.Infoln("收到未定义的消息类型", msg)
		logic.ReceiveFriendAdd(ws, msg)
	}
}
