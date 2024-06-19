package model

import (
	"github.com/eatmoreapple/openwechat"
	"time"
)

// 操作
const (
	SendMessage = iota + 1
	SendGroupMessage
	GetFriendList
	GetGroupList
	GetHeadImg
	GetGroupHeadImg
	AcceptFriendRequest
)

// 消息类型
const (
	TextMessage = iota + 1
	ImageMessage
	VideoMessage
	FriendAddMessage = 11
)

type RequestModel struct {
	Operation     int                 `json:"operation" validate:"required"` //操作 1:发送消息 2:发送群消息 3:获取好友列表 4:获取群列表 5:获取头像 6:获取群头像 7:同意好友请求
	MessageType   int                 `json:"messageType"`                   //消息类型 1:文本消息 2:图片消息 3:视频消息
	MessageTarget string              `json:"messageTarget"`                 //消息目标
	Content       string              `json:"content"`                       //文本消息内容
	FileName      string              `json:"fileName"`                      //文件消息文件名
	FileData      []byte              `json:"fileData"`                      //文件消息内容
	Timestamp     time.Time           `json:"timestamp"`                     //时间戳
	Msg           *openwechat.Message `json:"msg"`                           //系统消息体 同意好友申请时需原样返回
}
