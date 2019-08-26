package module

import (
	"github.com/tagDong/mvcrawler"
	"github.com/tagDong/mvcrawler/util"
	"testing"
)

func TestSilisili_Search(t *testing.T) {
	anal := mvcrawler.NewAnalysis(10, 2)
	l := util.NewLogger("log", "silisili")

	s := &Silisili{
		baseUrl: "www.silisili.me",
		update:  updateSilisili(),
		search:  searchSilisili(),

		analysis: anal,
		logger:   l,
	}

	ret := s.Search("海")
	for _, m := range ret {
		l.Infoln(m.Title, m.Img, m.Url)
	}
}
