package main

import (
	"fmt"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
)

func main() {
	resp, err := dhttp.Get("http://www.silisili.me/anime/998.html", 0)
	fmt.Println(err)

	length := resp.ContentLength
	fmt.Println("length", length)
	n, err := util.WriteFile("./silisili/", "998.html", resp.Body)
	fmt.Println(n, err)
	_ = resp.Body.Close()
}
