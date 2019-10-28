package main

import (
	"fmt"
	"github.com/yddeng/mvcrawler/dhttp"
	"github.com/yddeng/mvcrawler/util"
	"net/url"
)

func main() {
	data := url.Values{
		"wd": {"海贼王"},
	}

	resp, err := dhttp.PostUrlencoded("http://www.bimibimi.tv/vod/search/", data, 0)
	if err != nil {
		fmt.Println("1", err)
	}

	//str, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("2", err)
	//}
	util.WriteFile("./", "bimi_search_海贼王.html", resp.Body)
}
