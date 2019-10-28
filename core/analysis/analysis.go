package analysis

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/yddeng/mvcrawler/dhttp"
	"net/http"
	"strings"
	"sync/atomic"
)

var (
	started = int32(1)
	stopped = int32(0)
)

//分析器
type Analysis struct {
	processQue     chan *AnalysisReq
	processSize    int   //队列的容量
	goroutineCount int   //队列消费者数量
	flag           int32 //状态
	chStop         chan struct{}
}

//选择器
type Selector struct {
	Dom  string // DOM元素 选择器条件
	Exec []struct {
		//这一个Dom应该具体到哪一个标签
		Dom string
		//Attr获取指定属性,如果为空则获取Text
		Attr string
	}
}

//分析请求
type AnalysisReq struct {
	Url      string
	HttpResp *http.Response
	// 选择器
	Selector *Selector
	// 异步回调
	callBack func(resp *AnalysisReap)
}

//分析结果
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
	atomic.StoreInt32(&a.flag, started)

	//启动相应的处理线程
	for i := 1; i <= goroutineCount; i++ {
		id := i
		go a.run(id)
	}
	return a
}

//非阻塞异步投递，队列满丢弃
func (a *Analysis) Post(req *AnalysisReq, callback func(resp *AnalysisReap)) error {
	if atomic.LoadInt32(&a.flag) == stopped {
		return fmt.Errorf("analysis is stopped")
	}
	if len(a.processQue) == a.processSize {
		return fmt.Errorf("analysis processQue is full, discard %s", req.Url)
	}

	if callback == nil {
		return fmt.Errorf("analysis post callback is nil")
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

//停止分析器
func (a *Analysis) Stop() {
	if atomic.CompareAndSwapInt32(&a.flag, started, stopped) {
		close(a.chStop)
	}
}

func (a *Analysis) run(id int) {
	fmt.Printf("analysis consumer(%d) run\n", id)
	for {
		select {
		case <-a.chStop:
			fmt.Printf("analysis consumer(%d) stop\n", id)
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

// 页面分析处理
// 如果HttpResp为空，先请求
func (a *Analysis) exec(req *AnalysisReq) (ret [][]string, err error) {
	if req.HttpResp == nil {
		req.HttpResp, err = dhttp.Get(req.Url, 0)
		if err != nil {
			return
		}
	}
	resp := req.HttpResp

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
