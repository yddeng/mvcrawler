package analysis_test

import (
	"fmt"
	analysis2 "github.com/tagDong/mvcrawler/core/analysis"
	"net/http"
	"testing"
)

func TestNewAnalysis(t *testing.T) {
	analysis := analysis2.NewAnalysis(10, 2)

	_ = analysis.Post(&analysis2.AnalysisReq{
		Url: "http://www.silisili.me",
		Selector: &analysis2.Selector{
			Dom: ".time_con li",
			Exec: []struct {
				Dom  string
				Attr string
			}{
				{Dom: "p", Attr: ""},
				{Dom: "img", Attr: "src"},
			},
		},
	}, func(resp *analysis2.AnalysisReap) {
		fmt.Println("--------- post ------", resp.Url, resp.Err)
		for _, msg := range resp.RespData {
			for _, v := range msg {
				fmt.Println(v)
			}
		}

	})

	hreap, _ := http.Get("http://www.silisili.me")
	resp, _ := analysis.SyncPost(&analysis2.AnalysisReq{
		Url:      "http://www.silisili.me",
		HttpResp: hreap,
		Selector: &analysis2.Selector{
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
