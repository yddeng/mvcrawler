/*
 *
 */
package module

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tagDong/dutil/log"
	"github.com/tagDong/mvcrawler"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
	"strings"
)

type Dilidili struct {
	name    string
	baseUrl string
	logger  *log.Logger
}

func (this *Dilidili) GetName() string {
	return this.name
}

func (this *Dilidili) GetUrl() string {
	return this.baseUrl
}

// todo
func (this *Dilidili) Search(txt string) []*mvcrawler.Item {
	ret := []*mvcrawler.Item{}
	return ret
}

// 搜索结果的分页处理
func (this *Dilidili) search(doc *goquery.Document, result *[]*mvcrawler.Item) {

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

		*result = append(*result, &mvcrawler.Item{
			Title: title,
			From:  this.GetName(),
			Img:   util.CheckAndInsertHead(img, "http", this.baseUrl),
			Url:   util.CheckAndInsertHead(url, "http", this.baseUrl),
		})
	})

	page := doc.Find(".page").Eq(0)
	page.Find("a").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		if selection.Text() == "下一页" {
			if url, ok := selection.Attr("href"); ok {

				url = strings.Replace(url, "&amp;", "&", -1)
				url = util.MergeString(this.baseUrl, url)
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

func (this *Dilidili) Update() [][]*mvcrawler.Item {
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

	doc.Find(".side .wrp .book").Each(func(i int, sele1 *goquery.Selection) {
		msgs := []*mvcrawler.Item{}
		sele1.Find(".tooltip").Each(func(_ int, sele2 *goquery.Selection) {
			var title, img, url string
			var ok bool
			if url, ok = sele2.Attr("href"); !ok {
				return
			}
			if img, ok = sele2.Find("img").Attr("src"); !ok {
				return
			}

			title = sele2.Find("figcaption p").Eq(0).Text()

			//更新状态
			var status string
			status = sele2.Find("figcaption p").Eq(1).Text()

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
	mvcrawler.Register(mvcrawler.Dilidili, func(l *log.Logger) mvcrawler.Module {
		return &Dilidili{
			name:    "dilidili",
			baseUrl: "http://www.dilidili.name",
			logger:  l,
		}
	})
}
