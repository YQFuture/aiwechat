package handle

import (
	"aiwechat/model"
	"bytes"
	"encoding/json"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
	"io"
	"log"
)

func GetGroupHeadImg(ws *websocket.Conn, bot *openwechat.Bot, messageModel *model.RequestModel) {
	self, err := bot.GetCurrentUser()
	if err != nil {
		return
	}

	Groups, err := self.Groups()
	if err != nil {
		return
	}

	for _, group := range Groups {
		if group.AvatarID() == messageModel.MessageTarget {

			var buf bytes.Buffer
			func() {
				resp, err := group.GetAvatarResponse()
				if err != nil {
					return
				}
				if resp.ContentLength == 0 {
					return
				}
				defer func(Body io.ReadCloser) {
					err := Body.Close()
					if err != nil {
						log.Println("请求群头像返回体关闭失败:", err)
					}
				}(resp.Body)

				if _, err := io.Copy(&buf, resp.Body); err != nil {
					log.Println("请求群头像返回体复制失败:", err)
					return
				}
			}()

			//构建标准返回体
			responseModel := &model.ResponseModel{
				Operation:     model.ReturnGroupHeadImg,
				MessageTarget: group.AvatarID(),
				FileData:      buf.Bytes(),
			}
			responseModelBytes, err := json.Marshal(responseModel)
			if err != nil {
				return
			}

			err = ws.WriteMessage(websocket.BinaryMessage, responseModelBytes)
			if err != nil {
				log.Println("返回群头像失败:", err)
			}
		}
	}

}

func GetHeadImg(ws *websocket.Conn, bot *openwechat.Bot, messageModel *model.RequestModel) {
	self, err := bot.GetCurrentUser()
	if err != nil {
		return
	}
	Friends, err := self.Friends()
	if err != nil {
		return
	}

	for _, friend := range Friends {
		if friend.AvatarID() == messageModel.MessageTarget {

			var buf bytes.Buffer
			func() {
				resp, err := friend.GetAvatarResponse()
				if err != nil {
					return
				}
				if resp.ContentLength == 0 {
					return
				}
				defer func(Body io.ReadCloser) {
					err := Body.Close()
					if err != nil {
						log.Println("请求头像返回体关闭失败:", err)
					}
				}(resp.Body)

				if _, err := io.Copy(&buf, resp.Body); err != nil {
					log.Println("请求头像返回体复制失败:", err)
					return
				}
			}()

			//构建标准返回体
			responseModel := &model.ResponseModel{
				Operation:     model.ReturnHeadImg,
				MessageTarget: friend.AvatarID(),
				FileData:      buf.Bytes(),
			}
			responseModelBytes, err := json.Marshal(responseModel)
			if err != nil {
				return
			}

			err = ws.WriteMessage(websocket.BinaryMessage, responseModelBytes)
			if err != nil {
				log.Println("返回头像失败:", err)
			}
		}
	}
}

func GetGroupList(ws *websocket.Conn, bot *openwechat.Bot, messageModel *model.RequestModel) {
	self, err := bot.GetCurrentUser()
	if err != nil {
		return
	}
	Groups, err := self.Groups()
	if err != nil {
		return
	}

	//创建群列表
	groupList := make([]*model.GroupModel, 0)

	for _, group := range Groups {
		groupModel := &model.GroupModel{
			GroupName: group.NickName,
			AvatarID:  group.AvatarID(),
		}
		//将解析的群保存到列表中
		groupList = append(groupList, groupModel)
	}

	groupListBytes, err := json.Marshal(groupList)
	if err != nil {
		return
	}

	//构建标准返回体
	responseModel := &model.ResponseModel{
		Operation: model.ReturnGroupList,
		FileData:  groupListBytes,
	}
	responseModelBytes, err := json.Marshal(responseModel)
	if err != nil {
		return
	}

	err = ws.WriteMessage(websocket.BinaryMessage, responseModelBytes)
	if err != nil {
		log.Println("返回好友列表失败:", err)
	}

}

func GetFriendList(ws *websocket.Conn, bot *openwechat.Bot, messageModel *model.RequestModel) {
	self, err := bot.GetCurrentUser()
	if err != nil {
		return
	}
	Friends, err := self.Friends()
	if err != nil {
		return
	}

	//创建好友列表
	friendList := make([]*model.UserModel, 0)

	for _, friend := range Friends {
		userModel := &model.UserModel{
			UserName: friend.UserName,
			NickName: friend.NickName,
			AvatarID: friend.AvatarID(),
		}
		//将解析的好友列表保存到切片中
		friendList = append(friendList, userModel)
	}

	friendListBytes, err := json.Marshal(friendList)
	if err != nil {
		return
	}

	//构建标准返回体
	responseModel := &model.ResponseModel{
		Operation: model.ReturnFriendList,
		FileData:  friendListBytes,
	}
	responseModelBytes, err := json.Marshal(responseModel)
	if err != nil {
		return
	}

	err = ws.WriteMessage(websocket.BinaryMessage, responseModelBytes)
	if err != nil {
		log.Println("返回好友列表失败:", err)
	}
}
