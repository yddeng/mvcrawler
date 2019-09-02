package mvcrawler

import (
	"encoding/json"
	"github.com/tagDong/mvcrawler/db"
	"net/http"
)

/*
 http 搜索请求
*/

var pageNum = 20

//搜索
type SearchReq struct {
	Txt  string `json:"txt"`
	Page int    `json:"page"`
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

	data, ok := db.GetDB("search").Get(req.Txt)
	resp := &SearchRespone{Code: 0}
	var ret []*Message
	if ok {
		ret = data.([]*Message)
	} else {
		for _, m := range s.modules {
			ret = append(ret, m.Search(req.Txt)...)
		}
		db.GetDB("search").Set(req.Txt, ret)

	}

	// 分页
	length := len(ret)
	page := length / pageNum
	if length%pageNum != 0 {
		page += 1
	}

	if req.Page >= 0 && req.Page <= page {
		result := []*Message{}
		for i := req.Page * pageNum; i < (req.Page+1)*pageNum && i < length; i++ {
			result = append(result, ret[i])
		}
		resp.Code = 1
		resp.MsgNum = length
		resp.Messages = result
	}
	logger.Debugln("search respone", *resp)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Errorf("search write err: %s", err)
	}
}
