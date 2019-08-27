package module

import (
	"github.com/tagDong/mvcrawler"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
	"net/url"
)

type Bimibimi struct {
	baseUrl string
	today   int
	update  *mvcrawler.Selector
	search  *mvcrawler.Selector

	analysis   *mvcrawler.Analysis
	downloader *mvcrawler.Downloader
	logger     *util.Logger
}

func (sl *Bimibimi) Search(txt string) []*mvcrawler.Message {
	ret := []*mvcrawler.Message{}
	data := url.Values{
		"wd": {txt},
	}

	//把post表单发送给目标服务器
	resp, err := dhttp.PostUrlencoded("http://www.bimibimi.tv/vod/search/", data, 0)
	if err != nil {
		sl.logger.Errorf("bimibimi search err:%s", err)
		return ret
	}

	result, err := sl.analysis.SyncPost(&mvcrawler.AnalysisReq{
		HttpResp: resp,
		Selector: sl.search,
	})

	if err != nil {
		sl.logger.Errorf("bimibimi analysis syncPost err:%s", err)
		return ret
	}

	for _, msg := range result.RespData {
		ret = append(ret, &mvcrawler.Message{
			Title: msg[0],
			From:  "bimibimi",
			Img:   util.MergeString(sl.baseUrl, msg[1]),
			Url:   util.MergeString(sl.baseUrl, msg[2]),
		})
	}
	return ret
}

func (sl *Bimibimi) Update() []*mvcrawler.Message {
	ret := []*mvcrawler.Message{}

	return ret
}

// silisili日更新
func updateBimibimi() *mvcrawler.Selector {
	return &mvcrawler.Selector{
		Dom: ".tab-cont__wrap .item .bangumi-item",
		Exec: []struct {
			Dom  string
			Attr string
		}{
			{Dom: ".item-info a", Attr: ""},     //title
			{Dom: ".lazy-img img", Attr: "src"}, //img src
			{Dom: ".item-info a", Attr: "href"}, //url
		},
	}
}

func searchBimibimi() *mvcrawler.Selector {
	return &mvcrawler.Selector{
		Dom: ".v_tb .item",
		Exec: []struct {
			Dom  string
			Attr string
		}{
			{Dom: ".info a", Attr: ""},          //title
			{Dom: "img", Attr: "data-original"}, //img src
			{Dom: ".info a", Attr: "href"},      //url
		},
	}
}

func init() {
	mvcrawler.Register(mvcrawler.Bimibimi, func(
		anal *mvcrawler.Analysis, down *mvcrawler.Downloader, l *util.Logger) mvcrawler.Module {

		return &Bimibimi{
			baseUrl: "http://www.bimibimi.tv",
			today:   util.GetWeekDay(),
			update:  updateBimibimi(),
			search:  searchBimibimi(),

			analysis:   anal,
			downloader: down,
			logger:     l,
		}
	})
}
