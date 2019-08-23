package mvcrawler

import (
	"github.com/tagDong/mvcrawler/conf"
	"github.com/tagDong/mvcrawler/dhttp"
	"time"
)

var (
	tickDur time.Duration
)

type Service struct {
	hServer  *dhttp.HttpServer
	analysis *Analysis

	idxReq []*Request
}

func NewService() *Service {

	tickDur = 60 * time.Second

	InitLogger()
	s := new(Service)
	s.initReq()

	s.InitAnalysis()
	s.InitHttpServer()

	go s.tick()
	return s
}

func (s *Service) initReq() {
	config := conf.GetConfig().Update
	reqs := []*Request{}
	for _, v := range config.Urls {
		reqs = append(reqs, &Request{
			Url:      v.Index,
			Selector: v.Selectors[0],
			Depth:    1,
		})
	}
	s.idxReq = reqs
}

func (s *Service) InitHttpServer() {
	config := conf.GetConfig()
	s.hServer = dhttp.NewHttpServer(config.Common.HttpAddr)
	s.hServer.Register(config.Update.Route, update)

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
	s.analysis = NewAnalysis(respCh)

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
		s.analysis.Push(s.idxReq)
	}
}
