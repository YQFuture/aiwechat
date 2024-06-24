package init_application

import (
	"aiwechat/application/utils"
	"github.com/sirupsen/logrus"
	"os"
)

func LogrusInit() {
	utils.Logger = logrus.New()
	utils.Logger.SetReportCaller(true)
	utils.Logger.SetFormatter(&logrus.JSONFormatter{})
	utils.Logger.SetOutput(os.Stdout)
}
