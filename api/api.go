package api

import (
	"aiwechat/application/utils"
	"aiwechat/handle"
	"aiwechat/model"
	"encoding/json"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func Conn(c *gin.Context) {
	//升级到WebSocket
	r := c.Request
	w := c.Writer
	ws, err := utils.UPGRADER.Upgrade(w, r, nil)
	if err != nil {
		log.Println("升级连接为WebSocket连接失败", err)
		c.JSON(400, gin.H{
			"message": "升级连接为WebSocket连接失败",
		})
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			log.Fatal("WebSocket连接关闭失败", err)
		}
	}(ws)

	// 创建一个bot,一个bot对应一个登录的微信号。
	bot := openwechat.DefaultBot(openwechat.Desktop)

	//初始化bot配置
	botInit(ws, bot)

	// 登录
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return
	}

	// 处理消息
	go handle.ConnHandler(ws, bot)

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	err = bot.Block()
	if err != nil {
		return
	}
}

func botInit(ws *websocket.Conn, bot *openwechat.Bot) {
	// 注册登录二维码回调
	bot.UUIDCallback = func(uuid string) {
		url := openwechat.GetQrcodeUrl(uuid)
		// 将登录二维码链接发送给前端
		responseModel := &model.ResponseModel{
			Operation: model.ReturnQrCodeUrl,
			Content:   url,
			Timestamp: time.Now(),
		}
		responseModelBytes, err := json.Marshal(responseModel)
		if err != nil {
			log.Fatal("登录二维码序列化失败", err)
		}
		err = ws.WriteMessage(websocket.TextMessage, responseModelBytes)
		if err != nil {
			log.Fatal("发送二维码链接失败", err)
		}
	}
	//注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		go handle.ReceiveMessageAdapter(ws, bot, msg)
	}
	//登录回调函数
	bot.LoginCallBack = func(body openwechat.CheckLoginResponse) {
		handle.LoginCallBack(ws, bot, body)
	}
}
