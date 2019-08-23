package mvcrawler_test

import (
	"fmt"
	"github.com/tagDong/mvcrawler"
	"github.com/tagDong/mvcrawler/conf"
	"testing"
)

func TestNewAnalysis(t *testing.T) {
	conf.LoadConfig("conf/conf.json")

	anaConf := conf.GetConfig().Update.Urls[0]

	respCh := make(chan [][]string, 100)
	analysis := mvcrawler.NewAnalysis(respCh)

	reqs := []*mvcrawler.Request{}
	reqs = append(reqs, &mvcrawler.Request{
		Url:      anaConf.Index,
		Selector: conf.GetSilisili(),
		Depth:    1,
	})

	analysis.Push(reqs)

	for list := range respCh {
		for _, v := range list {
			for _, url := range v {
				fmt.Println(url)
			}
		}
	}
}
