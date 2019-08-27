package main

import (
	"fmt"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
	"net/url"
)

func main() {
	data := url.Values{
		"show": {"title"}, "tbname": {"movie"}, "tempid": {"1"}, "keyboard": {"海贼王"},
	}

	resp, err := dhttp.PostUrlencoded("http://www.silisili.me/e/search/index.php", data, 0)
	if err != nil {
		fmt.Println("1", err)
	}

	//str, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("2", err)
	//}
	util.WriteFile("./", "sili.html", resp.Body)
}
