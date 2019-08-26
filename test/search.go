package main

import (
	"fmt"
	"github.com/tagDong/mvcrawler/dhttp"
	"io/ioutil"
)

type SearchReq struct {
	Txt string `json:"txt"`
}

func main() {
	resp, err := dhttp.PostJson("http://127.0.0.1:2369/search", SearchReq{"海贼王"}, 0)
	if err != nil {
		fmt.Println("1", err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("2", err)
	}
	fmt.Println(string(data))
}
