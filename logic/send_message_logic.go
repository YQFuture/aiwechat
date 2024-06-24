package logic

import (
	"aiwechat/application/utils"
	"aiwechat/model"
	"github.com/eatmoreapple/openwechat"
	"os"
)

func SendMessage(bot *openwechat.Bot, messageModel *model.RequestModel) {
	self, err := bot.GetCurrentUser()
	if err != nil {
		return
	}

	//如果消息类型是图片/文件, 则预先创建
	var file *os.File
	if messageModel.MessageType == model.ImageMessage || messageModel.MessageType == model.VideoMessage {
		CreateFile(messageModel)
		file, err = OpenFile(messageModel)
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				utils.Logger.Errorln("关闭文件出错", err)
			}
		}(file)
		if err != nil {
			utils.Logger.Errorln("打开文件出错", err)
			return
		}
	}

	Friends, err := self.Friends()

	for _, friend := range Friends {
		if friend.AvatarID() == messageModel.MessageTarget {
			switch messageModel.MessageType {
			case model.TextMessage:
				_, err := friend.SendText(messageModel.Content)
				if err != nil {
					utils.Logger.Errorln("发送文本消息出错", err)
					return
				}
			case model.ImageMessage:
				_, err = friend.SendImage(file)
				if err != nil {
					utils.Logger.Errorln("发送图片消息出错", err)
					return
				}
			case model.VideoMessage:
				_, err = friend.SendVideo(file)
				if err != nil {
					utils.Logger.Errorln("发送视频消息出错", err)
					return
				}
			}
		}
	}
}

func SendGroupMessage(bot *openwechat.Bot, messageModel *model.RequestModel) {
	self, err := bot.GetCurrentUser()
	if err != nil {
		return
	}

	//如果消息类型是图片/文件, 则预先创建
	var file *os.File
	if messageModel.MessageType == model.ImageMessage || messageModel.MessageType == model.VideoMessage {
		CreateFile(messageModel)
		file, err = OpenFile(messageModel)
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				utils.Logger.Errorln("关闭文件出错", err)
			}
		}(file)
		if err != nil {
			utils.Logger.Errorln("打开文件出错", err)
			return
		}
	}

	Groups, err := self.Groups()

	for _, group := range Groups {
		if group.AvatarID() == messageModel.MessageTarget {
			switch messageModel.MessageType {
			case model.TextMessage:
				_, err := group.SendText(messageModel.Content)
				if err != nil {
					utils.Logger.Errorln("发送文本消息出错", err)
					return
				}
			case model.ImageMessage:
				_, err = group.SendImage(file)
				if err != nil {
					utils.Logger.Errorln("发送图片消息出错", err)
					return
				}
			case model.VideoMessage:
				_, err = group.SendVideo(file)
				if err != nil {
					utils.Logger.Errorln("发送视频消息出错", err)
					return
				}
			}
		}
	}
}

func CreateFile(messageModel *model.RequestModel) {
	file, err := os.Create(messageModel.FileName)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			utils.Logger.Errorln("关闭文件出错", err)
		}
	}(file)
	if err != nil {
		utils.Logger.Errorln("创建文件出错", err)
		return
	}
	_, err = file.Write(messageModel.FileData)
	if err != nil {
		utils.Logger.Errorln("写入文件出错", err)
		return
	}
}

func OpenFile(messageModel *model.RequestModel) (*os.File, error) {
	file, err := os.Open(messageModel.FileName)
	if err != nil {
		utils.Logger.Errorln("打开文件出错", err)
		return nil, err
	}
	return file, nil
}
