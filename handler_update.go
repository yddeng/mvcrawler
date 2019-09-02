package mvcrawler

import (
	"encoding/json"
	"github.com/tagDong/mvcrawler/db"
	"github.com/tagDong/mvcrawler/util"
	"net/http"
)

/*
 http 周更新请求
*/

//更新
type UpdateReq struct {
	Modules int `json:"modules"`
}

func (s *Service) update(w http.ResponseWriter, r *http.Request) {
	logger.Infoln("http update request", r.Method)

	//跨域
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	var req UpdateReq
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Errorf("json read err: %s", err)
		return
	}
	logger.Infoln("update request", req)

	data, ok := db.GetDB("update").Get("update")
	resp := &UpdateRespone{Code: 0}
	if ok {
		resp.Code = 1
		resp.Messages = data.([][]*Message)
	}
	logger.Debugln("update respone", *resp)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Errorf("update write err: %s", err)
	}
}

//
var weekDay = []int{
	6, 0, 1, 2, 3, 4, 5,
}

func GetWebWeekDay() int {
	return weekDay[util.GetWeekDay()]
}
