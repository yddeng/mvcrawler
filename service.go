package mvcrawler

import (
	"github.com/tagDong/mvcrawler/conf"
	"github.com/tagDong/mvcrawler/dhttp"
	"time"
)

type Service struct {
	modules map[ModuleType]Module

	analysis   *Analysis
	downloader *Downloader
	hServer    *dhttp.HttpServer
}

func NewService() *Service {
	InitLogger()
	s := new(Service)

	s.initAnalysis()
	s.initDownloader()
	s.initModules()

	//
	s.initHttpServer()

	go s.tick()
	return s
}

func (s *Service) initModules() {
	s.modules = map[ModuleType]Module{}
	for mt, fn := range moduleFunc {
		s.modules[mt] = fn(s.analysis, s.downloader, logger)
	}
}

//初始化分析器
func (s *Service) initAnalysis() {
	config := conf.GetConfig().Analysis
	s.analysis = NewAnalysis(config.QueueSize, config.GoroutineCount)

	logger.Infoln("init analysis ok")
}

//初始化下载器
func (s *Service) initDownloader() {
	config := conf.GetConfig().DownLoad
	s.downloader = NewDownLoader(
		config.OutPath, config.QueueSize, config.GoroutineCount, logger)

	logger.Infoln("init analysis ok")
}

//http服务
func (s *Service) initHttpServer() {
	config := conf.GetConfig()
	s.hServer = dhttp.NewHttpServer(config.HttpAddr)

	//注册路由
	s.hServer.Register("/search", s.search)
	s.hServer.Register("/update", s.update)

	go func() {
		err := s.hServer.Listen()
		if err != nil {
			panic(err)
		}
	}()

	logger.Infof("httpServer start on %s", config.HttpAddr)
}

//定时抓取
func (s *Service) tick() {
	_updata = &UpdateResp{
		resp: map[ModuleType][]*Message{},
	}

	config := conf.GetConfig()
	tick := time.NewTicker(time.Duration(config.TickDur) * time.Second)
	for {
		for k, m := range s.modules {
			_updata.resp[k] = m.Update()
		}
		<-tick.C
		//logger.Infof("-- tick %s\n", now.String())
	}
}
