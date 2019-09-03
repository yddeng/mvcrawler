/*
 *
 */
package main

import (
	"fmt"
	"github.com/tagDong/mvcrawler"
)

func main() {
	data := []*mvcrawler.Message{{Title: "1", Status: "1"}, {Title: "1", Status: "2"}, {Title: "2", Status: "3"}}
	for _, v := range data {
		fmt.Println(*v)
	}
	mvcrawler.Process(&data)
	for _, v := range data {
		fmt.Println(*v)
	}
}
