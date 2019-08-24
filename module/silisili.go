package module

import (
	"fmt"
	"github.com/tagDong/mvcrawler"
	"github.com/tagDong/mvcrawler/conf"
	"github.com/tagDong/mvcrawler/util"
)

type Silisili struct {
	baseUrl string
	update  *conf.Selector
	search  *conf.Selector
}

func (sl *Silisili) Search(txt string) []*mvcrawler.Message {
	ret := []*mvcrawler.Message{}

	return ret
}

func (sl *Silisili) Update() {

}

// silisili日更新
func updateSilisili() *conf.Selector {
	var siliWeek = []int{
		6, 0, 1, 2, 3, 4, 5,
	}

	n := util.GetWeekDay()
	dom := fmt.Sprintf(".xfswiper%d li", siliWeek[n])

	return &conf.Selector{
		Dom: dom,
		Exec: []struct {
			Dom  string
			Attr string
		}{
			{Dom: "p", Attr: ""},
			{Dom: "img", Attr: "src"},
		},
	}
}

func searchSilisili() *conf.Selector {
	return &conf.Selector{
		Dom: ".anime_list dl",
		Exec: []struct {
			Dom  string
			Attr string
		}{
			{Dom: "dd h3 a", Attr: ""},
			{Dom: "dt img", Attr: "src"},
			{Dom: "dd h3 a", Attr: "href"},
		},
	}
}

func init() {
	mvcrawler.Register(mvcrawler.Silisili, func() mvcrawler.Module {

		return &Silisili{
			baseUrl: "www.silisili.me",
			update:  updateSilisili(),
			search:  searchSilisili(),
		}
	})
}
