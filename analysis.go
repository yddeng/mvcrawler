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
	chStop         chan struct{}
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
	callBack func(resp *AnalysisReap)
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
		chStop:         make(chan struct{}),
	}

	//启动相应的处理线程
	for i := 1; i <= goroutineCount; i++ {
		id := i
		go a.run(id)
	}
	return a
}

//非阻塞异步投递，队列满丢弃
func (a *Analysis) Post(req *AnalysisReq, callback func(resp *AnalysisReap)) error {
	if len(a.processQue) == a.processSize {
		return fmt.Errorf("processQue is full, discard %s", req.Url)
	}

	if callback == nil {
		return fmt.Errorf("callback is nil")
	}

	req.callBack = callback
	a.processQue <- req
	return nil
}

//同步投递
func (a *Analysis) SyncPost(req *AnalysisReq) (resp *AnalysisReap, err error) {
	ch := make(chan *AnalysisReap)
	if err = a.Post(req, func(resp *AnalysisReap) {
		ch <- resp
	}); err != nil {
		return
	}
	resp = <-ch
	return resp, err
}

func (a *Analysis) Stop() {
	close(a.chStop)
}

func (a *Analysis) run(id int) {
	fmt.Printf("analysis consumer(%d) run\n", id)
	for {
		select {
		case <-a.chStop:
			fmt.Printf("analysis consumer(%d) close\n", id)
			return
		case req := <-a.processQue:
			ret, err := a.exec(req)
			req.callBack(&AnalysisReap{
				Url:      req.Url,
				RespData: ret,
				Err:      err,
			})
		}
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
