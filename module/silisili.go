package module

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tagDong/mvcrawler"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
	"net/url"
)

type Silisili struct {
	baseUrl string
	logger  *util.Logger
}

func (sl *Silisili) GetUrl() string {
	return sl.baseUrl
}

func (sl *Silisili) Search(txt string) []*mvcrawler.Message {
	ret := []*mvcrawler.Message{}
	data := url.Values{
		"show": {"title"}, "tbname": {"movie"}, "tempid": {"1"}, "keyboard": {txt},
	}

	resp, err := dhttp.PostUrlencoded("http://www.silisili.me/e/search/index.php", data, 0)
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

	doc.Find(".anime_list dl").Each(func(i int, selection *goquery.Selection) {
		var title, img, url string
		var ok bool
		title = selection.Find("dd h3 a").Text()
		if img, ok = selection.Find("dt img").Attr("src"); !ok {
			return
		}
		if url, ok = selection.Find("dd h3 a").Attr("href"); !ok {
			return
		}

		ret = append(ret, &mvcrawler.Message{
			Title: title,
			From:  "silisili",
			Img:   img,
			Url:   util.MergeString(sl.baseUrl, url),
		})
	})

	return ret
}

func (sl *Silisili) Update() [][]*mvcrawler.Message {
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

	doc.Find(".time_con").Each(func(i int, sele1 *goquery.Selection) {
		msgs := []*mvcrawler.Message{}
		sele1.Find("li").Each(func(_ int, sele2 *goquery.Selection) {
			var title, img, url string
			var ok bool
			if title, ok = sele2.Find("a").Attr("title"); !ok {
				return
			}
			if img, ok = sele2.Find("img").Attr("src"); !ok {
				return
			}
			if url, ok = sele2.Find("a").Attr("href"); !ok {
				return
			}

			msgs = append(msgs, &mvcrawler.Message{
				Title: title,
				From:  "silisili",
				Img:   img,
				Url:   util.MergeString(sl.baseUrl, url),
			})
		})
		ret = append(ret, msgs)
	})
	return ret
}

func init() {
	mvcrawler.Register(mvcrawler.Silisili, func(l *util.Logger) mvcrawler.Module {
		return &Silisili{
			baseUrl: "http://www.silisili.me",
			logger:  l,
		}
	})
}
