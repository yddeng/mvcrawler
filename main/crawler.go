package main

import (
	"fmt"
	"github.com/tagDong/mvcrawler"
	"github.com/tagDong/mvcrawler/conf"
	_ "github.com/tagDong/mvcrawler/module"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Printf("usage config\n")
		return
	}

	fmt.Println("------------------- start ------------------")
	conf.LoadConfig(os.Args[1])
	mvcrawler.NewService()

	stop := make(chan struct{}, 1)
	select {
	case <-stop:

	}
}
