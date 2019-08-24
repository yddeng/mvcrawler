package main

import (
	"fmt"
	"github.com/tagDong/mvcrawler"
	"github.com/tagDong/mvcrawler/conf"
	_ "github.com/tagDong/mvcrawler/module"
)

func main() {
	conf.LoadConfig("conf/conf.json")
	mvcrawler.NewService()

	fmt.Print("------------------- start ------------------")
	stop := make(chan struct{}, 1)
	select {
	case <-stop:

	}
}
