package mvcrawler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//搜索
type SearchReq struct {
	Txt     string `json:"txt"`
	Modules []int  `json:"modules"`
}

func (s *Service) search(w http.ResponseWriter, r *http.Request) {
	logger.Infoln("http search request")

	var req SearchReq
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Errorf("search read err: %s", err)
		return
	}
	fmt.Print("search", req)

	for _, m := range s.modules {
		ret := m.Search(req.Txt)
		if err := json.NewEncoder(w).Encode(ret); err != nil {
			logger.Errorf("search write err: %s", err)
		}
	}
}

//日更新
type UpdateReq struct {
	Modules int `json:"modules"`
}
type UpdateResp struct {
	resp map[ModuleType][]*Message
}

var _updata *UpdateResp

func (s *Service) update(w http.ResponseWriter, r *http.Request) {
	logger.Infoln("http update request")

	if _updata.resp == nil {
		logger.Errorf("http update _update is nil")
		return
	}

	var req UpdateReq
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Errorf("update read err: %s", err)
		return
	}
	fmt.Print("update", req)

	//获取全部
	if req.Modules == 0 {
		ret := []*Message{}
		for _, v := range _updata.resp {
			for _, m := range v {
				ret = append(ret, m)
			}
		}
		if err := json.NewEncoder(w).Encode(ret); err != nil {
			logger.Errorf("update write err: %s", err)
		}
	} else {
		ret := _updata.resp[ModuleType(req.Modules)]
		if err := json.NewEncoder(w).Encode(ret); err != nil {
			logger.Errorf("update write err: %s", err)
		}
	}
}
