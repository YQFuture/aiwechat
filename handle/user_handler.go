package handle

import (
	"aiwechat/model"
	"encoding/json"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func LoginCallBack(ws *websocket.Conn, bot *openwechat.Bot, body openwechat.CheckLoginResponse) {
	fmt.Println(string(body))
	code, err := body.Code()
	if err != nil {
		log.Println("获取登录返回码值失败", err)
		return
	}
	if code == openwechat.LoginCodeSuccess {
		self, err := bot.GetCurrentUser()
		if err != nil {
			fmt.Println(err)
		}
		userModel := &model.UserModel{
			UserName: self.UserName,
			NickName: self.NickName,
			AvatarID: self.AvatarID(),
		}
		marshal, err := json.Marshal(userModel)
		if err != nil {
			log.Println("序列化登录用户信息失败", err)
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
