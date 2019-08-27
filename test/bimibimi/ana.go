package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

//选择器
type Selector struct {
	Dom  string // DOM元素 选择器条件
	Exec []struct {
		//这一个Dom应该具体到哪一个标签
		Dom string
		//Attr获取指定属性,如果为空则获取Text
		Attr string
	}
	Next *Selector
}

func tt() {
	hreap, _ := http.Get("http://www.bimibimi.tv")
	doc, _ := goquery.NewDocumentFromReader(hreap.Body)

	k2 := doc.Find(".tab-cont__wrap .tab-content").Eq(0)
	fmt.Println(k2.Find(".bangumi-item .item-info a").Text())
	var s1 *goquery.Selection
	doc.Find(".tab-cont__wrap .tab-content").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		c := func() bool {
			if i == 0 {
				fmt.Println(i, selection.Find(".bangumi-item .item-info a").Text())
				s1 = selection
				return false
			}
			return true
		}
		if c() {
			return false
		}
		fmt.Println(i)
		return true
	})
	fmt.Println(s1, s1.Size())

}

func main() {
	//analysis := mvcrawler.NewAnalysis(10, 2)

	hreap, _ := http.Get("http://www.bimibimi.tv")
	doc, _ := goquery.NewDocumentFromReader(hreap.Body)

	k2 := doc.Find(".tab-cont__wrap .tab-content").Eq(0)
	fmt.Println(k2.Find(".bangumi-item .item-info a").Text())
	var s1 *goquery.Selection
	doc.Find(".tab-cont__wrap .tab-content").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		c := func() bool {
			if i == 0 {
				fmt.Println(i, selection.Find(".bangumi-item .item-info a").Text())
				s1 = selection
				return false
			}
			return true
		}
		if c() {
			return false
		}
		fmt.Println(i)
		return true
	})
	fmt.Println(s1, s1.Size())

	//s1.Find()
	/*
		resp, _ := analysis.SyncPost(&mvcrawler.AnalysisReq{
			Url:      "http://www.bimibimi.tv",
			HttpResp: hreap,
			Selector: &mvcrawler.Selector{
				Dom: ".tab-cont__wrap .item .bangumi-item",
				Exec: []struct {
					Dom  string
					Attr string
				}{
					{Dom: ".item-info a", Attr: ""},     //title
					{Dom: ".lazy-img img", Attr: "src"}, //img src
					{Dom: ".item-info a", Attr: "href"}, //url
				},
			},
		})

		fmt.Println("--------- sysn ------", resp.Url, resp.Err)
		for _, msg := range resp.RespData {
			for _, v := range msg {
				fmt.Println(v) //, util.ComperAndInsertHead(v, "https:/"))
			}
		}
	*/
}
