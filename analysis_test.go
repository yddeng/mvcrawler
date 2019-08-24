package mvcrawler_test

import (
	"fmt"
	"github.com/tagDong/mvcrawler"
	"net/http"
	"testing"
)

func TestNewAnalysis(t *testing.T) {
	analysis := mvcrawler.NewAnalysis(10, 2)

	_ = analysis.Post(&mvcrawler.AnalysisReq{
		Url: "http://www.silisili.me",
		Selector: &mvcrawler.Selector{
			Dom: ".time_con li",
			Exec: []struct {
				Dom  string
				Attr string
			}{
				{Dom: "p", Attr: ""},
				{Dom: "img", Attr: "src"},
			},
		},
	}, func(resp *mvcrawler.AnalysisReap) {
		fmt.Println("--------- post ------", resp.Url, resp.Err)
		for _, msg := range resp.RespData {
			for _, v := range msg {
				fmt.Println(v)
			}
		}

	})

	hreap, _ := http.Get("http://www.silisili.me")
	resp, _ := analysis.SyncPost(&mvcrawler.AnalysisReq{
		Url:      "http://www.silisili.me",
		HttpResp: hreap,
		Selector: &mvcrawler.Selector{
			Dom: ".time_con li",
			Exec: []struct {
				Dom  string
				Attr string
			}{
				{Dom: "p", Attr: ""},
				{Dom: "img", Attr: "src"},
			},
		},
	})

	fmt.Println("--------- sysn ------", resp.Url, resp.Err)
	for _, msg := range resp.RespData {
		for _, v := range msg {
			fmt.Println(v)
		}
	}

	analysis.Stop()
	select {}
}
