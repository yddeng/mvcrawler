package mvcrawler

import (
	"encoding/json"
	"net/http"
)

//搜索
type SearchReq struct {
	Txt string `json:"txt"`
	//Modules []int  `json:"modules"`
}

func (s *Service) search(w http.ResponseWriter, r *http.Request) {
	logger.Infoln("http search request", r.Method)

	//跨域
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	var req SearchReq
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Errorf("json err: %s", err)
		return
	}
	logger.Infoln("search request", req)

	resp := []*Message{}
	for _, m := range s.modules {
		resp = append(resp, m.Search(req.Txt)...)

	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Errorf("search write err: %s", err)
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
	logger.Infoln("http update request", r.Method)

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
	logger.Infoln("update request", req)

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
