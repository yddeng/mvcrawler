package mvcrawler

import (
	"encoding/json"
	"github.com/tagDong/mvcrawler/db"
	"net/http"
	"sync"
)

/*
 http 搜索请求
*/

var pageNum = 20

// 存储结构
type SearchDB struct {
	Name    string       //搜索字
	MsgNum  int          //结果数量
	PageNum int          //分页数量
	PageMsg [][]*Message //分页后的项目集合
}

//搜索
type SearchReq struct {
	Txt  string `json:"txt"`
	Page int    `json:"page"`
}

//todo msgs 的数量大于一页显示的数量 立即返回给客户端，后端继续抓取所有数据
func (s *Service) search(w http.ResponseWriter, r *http.Request) {
	//logger.Infoln("http search request", r.Method)

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

	resp := &SearchRespone{Code: 0}
	if req.Txt == "" {
		logger.Errorln("search txt is nil")

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			logger.Errorf("search write err: %s", err)
		}
		return
	}

	var ret *SearchDB
	data, ok := db.GetClient("search").Get(req.Txt)
	if ok {
		ret = data.(*SearchDB)
	} else {
		ret = s.searchOnWeb(req.Txt)
	}

	if req.Page >= 0 && req.Page < ret.PageNum {
		resp.Code = 1
		resp.PageNum = ret.PageNum
		resp.Messages = ret.PageMsg[req.Page]

	}

	logger.Debugln("search respone", *resp)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Errorf("search write err: %s", err)
	}
}

var webMtx sync.Mutex

func (s *Service) searchOnWeb(txt string) *SearchDB {
	logger.Infof("search txt:%s on web\n", txt)

	msgs := []*Message{}
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
	pageMsg := [][]*Message{}
	length := len(msgs)
	var i = 0
	for i < length {
		var k = 0
		var msg = []*Message{}
		for k < pageNum && i < length {
			msg = append(msg, msgs[i])
			i++
			k++
		}
		pageMsg = append(pageMsg, msg)
	}

	sdb := &SearchDB{
		Name:    txt,
		MsgNum:  length,
		PageNum: len(pageMsg),
		PageMsg: pageMsg,
	}

	logger.Debugf("sdb name %s, page_num %d, msg_num %d \n", sdb.Name, sdb.PageNum, sdb.MsgNum)

	if length > 0 {
		db.GetClient("search").Set(sdb.Name, sdb)
	}

	return sdb
}
