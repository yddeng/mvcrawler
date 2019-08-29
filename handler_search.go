package mvcrawler

import (
	"encoding/json"
	"net/http"
)

/*
 http 搜索请求
*/

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
	if req.Txt == "" {
		logger.Errorln("search txt is nil")
		return
	}

	resp := []*Message{}
	for _, m := range s.modules {
		resp = append(resp, m.Search(req.Txt)...)

	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Errorf("search write err: %s", err)
	}
}
