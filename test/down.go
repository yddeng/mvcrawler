package main

import (
	"fmt"
	"github.com/yddeng/mvcrawler/dhttp"
	"github.com/yddeng/mvcrawler/util"
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
