package mvcrawler

/*
 mvcrawler logger 日志
*/
import (
	"github.com/tagDong/mvcrawler/conf"
	"github.com/tagDong/mvcrawler/util"
)

var logger *util.Logger

func InitLogger() {
	logConf := conf.GetConfig().Common.Log
	logger = util.NewLogger(logConf.LogPath, logConf.LogName)
	logger.Infoln("crawler logger init")
}
