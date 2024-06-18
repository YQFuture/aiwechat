package model

// 操作
const (
	ReturnQrCodeUrl = iota + 1
	ReturnFriendList
	ReturnGroupList
	ReturnMessage
	ReturnHeadImg
	ReturnGroupHeadImg
)

type ResponseModel struct {
	Operation     int    `json:"operation" validate:"required"` //操作 1:返回登录二维码url 2:返回好友列表 3:返回群列表 4:返回消息 5:返回头像
	MessageType   int    `json:"messageType"`                   //消息类型 1:文本消息 2:图片消息 3:文件消息
	MessageTarget string `json:"messageTarget"`                 //消息目标
	Content       string `json:"content"`                       //文本消息内容
	FileName      string `json:"fileName"`                      //文件消息文件名
	FileData      []byte `json:"fileData"`                      //文件消息内容
}
