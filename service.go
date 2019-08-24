package mvcrawler

import (
	"fmt"
	"github.com/tagDong/mvcrawler/conf"
	"github.com/tagDong/mvcrawler/dhttp"
	"time"
)

var (
	tickDur time.Duration = 60 * time.Second
)

type Service struct {
	modules map[ModuleType]Module

	analysis *Analysis
	hServer  *dhttp.HttpServer
}

func NewService() *Service {
	InitLogger()
	s := new(Service)
	s.initModules()

	s.InitAnalysis()
	s.InitHttpServer()

	go s.tick()
	return s
}

func (s *Service) initModules() {
	for mt, fn := range moduleFunc {
		s.modules[mt] = fn()
	}
}

func (s *Service) InitHttpServer() {
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

func (s *Service) InitAnalysis() {
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
