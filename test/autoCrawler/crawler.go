/*
 *
 */
package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tagDong/dutil/log"
	"github.com/tagDong/dutil/queue"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
	"sync"
	"time"
)

type item struct {
	base  string
	url   string
	name  string
	depth int
}

func getItem(elem *item) []*item {
	ret := []*item{}

	resp, err := dhttp.Get(elem.url, 0)
	if err != nil {
		errLog.Errorln(elem.url, err)
		return ret
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		errLog.Errorln(elem.url, err)
		return ret
	}
	_ = resp.Body.Close()

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		var title, href, rets string
		var ok bool
		rets, _ = s.Attr("target")
		if rets == "_blank" {
			return
		}
		if href, ok = s.Attr("href"); !ok {
			return
		}
		title, ok = s.Attr("title")
		if title == "" {
			title = s.Text()
		}
		if title == "" {
			errLog.Debugln("------------", elem.url)
			return
		}
		if rets == "_self" {
			itemLog.Infoln(title, util.CheckAndInsertHead(href, "http", elem.base))
			return
		}
		ret = append(ret, &item{
			base:  elem.base,
			url:   util.CheckAndInsertHead(href, "http", elem.base),
			name:  title,
			depth: elem.depth + 1,
		})
	})
	return ret
}

var (
	crawlerLog = log.NewLogger("./", "url.log")
	errLog     = log.NewLogger("./", "err.log")
	itemLog    = log.NewLogger("./", "item.log")
	urlMap     = map[string]struct{}{}
	mu         = sync.Mutex{}
)

func main() {

	chItems := make(chan []*item, 1024)

	chItem := queue.NewEventQueue(1024, func(i interface{}) {
		item := i.(*item)
		go func() {
			ret := getItem(item)
			chItems <- ret
		}()
	})

	chItem.Run(4)
	chItem.Push(&item{name: "silisili", base: "http://www.silisili.me", url: "http://www.silisili.me", depth: 0})
	urlMap["http://www.silisili.me/"] = struct{}{}
	for v := range chItems {
		for _, i := range v {
			time.Sleep(100 * time.Millisecond)
			// 去重
			mu.Lock()
			if _, ok := urlMap[i.url]; !ok && i.depth < 3 {
				urlMap[i.url] = struct{}{}
				_ = chItem.Push(i)
				crawlerLog.Infoln(i.name, i.url, i.depth)
			}
			mu.Unlock()
		}
	}
}
