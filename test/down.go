package main

import (
	"fmt"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
)

func main() {
	resp, err := dhttp.Get("http://www.bimibimi.tv/", 0)
	fmt.Println(err)

	n, err := util.WriteFile("./", "bimi_search_的.html", resp.Body)
	fmt.Println(n, err)
	_ = resp.Body.Close()
}
