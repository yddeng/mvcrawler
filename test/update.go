package main

import (
	"fmt"
	"github.com/tagDong/mvcrawler/dhttp"
	"io/ioutil"
)

type UpdateReq struct {
	Modules int `json:"modules"`
}

func main() {
	resp, err := dhttp.PostJson("http://127.0.0.1:12345/update", UpdateReq{4}, 0)
	//resp, err := dhttp.PostJson("http://104.168.165.226:12345/update", UpdateReq{}, 0)
	if err != nil {
		fmt.Println("1", err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("2", err)
	}
	fmt.Println(string(data))
}
