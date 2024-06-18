package handle

import (
	"aiwechat/model"
	"github.com/eatmoreapple/openwechat"
	"log"
	"os"
)

func SendMessage(bot *openwechat.Bot, messageModel *model.RequestModel) {
	self, err := bot.GetCurrentUser()
	if err != nil {
		return
	}

	//如果消息类型是图片/文件, 则预先创建
	var file *os.File
	if messageModel.MessageType == model.ImageMessage || messageModel.MessageType == model.FileMessage {
		file, err := os.Create(messageModel.FileName)
		if err != nil {
			log.Println("创建文件出错", err)
			return
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Println("关闭文件出错", err)
				return
			}
		}(file)

		_, err = file.Write(messageModel.FileData)
		if err != nil {
			log.Println("写入文件出错", err)
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
					log.Println("发送文本消息出错", err)
					return
				}
			case model.ImageMessage:
				_, err = friend.SendImage(file)
				if err != nil {
					log.Println("发送图片消息出错", err)
					return
				}
			case model.FileMessage:
				_, err = friend.SendFile(file)
				if err != nil {
					log.Println("发送文件消息出错", err)
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
	if messageModel.MessageType == model.ImageMessage || messageModel.MessageType == model.FileMessage {
		file, err := os.Create(messageModel.FileName)
		if err != nil {
			log.Println("创建文件出错", err)
			return
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Println("关闭文件出错", err)
				return
			}
		}(file)

		_, err = file.Write(messageModel.FileData)
		if err != nil {
			log.Println("写入文件出错", err)
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
					log.Println("发送文本消息出错", err)
					return
				}
			case model.ImageMessage:
				_, err = group.SendImage(file)
				if err != nil {
					log.Println("发送图片消息出错", err)
					return
				}
			case model.FileMessage:
				_, err = group.SendFile(file)
				if err != nil {
					log.Println("发送文件消息出错", err)
					return
				}
			}
		}
	}
}
