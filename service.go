package mvcrawler

import (
	"fmt"
	"github.com/tagDong/mvcrawler/conf"
	"github.com/tagDong/mvcrawler/dhttp"
	"time"
)

var (
	tickDur = 60 * time.Second
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
	for mt, fn := range moduleFunc {
		s.modules[mt] = fn(s.analysis, s.downloader, logger)
	}
}

//初始化分析器
func (s *Service) initAnalysis() {
	respCh := make(chan [][]string, 100)

	go func() {
		for list := range respCh {
			for _, url := range list {
				logger.Infoln(url[0], url[1])
			}
		}
	}()

	logger.Infoln("init analysis ok")
}

//初始化下载器
func (s *Service) initDownloader() {

}

//http服务
func (s *Service) initHttpServer() {
	config := conf.GetConfig()
	s.hServer = dhttp.NewHttpServer(config.Common.HttpAddr)

	go func() {
		err := s.hServer.Listen()
		if err != nil {
			panic(err)
		}
	}()

	logger.Infoln("init httpServer ok")
}

//定时抓取
func (s *Service) tick() {
	tick := time.NewTicker(tickDur)
	for {
		now := <-tick.C
		logger.Infof("-------- tick %s------", now.String())
		for mt, module := range s.modules {
			module.Update()
			fmt.Printf("%s update", MT2String(mt))
		}
	}
}
