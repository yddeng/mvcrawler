package main

import (
	"fmt"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
)

func main() {
	resp, err := dhttp.Get("https://www.5dm.tv/timeline", 0)
	fmt.Println(err)

	n, err := util.WriteFile("./", "www.5dm.tv.html", resp.Body)
	fmt.Println(n, err)
	_ = resp.Body.Close()
}
