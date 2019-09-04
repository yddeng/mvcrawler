package mvcrawler

import (
	"github.com/tagDong/mvcrawler/conf"
	"github.com/tagDong/mvcrawler/db"
	"github.com/tagDong/mvcrawler/dhttp"
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
	db.NewClient("update", false, &UpdateDB{})
	db.NewClient("search", true, &SearchDB{})

	//module
	s.modules = map[ModuleType]Module{}
	for mt, fn := range moduleFunc {
		s.modules[mt] = fn(logger)
	}

	//httpServer
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

	go s.updateLoop(time.Duration(config.UpdateDur) * time.Second)
	go s.searchLoop(time.Duration(config.SearchDur) * time.Hour)
	return s
}

//update 抓取
func (s *Service) updateLoop(dur time.Duration) {
	tick := time.NewTicker(dur)
	updateDB := db.GetClient("update")
	for {
		data := &UpdateDB{}
		result := [][]*Message{}
		for i := 0; i < 7; i++ {
			result = append(result, []*Message{})
		}
		for _, m := range s.modules {
			ret := m.Update()
			for i, v1 := range ret {
				result[i] = append(result[i], v1...)
			}
		}
		data.Msgs = result
		updateDB.Set("update", data)

		logger.Infoln("updateLoop ok")
		<-tick.C
	}
}

// search
// 只更新缓存中的热数据
func (s *Service) searchLoop(dur time.Duration) {
	tick := time.NewTicker(dur)
	searchDB := db.GetClient("search")
	for {
		kv := searchDB.GetAll()
		for k := range kv {
			s.searchOnWeb(k)
		}

		logger.Infof("searchLoop len %d ok\n", len(kv))
		<-tick.C
	}
}
