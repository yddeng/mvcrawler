package mvcrawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tagDong/mvcrawler/conf"
	"github.com/tagDong/mvcrawler/dhttp"
	"net/http"
	"strings"
)

//分析器
type Analysis struct {
	processQue  chan []*Request
	processSize int
}

type Request struct {
	Url string
	// 选择器
	Selector *conf.Selector
	// 请求深度
	Depth int
}

func NewAnalysis(chResp chan<- [][]string) *Analysis {
	analConf := conf.GetConfig().Common.Analysis
	a := &Analysis{
		processQue:  make(chan []*Request, analConf.ChanSize),
		processSize: analConf.ChanSize,
	}

	for i := 0; i < analConf.GoroutineCount; i++ {
		go a.run(chResp)
	}

	return a
}

func (a *Analysis) Push(req []*Request) {
	if len(a.processQue) == a.processSize {
		logger.Errorf("processQue is full, discard")
	} else {
		a.processQue <- req
	}
}

func (a *Analysis) run(chResp chan<- [][]string) {
	for {
		reqs := <-a.processQue
		for _, req := range reqs {
			ret, err := a.exec(req)
			if err != nil {
				logger.Errorf("analysis: err %s", err)
			} else {
				resp := [][]string{}
				for _, msg := range ret {
					resp = append(resp, msg.data)
				}
				chResp <- resp
			}
		}
	}
}

type message struct {
	depth int
	data  []string
}

func (a *Analysis) exec(req *Request) (ret []*message, err error) {

	var resp *http.Response
	resp, err = dhttp.Get(req.Url)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http %s StatusCode %d", req.Url, resp.StatusCode)
		return
	}

	defer resp.Body.Close()
	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	ret = []*message{}
	doc.Find(req.Selector.Dom).Each(func(i int, selection *goquery.Selection) {
		msg := new(message)
		msg.depth = req.Depth + 1
		for _, sel := range req.Selector.Exec {
			var s = selection
			var txt string
			// 内容搜集
			if sel.Dom != "" {
				s = selection.Find(sel.Dom)
			}

			if sel.Attr != "" {
				if attr, ok := s.Attr(sel.Attr); ok {
					txt = strings.TrimSpace(attr)
				}
			} else {
				txt = s.Text()
			}
			msg.data = append(msg.data, txt)
		}
		ret = append(ret, msg)
	})

	return
}
