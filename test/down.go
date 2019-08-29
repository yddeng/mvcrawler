package main

import (
	"fmt"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
)

func main() {
	resp, err := dhttp.Get("http://10.128.2.252:2323", 0)
	fmt.Println(err)

	n, err := util.WriteFile("./", "index.html", resp.Body)
	fmt.Println(n, err)
	_ = resp.Body.Close()
}
