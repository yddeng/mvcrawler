package mvcrawler

import (
	"encoding/json"
	"net/http"
	"sync"
)

/*
 http 周更新请求
*/

func (s *Service) handleUpdate(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	logger.Infoln("handleUpdate request", r.Method, r.Form)
	//跨域
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	var ret *UpdateDB
	data, ok := GetClient("update").Get("update")
	resp := &UpdateRespone{Code: "ERR"}
	if ok {
		ret = data.(*UpdateDB)
		resp.Code = "OK"
		resp.Items = ret.Items
	}

	logger.Debugln("handleUpdate response", *resp)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Errorf("update write err: %s", err)
	}
}

func (s *Service) update() {
	client := GetClient("update")

	var webMtx sync.Mutex
	result := [][]*Item{}
	for i := 0; i < 7; i++ {
		result = append(result, []*Item{})
	}

	wg := sync.WaitGroup{}
	wg.Add(len(s.modules))
	for _, v := range s.modules {
		m := v
		go func() {
			ret := m.Update()
			webMtx.Lock()
			for i, v1 := range ret {
				result[i] = append(result[i], v1...)
			}
			webMtx.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	data := &UpdateDB{
		Items: result,
	}
	client.Set("update", data)
}
