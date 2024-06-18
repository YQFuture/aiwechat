package handle

import (
	"aiwechat/application/utils"
	"aiwechat/model"
	"bytes"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
	"log"
)

func ReceiveFriendMessage(ws *websocket.Conn, bot *openwechat.Bot, msg *openwechat.Message) {
	//获取消息发送者
	sender, err := msg.Sender()
	if err != nil {
		log.Println("获取文本消息发送者失败:", err)
		return
	}
	switch msg.MsgType {
	case openwechat.MsgTypeText:
		//文本消息
		text := msg.Content
		responseModel := &model.ResponseModel{
			Operation:         model.ReturnMessage,
			MessageType:       model.TextMessage,
			MessageTargetType: model.FriendType,
			MessageTarget:     sender.AvatarID(),
			Content:           text,
		}
		model.ReturnModel(ws, responseModel)
	case openwechat.MsgTypeImage:
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
			MessageTargetType: model.FriendType,
			MessageTarget:     sender.AvatarID(),
			FileName:          msg.FileName,
			FileData:          buf.Bytes(),
		}
		model.ReturnModel(ws, responseModel)
		//TODO

	}
}
