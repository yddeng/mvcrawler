/*
 * 爬取所有项目资源
 */
package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tagDong/dutil/dstring"
	"github.com/tagDong/dutil/io"
	"github.com/tagDong/dutil/log"
	"github.com/tagDong/dutil/queue"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
	"strings"
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

	crawlerLog.Infoln("crawler url", elem.url)
	resp, err := dhttp.Get(elem.url, 0)
	if err != nil {
		crawlerLog.Errorln(elem.url, err)
		return ret
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		crawlerLog.Errorln(elem.url, err)
		return ret
	}
	_ = resp.Body.Close()

	var name string
	doc.Find(".m-60").Eq(0).Each(func(i int, selection *goquery.Selection) {
		name = selection.Find("a").Eq(1).Text()
	})
	if dstring.IsEmpty(name) {
		crawlerLog.Errorln(elem.url, "name is nil")
		return ret
	}

	doc.Find("#show .xfswiper1").Each(func(i int, selection *goquery.Selection) {
		selection.Find("a").Each(func(i int, s *goquery.Selection) {
			var title, href string
			var ok bool
			if href, ok = s.Attr("href"); !ok {
				crawlerLog.Errorln(elem.url, "href is nil")
				return
			}
			title = genTitle(s.Text())
			if dstring.IsEmpty(name) {
				crawlerLog.Errorln(elem.url, "title is nil")
				return
			}

			itemTxt.WriteString(fmt.Sprintf("%s@%s@%s\n", name, title, util.CheckAndInsertHead(href, "http", elem.base)))
		})

	})
	urlTxt.WriteString(fmt.Sprintf("%s\n", elem.url))
	return ret
}

func genTitle(s string) string {
	s1 := strings.ReplaceAll(s, "\n", "")
	s1 = dstring.ContrinuousSpaceToOnce(s1)
	s1 = strings.TrimSuffix(s1, " ")
	s2 := strings.Split(s1, " ")
	ret := ""
	if len(s2) > 2 {
		for _, v := range s2[2:] {
			ret += v
		}
	} else {
		ret = s1
	}
	return ret
}

var (
	crawlerLog = log.NewLogger("./", "crawler.log")
	itemTxt    = io.NewFile("./data", "item.txt")
	urlTxt     = io.NewFile("./data", "url.txt")
	urlMap     = map[string]struct{}{}
)

func main() {

	bt, err := io.ReadFile("./data/url.txt")
	if err == nil {
		s := string(bt)
		strs := strings.Split(s, "\n")
		for _, line := range strs {
			crawlerLog.Debugln("read line", line)
			urlMap[line] = struct{}{}
		}
	}

	chItem := queue.NewEventQueue(1024, func(i interface{}) {
		item := i.(*item)
		getItem(item)
	})

	chItem.Run(1)
	for i := 1; i <= 2300; i++ {
		url := fmt.Sprintf("http://www.silisili.me/anime/%d.html", i)
		if _, ok := urlMap[url]; !ok {
			time.Sleep(1000 * time.Millisecond)
			chItem.Push(&item{name: "silisili", base: "http://www.silisili.me", url: url, depth: 0})
		}
	}
}
