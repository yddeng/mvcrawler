package module

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tagDong/mvcrawler"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
	"net/url"
)

type Bimibimi struct {
	baseUrl string
	logger  *util.Logger
}

func (sl *Bimibimi) GetUrl() string {
	return sl.baseUrl
}

func (sl *Bimibimi) Search(txt string) []*mvcrawler.Message {
	ret := []*mvcrawler.Message{}
	data := url.Values{
		"wd": {txt},
	}

	resp, err := dhttp.PostUrlencoded("http://www.bimibimi.tv/vod/search/", data, 0)
	if err != nil {
		sl.logger.Errorln(err)
		return ret
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		sl.logger.Errorln(err)
		return ret
	}
	_ = resp.Body.Close()

	doc.Find(".v_tb .item").Each(func(i int, selection *goquery.Selection) {
		var title, img, url string
		var ok bool
		title = selection.Find(".info a").Text()
		if img, ok = selection.Find("img").Attr("data-original"); !ok {
			return
		}
		if url, ok = selection.Find(".info a").Attr("href"); !ok {
			return
		}

		ret = append(ret, &mvcrawler.Message{
			Title: title,
			From:  "bimibimi",
			Img:   util.MergeString(sl.baseUrl, img),
			Url:   util.MergeString(sl.baseUrl, url),
		})
	})
	return ret
}

func (sl *Bimibimi) Update() [][]*mvcrawler.Message {
	ret := [][]*mvcrawler.Message{}
	resp, err := dhttp.Get(sl.baseUrl, 0)
	if err != nil {
		sl.logger.Errorln(err)
		return ret
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		sl.logger.Errorln(err)
		return ret
	}
	_ = resp.Body.Close()

	doc.Find(".tab-content").Each(func(i int, sele1 *goquery.Selection) {
		msgs := []*mvcrawler.Message{}
		sele1.Find(".bangumi-item").Each(func(_ int, sele2 *goquery.Selection) {
			var title, img, url string
			var ok bool
			if title, ok = sele2.Find(".item-info a").Attr("title"); !ok {
				return
			}
			if img, ok = sele2.Find(".lazy-img img").Attr("src"); !ok {
				return
			}
			if url, ok = sele2.Find(".item-info a").Attr("href"); !ok {
				return
			}

			msgs = append(msgs, &mvcrawler.Message{
				Title: title,
				From:  "bimibimi",
				Img:   util.MergeString(sl.baseUrl, img),
				Url:   util.MergeString(sl.baseUrl, url),
			})
		})
		ret = append(ret, msgs)
	})
	return ret
}

func init() {
	mvcrawler.Register(mvcrawler.Bimibimi, func(l *util.Logger) mvcrawler.Module {
		return &Bimibimi{
			baseUrl: "http://www.bimibimi.tv",
			logger:  l,
		}
	})
}
