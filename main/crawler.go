package main

import (
	"fmt"
	"github.com/yddeng/mvcrawler"
	"github.com/yddeng/mvcrawler/conf"
	_ "github.com/yddeng/mvcrawler/module"
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
