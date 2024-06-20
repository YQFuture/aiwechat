package model

import (
	"encoding/json"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

// 操作
const (
	ReturnQrCodeUrl = iota + 1
	ReturnUserInfo
	ReturnFriendList
	ReturnGroupList
	ReturnMessage
	ReturnHeadImg
	ReturnGroupHeadImg
)

// 消息目标类型
const (
	FriendType = iota + 1
	GroupType
)

type ResponseModel struct {
	Operation         int                 `json:"operation" validate:"required"` //操作 1:返回登录二维码url 2:返回用户信息 3:返回好友列表 4:返回群列表 5:返回消息 6:返回头像 7:返回群头像
	MessageType       int                 `json:"messageType"`                   //消息类型 1:文本消息 2:图片消息 3:视频消息 11:好友请求 21未知消息
	MessageTargetType int                 `json:"messageTargetType"`             //消息来源类型 1:好友 2:群
	MessageTarget     string              `json:"messageTarget"`                 //消息来源
	Content           string              `json:"content"`                       //文本消息内容
	FileName          string              `json:"fileName"`                      //文件消息文件名
	FileData          []byte              `json:"fileData"`                      //文件消息内容
	Timestamp         time.Time           `json:"timestamp"`                     //时间戳
	Msg               *openwechat.Message `json:"msg"`                           //系统消息体 同意好友申请时需原样返回
	MsgUserModel      *UserModel          `json:"msgUserModel"`                  //消息来源的用户对象 在收到群消息时可以使用该对象获取发送者昵称和头像
}

func ReturnModel(ws *websocket.Conn, responseModel *ResponseModel) {
	responseModelBytes, err := json.Marshal(responseModel)
	if err != nil {
		return
	}
	err = ws.WriteMessage(websocket.TextMessage, responseModelBytes)
	if err != nil {
		log.Println("返回消息失败:", err)
	}
}
