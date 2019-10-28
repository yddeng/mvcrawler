package mvcrawler

import (
	"encoding/json"
	"github.com/yddeng/dutil/dstring"
	"net/http"
	"sync"
)

/*
 http 搜索请求
*/

var pageNum = 20

// txt  搜索内容
// page 页码

//todo msgs 的数量大于一页显示的数量 立即返回给客户端，后端继续抓取所有数据
func (s *Service) handleSearch(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	logger.Infoln("handleSearch request", r.Method, r.Form)

	//跨域
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	var txt string
	var page int
	if t, ok := r.Form["txt"]; ok {
		txt = t[0]
	}

	if p, ok := r.Form["page"]; ok {
		page = dstring.ToInt(p[0])
	}

	resp := &SearchRespone{Code: "ERR", Txt: txt, Page: page}
	if txt == "" {
		logger.Errorln("search txt is nil")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			logger.Errorf("search write err: %s", err)
		}
		return
	}

	var ret *SearchDB
	data, ok := GetClient("search").Get(txt)
	if ok {
		ret = data.(*SearchDB)
	} else {
		ret = s.search(txt)
	}

	if page >= 0 && page < ret.TotalPage {
		resp.Code = "OK"
		resp.TotalPage = ret.TotalPage
		resp.TotalItem = ret.TotalItem
		resp.Items = ret.PageItems[page]

	}

	logger.Debugln("handleSearch response", *resp)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Errorf("search write err: %s", err)
	}
}

func (s *Service) search(txt string) *SearchDB {
	logger.Infof("search txt:%s on web\n", txt)
	var webMtx sync.Mutex

	msgs := []*Item{}
	wg := sync.WaitGroup{}
	wg.Add(len(s.modules))
	for _, v := range s.modules {
		m := v
		go func() {
			ret := m.Search(txt)
			webMtx.Lock()
			msgs = append(msgs, ret...)
			webMtx.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	logger.Infof("search txt:%s on web ok\n", txt)

	// 分页
	pageMsg := [][]*Item{}
	length := len(msgs)
	var i = 0
	for i < length {
		var k = 0
		var msg = []*Item{}
		for k < pageNum && i < length {
			msg = append(msg, msgs[i])
			i++
			k++
		}
		pageMsg = append(pageMsg, msg)
	}

	sdb := &SearchDB{
		Name:      txt,
		TotalItem: length,
		TotalPage: len(pageMsg),
		PageItems: pageMsg,
	}

	logger.Debugf("SearchDB name %s, totalItem %d, totalPage %d \n", sdb.Name, sdb.TotalItem, sdb.TotalPage)

	if length > 0 {
		GetClient("search").Set(sdb.Name, sdb)
	}

	return sdb
}
