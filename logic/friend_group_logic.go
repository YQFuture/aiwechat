package logic

import (
	"aiwechat/application/utils"
	"aiwechat/model"
	"bytes"
	"encoding/json"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

func AcceptFriendRequest(ws *websocket.Conn, bot *openwechat.Bot, messageModel *model.RequestModel) {
	msgId := messageModel.Msg.MsgId
	value, ok := MsgMap.Load(msgId)
	if ok {
		msg := value.(*openwechat.Message)
		_, err := msg.Agree()
		if err != nil {
			utils.Logger.Errorln("同意好友请求失败")
			return
		}
		MsgMap.Delete(msgId)
		//同意好友请求后返回新的好友列表
		GetFriendList(ws, bot)
		GetFriendHeadImgList(ws, bot)
	}
}

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
			resp, err := group.GetAvatarResponse()
			if err != nil {
				return
			}
			utils.RespToBuf(resp, &buf)
			//构建标准返回体
			responseModel := &model.ResponseModel{
				Operation:     model.ReturnGroupHeadImg,
				MessageTarget: group.AvatarID(),
				FileData:      buf.Bytes(),
				Timestamp:     time.Now(),
			}
			model.ReturnModel(ws, responseModel)
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
			resp, err := friend.GetAvatarResponse()
			if err != nil {
				return
			}
			utils.RespToBuf(resp, &buf)

			//构建标准返回体
			responseModel := &model.ResponseModel{
				Operation:     model.ReturnHeadImg,
				MessageTarget: friend.AvatarID(),
				FileData:      buf.Bytes(),
				Timestamp:     time.Now(),
			}
			model.ReturnModel(ws, responseModel)
		}
	}
}

func GetGroupList(ws *websocket.Conn, bot *openwechat.Bot) {
	self, err := bot.GetCurrentUser()
	if err != nil {
		return
	}
	Groups, err := self.Groups()
	if err != nil {
		return
	}

	//创建群列表
	groupList := model.GroupModelList{}

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
		Timestamp: time.Now(),
	}
	model.ReturnModel(ws, responseModel)
}

func GetGroupHeadImgList(ws *websocket.Conn, bot *openwechat.Bot) {
	self, err := bot.GetCurrentUser()
	if err != nil {
		return
	}
	Groups, err := self.Groups()
	if err != nil {
		return
	}

	//创建群列表
	groupList := model.GroupModelList{}

	var headImgMap sync.Map
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(Groups))

	for _, group := range Groups {
		go func() {
			defer waitGroup.Done()
			//获取群头像
			var buf bytes.Buffer
			resp, err := group.GetAvatarResponse()
			if err != nil {
				return
			}
			utils.RespToBuf(resp, &buf)
			groupModel := &model.GroupModel{
				GroupName: group.NickName,
				AvatarID:  group.AvatarID(),
				FileData:  buf.Bytes(),
			}
			headImgMap.Store(group.AvatarID(), groupModel)
		}()
	}

	waitGroup.Wait()

	headImgMap.Range(func(key, value interface{}) bool {
		groupModel := value.(*model.GroupModel)
		groupList = append(groupList, groupModel)
		return true
	})

	groupListBytes, err := json.Marshal(groupList)
	if err != nil {
		return
	}

	//构建标准返回体
	responseModel := &model.ResponseModel{
		Operation: model.ReturnGroupHeadImg,
		FileData:  groupListBytes,
		Timestamp: time.Now(),
	}
	model.ReturnModel(ws, responseModel)
}

func GetFriendList(ws *websocket.Conn, bot *openwechat.Bot) {
	self, err := bot.GetCurrentUser()
	if err != nil {
		return
	}
	Friends, err := self.Friends(true)
	if err != nil {
		return
	}

	//创建好友列表
	friendList := model.UserModelList{}
	for _, friend := range Friends {
		remarkName := friend.RemarkName
		if len(friend.RemarkName) == 0 {
			remarkName = friend.NickName
		}
		userModel := &model.UserModel{
			UserName:   friend.UserName,
			NickName:   friend.NickName,
			RemarkName: remarkName,
			AvatarID:   friend.AvatarID(),
		}
		//将解析的好友列表保存到切片中
		friendList = append(friendList, userModel)
	}

	//按照首字母分组
	groupByInitial := model.UserGroupByInitial(friendList)

	friendListBytes, err := json.Marshal(groupByInitial)
	if err != nil {
		return
	}

	//构建标准返回体
	responseModel := &model.ResponseModel{
		Operation: model.ReturnFriendList,
		FileData:  friendListBytes,
		Timestamp: time.Now(),
	}
	model.ReturnModel(ws, responseModel)
}

func GetFriendHeadImgList(ws *websocket.Conn, bot *openwechat.Bot) {
	self, err := bot.GetCurrentUser()
	if err != nil {
		return
	}
	Friends, err := self.Friends(true)
	if err != nil {
		return
	}

	//创建好友列表
	friendList := model.UserModelList{}

	var headImgMap sync.Map
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(Friends))

	for _, friend := range Friends {
		go func() {
			defer waitGroup.Done()
			//获取用户头像
			var buf bytes.Buffer
			resp, err := friend.GetAvatarResponse()
			if err != nil {
				return
			}
			utils.RespToBuf(resp, &buf)
			remarkName := friend.RemarkName
			if len(friend.RemarkName) == 0 {
				remarkName = friend.NickName
			}
			userModel := &model.UserModel{
				UserName:   friend.UserName,
				NickName:   friend.NickName,
				RemarkName: remarkName,
				AvatarID:   friend.AvatarID(),
				FileData:   buf.Bytes(),
			}
			headImgMap.Store(friend.AvatarID(), userModel)
		}()
	}

	waitGroup.Wait()

	headImgMap.Range(func(key, value interface{}) bool {
		userModel := value.(*model.UserModel)
		friendList = append(friendList, userModel)
		return true
	})

	//按照首字母分组
	groupByInitial := model.UserGroupByInitial(friendList)

	friendListBytes, err := json.Marshal(groupByInitial)
	if err != nil {
		return
	}

	//构建标准返回体
	responseModel := &model.ResponseModel{
		Operation: model.ReturnHeadImg,
		FileData:  friendListBytes,
		Timestamp: time.Now(),
	}
	model.ReturnModel(ws, responseModel)
}
