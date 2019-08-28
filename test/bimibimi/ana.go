package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func main() {
	hreap, _ := http.Get("http://www.bimibimi.tv")
	doc, _ := goquery.NewDocumentFromReader(hreap.Body)

	doc.Find(".tab-content").Each(func(i int, sele1 *goquery.Selection) {
		fmt.Println("i ---------------", i)
		sele1.Find(".bangumi-item").Each(func(_ int, sele2 *goquery.Selection) {
			txt, _ := sele2.Find(".item-info a").Attr("title")
			fmt.Println("title", txt)
			txt, _ = sele2.Find(".lazy-img img").Attr("src")
			fmt.Println("img", txt)
			txt, _ = sele2.Find(".item-info a").Attr("href")
			fmt.Println("href", txt)
		})
	})

}
