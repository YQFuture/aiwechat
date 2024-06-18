package handle

import (
	"aiwechat/model"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
)

func MessageDispatcher(ws *websocket.Conn, bot *openwechat.Bot, messageModel *model.RequestModel) {
	//分发处理不同类型的消息
	switch messageModel.Operation {
	case model.SendMessage:
		SendMessage(bot, messageModel)
	case model.SendGroupMessage:
		SendGroupMessage(bot, messageModel)
	case model.GetFriendList:
		GetFriendList(ws, bot, messageModel)
	case model.GetGroupList:
		GetGroupList(ws, bot, messageModel)
	case model.GetHeadImg:
		GetHeadImg(ws, bot, messageModel)
	case model.GetGroupHeadImg:
		GetGroupHeadImg(ws, bot, messageModel)
	default:
		//TODO
	}
}

func ReceiveMessageAdapter(ws *websocket.Conn, bot *openwechat.Bot, msg *openwechat.Message) {
	if msg.IsSendByFriend() {
		ReceiveFriendMessage(ws, bot, msg)
	}
}
