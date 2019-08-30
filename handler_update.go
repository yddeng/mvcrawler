package mvcrawler

import (
	"encoding/json"
	"github.com/tagDong/mvcrawler/util"
	"net/http"
	"sync"
)

/*
 http 周更新请求
*/

//
var weekDay = []int{
	6, 0, 1, 2, 3, 4, 5,
}

func GetWebWeekDay() int {
	return weekDay[util.GetWeekDay()]
}

//更新
type UpdateReq struct {
	Modules int `json:"modules"`
}

type UpdateResp struct {
	resp map[ModuleType][][]*Message
	mu   sync.Mutex
}

var _updata *UpdateResp

func (s *Service) update(w http.ResponseWriter, r *http.Request) {
	logger.Infoln("http update request", r.Method)
	_updata.mu.Lock()
	defer _updata.mu.Unlock()

	//跨域
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	if _updata.resp == nil {
		logger.Errorf("http update _update is nil")
		return
	}

	var req UpdateReq
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Errorf("json read err: %s", err)
		return
	}
	logger.Infoln("update request", req, "update data len", len(_updata.resp))
	if req.Modules >= int(End) {
		logger.Errorf("update request module:%d failed", req.Modules)
		return
	}
	//获取全部
	ret := []*Message{}
	if req.Modules == 0 {
		for k, m := range _updata.resp {
			if len(m) == 7 {
				msgs := m[GetWebWeekDay()]
				ret = append(ret, msgs...)
			} else {
				logger.Errorf("module name:%s data len:%d failed", GetName(k), len(m))
			}
		}
	} else {
		m := _updata.resp[ModuleType(req.Modules)]
		if len(m) == 7 {
			msgs := m[GetWebWeekDay()]
			ret = append(ret, msgs...)
		} else {
			logger.Errorf("module name:%s data len:%d failed", GetName(ModuleType(req.Modules)), len(m))
		}
	}

	if err := json.NewEncoder(w).Encode(ret); err != nil {
		logger.Errorf("update write err: %s", err)
	}
}
