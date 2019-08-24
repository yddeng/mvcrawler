package mvcrawler_test

import (
	"fmt"
	"github.com/tagDong/mvcrawler"
	"testing"
)

func TestNewAnalysis(t *testing.T) {
	analysis := mvcrawler.NewAnalysis(10, 1)

	_ = analysis.Push(&mvcrawler.AnalysisReq{
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
		CallBack: func(resp *mvcrawler.AnalysisReap) {
			fmt.Println(resp.Url, resp.Err)
			for _, msg := range resp.RespData {
				for _, v := range msg {
					fmt.Println(v)
				}
			}
		},
	})
	select {}
}
