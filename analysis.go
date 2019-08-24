package mvcrawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

//分析器
type Analysis struct {
	processQue     chan *AnalysisReq
	processSize    int //队列的容量
	goroutineCount int //队列消费者数量
}

//选择器
type Selector struct {
	Dom  string // DOM元素 选择器条件
	Exec []struct {
		//这一个选择器应该具体到哪一个标签
		Dom string
		//Attr获取指定属性,如果为空则获取Text
		Attr string
	}
}

type AnalysisReq struct {
	Url string
	// 选择器
	Selector *Selector
	// 异步回调
	CallBack func(resp *AnalysisReap)
}

type AnalysisReap struct {
	Url string
	// 结果集
	RespData [][]string
	// 错误信息
	Err error
}

//NewAnalysis
//size:队列的容量，goroutineCount:队列消费者数量
func NewAnalysis(size, goroutineCount int) *Analysis {
	a := &Analysis{
		processQue:     make(chan *AnalysisReq, size),
		processSize:    size,
		goroutineCount: goroutineCount,
	}

	for i := 0; i < goroutineCount; i++ {
		go a.run()
	}
	return a
}

//非阻塞投递，队列满丢弃
func (a *Analysis) Push(req *AnalysisReq) error {
	if len(a.processQue) == a.processSize {
		return fmt.Errorf("processQue is full, discard %s", req.Url)
	}
	a.processQue <- req
	return nil
}

func (a *Analysis) run() {
	for {
		req := <-a.processQue
		if req.CallBack == nil {
			//log
			continue
		}
		ret, err := a.exec(req)
		req.CallBack(&AnalysisReap{
			Url:      req.Url,
			RespData: ret,
			Err:      err,
		})
	}
}

func (a *Analysis) exec(req *AnalysisReq) (ret [][]string, err error) {
	var resp *http.Response
	resp, err = http.Get(req.Url)
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

	ret = [][]string{}
	doc.Find(req.Selector.Dom).Each(func(i int, selection *goquery.Selection) {
		msg := []string{}
		for _, sel := range req.Selector.Exec {
			var sele = selection
			var txt string
			// 内容搜集
			if sel.Dom != "" {
				sele = selection.Find(sel.Dom)
			}

			if sel.Attr != "" {
				if attr, ok := sele.Attr(sel.Attr); ok {
					txt = strings.TrimSpace(attr)
				}
			} else {
				txt = sele.Text()
			}
			msg = append(msg, txt)
		}
		ret = append(ret, msg)
	})
	return
}
