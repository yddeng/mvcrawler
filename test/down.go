package main

import (
	"fmt"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
)

func main() {
	resp, err := dhttp.Get("http://dilidili.name", 0)
	fmt.Println(err)

	n, err := util.WriteFile("./", "dilidili.name.html", resp.Body)
	fmt.Println(n, err)
	_ = resp.Body.Close()
}
