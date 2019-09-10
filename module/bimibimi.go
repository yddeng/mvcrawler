package module

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tagDong/dutil/dstring"
	"github.com/tagDong/dutil/log"
	"github.com/tagDong/mvcrawler"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
	"net/url"
)

type Bimibimi struct {
	name    string
	baseUrl string
	logger  *log.Logger
}

func (this *Bimibimi) GetName() string {
	return this.name
}

func (this *Bimibimi) GetUrl() string {
	return this.baseUrl
}

func (this *Bimibimi) Search(txt string) []*mvcrawler.Item {
	ret := []*mvcrawler.Item{}
	data := url.Values{
		"wd": {txt},
	}

	resp, err := dhttp.PostUrlencoded("http://www.bimibimi.tv/vod/search/", data, 0)
	if err != nil {
		this.logger.Errorln(err)
		return ret
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		this.logger.Errorln(err)
		return ret
	}
	this.search(doc, &ret)
	return ret
}

// 搜索结果的分页处理
func (this *Bimibimi) search(doc *goquery.Document, result *[]*mvcrawler.Item) {

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

		*result = append(*result, &mvcrawler.Item{
			Title: title,
			From:  this.GetName(),
			Img:   util.CheckAndInsertHead(img, "http", this.baseUrl),
			Url:   util.CheckAndInsertHead(url, "http", this.baseUrl),
		})
	})

	doc.Find("#long-page li a").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		if selection.Text() == "下一页" {
			if url, ok := selection.Attr("href"); ok {
				url = dstring.MergeString(this.baseUrl, url)
				this.logger.Debugln("next page", url)

				resp, err := dhttp.Get(url, 0)
				if err != nil {
					this.logger.Errorln(err)
					return false
				}

				doc, err := goquery.NewDocumentFromReader(resp.Body)
				_ = resp.Body.Close()
				if err != nil {
					this.logger.Errorln(err)
					return false
				}

				this.search(doc, result)
			}
			return false
		}
		return true
	})
}

func (this *Bimibimi) Update() [][]*mvcrawler.Item {
	ret := [][]*mvcrawler.Item{}
	resp, err := dhttp.Get(this.baseUrl, 0)
	if err != nil {
		this.logger.Errorln(err)
		return ret
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		this.logger.Errorln(err)
		return ret
	}
	_ = resp.Body.Close()

	doc.Find(".tab-content").Each(func(i int, sele1 *goquery.Selection) {
		msgs := []*mvcrawler.Item{}
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

			//更新状态
			var status string
			status = sele2.Find(".item-info p").Text()

			msgs = append(msgs, &mvcrawler.Item{
				Title:  title,
				From:   this.GetName(),
				Img:    util.CheckAndInsertHead(img, "http", this.baseUrl),
				Url:    util.CheckAndInsertHead(url, "http", this.baseUrl),
				Status: status,
			})
		})
		ret = append(ret, msgs)
	})
	return ret
}

func init() {
	mvcrawler.Register(mvcrawler.Bimibimi, func(l *log.Logger) mvcrawler.Module {
		return &Bimibimi{
			name:    "bimibimi",
			baseUrl: "http://www.bimibimi.tv",
			logger:  l,
		}
	})
}
