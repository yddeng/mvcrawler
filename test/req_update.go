package main

import (
	"fmt"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
	"net/url"
)

func main() {
	data := url.Values{
		"wd": {"海贼"},
	}

	resp, err := dhttp.PostUrlencoded("http://www.bimibimi.tv/vod/search/", data, 0)
	if err != nil {
		fmt.Println("1", err)
	}

	//str, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("2", err)
	//}
	util.WriteFile("./", "bili.html", resp.Body)
}
