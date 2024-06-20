package main

import (
	"aiwechat/application"
	"aiwechat/application/utils"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	utils.Logger = logrus.New()
	utils.Logger.SetFormatter(&logrus.JSONFormatter{})
	utils.Logger.SetOutput(os.Stdout)
	utils.Logger.Infoln("aiwechat start")
	application.RUN()
}
