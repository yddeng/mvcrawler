package mvcrawler

import (
	"github.com/yddeng/mvcrawler/conf"
	"github.com/yddeng/mvcrawler/dhttp"
	"time"
)

type Service struct {
	modules map[ModuleType]Module
	hServer *dhttp.HttpServer
}

func NewService() *Service {
	config := conf.GetConfig()
	InitLogger()
	s := new(Service)

	//db
	NewClient("update", false, &UpdateDB{})
	NewClient("search", true, &SearchDB{})

	//module
	s.modules = map[ModuleType]Module{}
	for mt, fn := range moduleFunc {
		s.modules[mt] = fn(logger)
	}

	//httpServer
	s.hServer = dhttp.NewHttpServer(config.HttpAddr)

	//注册路由
	s.hServer.Register("/search", s.handleSearch)
	s.hServer.Register("/update", s.handleUpdate)

	go func() {
		err := s.hServer.Listen()
		if err != nil {
			panic(err)
		}
	}()

	logger.Infof("httpServer start on %s", config.HttpAddr)

	go s.updateLoop(time.Duration(config.UpdateDur) * time.Second)
	go s.searchLoop(time.Duration(config.SearchDur) * time.Hour)
	return s
}

//update 抓取
func (s *Service) updateLoop(dur time.Duration) {
	tick := time.NewTicker(dur)
	for {
		s.update()
		logger.Debugf("updateLoop ok")
		<-tick.C
	}
}

// search
// 只更新缓存中的热数据
func (s *Service) searchLoop(dur time.Duration) {
	tick := time.NewTicker(dur)
	client := GetClient("search")
	for {
		kv := client.GetAll()
		for k := range kv {
			s.search(k)
		}

		logger.Debugf("searchLoop len %d ok\n", len(kv))
		<-tick.C
	}
}
