package module

import (
	"fmt"
	"github.com/tagDong/mvcrawler"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
	"net/url"
)

type Silisili struct {
	baseUrl string
	today   int
	update  *mvcrawler.Selector
	search  *mvcrawler.Selector

	analysis   *mvcrawler.Analysis
	downloader *mvcrawler.Downloader
	logger     *util.Logger
}

func (sl *Silisili) Search(txt string) []*mvcrawler.Message {
	ret := []*mvcrawler.Message{}
	data := url.Values{
		"show": {"title"}, "tbname": {"movie"}, "tempid": {"1"}, "keyboard": {txt},
	}

	//把post表单发送给目标服务器
	resp, err := dhttp.PostUrlencoded("http://www.silisili.me/e/search/index.php", data, 0)
	if err != nil {
		sl.logger.Errorf("silisili search err:%s", err)
		return ret
	}

	result, err := sl.analysis.SyncPost(&mvcrawler.AnalysisReq{
		HttpResp: resp,
		Selector: sl.search,
	})

	if err != nil {
		sl.logger.Errorf("silisili analysis syncPost err:%s", err)
		return ret
	}

	for _, msg := range result.RespData {
		ret = append(ret, &mvcrawler.Message{
			Title: msg[0],
			From:  "silisili",
			Img:   msg[1],
			Url:   util.MergeString(sl.baseUrl, msg[2]),
		})
	}
	return ret
}

func (sl *Silisili) Update() []*mvcrawler.Message {

	if util.GetWeekDay() != sl.today {
		sl.update = updateSilisili()
		sl.today = util.GetWeekDay()
	}

	ret := []*mvcrawler.Message{}
	result, err := sl.analysis.SyncPost(&mvcrawler.AnalysisReq{
		Url:      sl.baseUrl,
		Selector: sl.update,
	})

	if err != nil {
		sl.logger.Errorf("silisili analysis syncPost err:%s", err)
		return ret
	}

	for _, msg := range result.RespData {
		ret = append(ret, &mvcrawler.Message{
			Title: msg[0],
			From:  "silisili",
			Img:   msg[1],
			Url:   util.MergeString(sl.baseUrl, msg[2]),
		})
	}
	return ret
}

// silisili日更新
func updateSilisili() *mvcrawler.Selector {
	var siliWeek = []int{
		6, 0, 1, 2, 3, 4, 5,
	}

	n := util.GetWeekDay()
	dom := fmt.Sprintf(".xfswiper%d li", siliWeek[n])

	return &mvcrawler.Selector{
		Dom: dom,
		Exec: []struct {
			Dom  string
			Attr string
		}{
			{Dom: "p", Attr: ""},      //title
			{Dom: "img", Attr: "src"}, //img src
			{Dom: "a", Attr: "href"},  //url
		},
	}
}

func searchSilisili() *mvcrawler.Selector {
	return &mvcrawler.Selector{
		Dom: ".anime_list dl",
		Exec: []struct {
			Dom  string
			Attr string
		}{
			{Dom: "dd h3 a", Attr: ""},     //title
			{Dom: "dt img", Attr: "src"},   //img src
			{Dom: "dd h3 a", Attr: "href"}, //url
		},
	}
}

func init() {
	mvcrawler.Register(mvcrawler.Silisili, func(
		anal *mvcrawler.Analysis, down *mvcrawler.Downloader, l *util.Logger) mvcrawler.Module {

		return &Silisili{
			baseUrl: "http://www.silisili.me",
			today:   util.GetWeekDay(),
			update:  updateSilisili(),
			search:  searchSilisili(),

			analysis:   anal,
			downloader: down,
			logger:     l,
		}
	})
}
