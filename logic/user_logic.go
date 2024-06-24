package logic

import (
	"aiwechat/application/utils"
	"aiwechat/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
	"time"
)

func LoginCallBack(ws *websocket.Conn, bot *openwechat.Bot, body openwechat.CheckLoginResponse) {
	fmt.Println(string(body))
	code, err := body.Code()
	if err != nil {
		utils.Logger.Errorln("获取登录返回码值失败", err)
		return
	}
	if code == openwechat.LoginCodeSuccess {
		self, err := bot.GetCurrentUser()
		if err != nil {
			fmt.Println(err)
		}

		//获取用户头像
		var buf bytes.Buffer
		resp, err := self.GetAvatarResponse()
		if err != nil {
			return
		}
		utils.RespToBuf(resp, &buf)

		userModel := &model.UserModel{
			UserName: self.UserName,
			NickName: self.NickName,
			AvatarID: self.AvatarID(),
			FileData: buf.Bytes(),
		}
		marshal, err := json.Marshal(userModel)
		if err != nil {
			utils.Logger.Errorln("序列化登录用户信息失败", err)
			return
		}

		responseModel := &model.ResponseModel{
			Operation: model.ReturnUserInfo,
			FileData:  marshal,
			Timestamp: time.Now(),
		}
		model.ReturnModel(ws, responseModel)
	}
}
