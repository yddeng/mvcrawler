package mvcrawler

/*
 mvcrawler logger 日志
*/
import (
	"github.com/tagDong/dutil/log"
	"github.com/tagDong/mvcrawler/conf"
)

var logger *log.Logger

func InitLogger() {
	logConf := conf.GetConfig().Log
	logger = log.NewLogger(logConf.LogPath, logConf.LogName, 1024*1024)
	//log.CloseStdOut()
	logger.AsyncOut()
	logger.Infoln("crawler logger init")
}
