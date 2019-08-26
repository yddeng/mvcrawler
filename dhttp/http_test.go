package dhttp_test

import (
	"fmt"
	"github.com/tagDong/mvcrawler/dhttp"
	"io/ioutil"
	"net/url"
	"testing"
)

type Req struct {
	Show     string `json:"show"`
	Tbname   string `json:"tbname"`
	Tempid   string `json:"tempid"`
	Keyboard string `json:"keyboard"`
}

func TestHttp(t *testing.T) {
	data := make(url.Values)
	data["show"] = []string{"title"}
	data["tbname"] = []string{"movie"}
	data["tempid"] = []string{"1"}
	data["keyboard"] = []string{"海"}

	//把post表单发送给目标服务器
	resp, err := dhttp.PostUrlencoded("http://www.silisili.me/e/search/index.php", data, 0)
	result, _ := ioutil.ReadAll(resp.Body)
	content := string(result)
	fmt.Println("qq", content, err)

}
