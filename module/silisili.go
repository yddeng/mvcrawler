package module

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tagDong/dutil/log"
	"github.com/tagDong/mvcrawler"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
	"net/url"
	"strings"
)

type Silisili struct {
	name    string
	baseUrl string
	logger  *log.Logger
}

func (this *Silisili) GetName() string {
	return this.name
}

func (this *Silisili) GetUrl() string {
	return this.baseUrl
}

func (this *Silisili) Search(txt string) []*mvcrawler.Item {
	ret := []*mvcrawler.Item{}
	data := url.Values{
		"show": {"title"}, "tbname": {"movie"}, "tempid": {"1"}, "keyboard": {txt},
	}

	resp, err := dhttp.PostUrlencoded("http://www.silisili.me/e/search/index.php", data, 0)
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
func (this *Silisili) search(doc *goquery.Document, result *[]*mvcrawler.Item) {

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

func (this *Silisili) Update() [][]*mvcrawler.Item {
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

	doc.Find(".time_con").Each(func(i int, sele1 *goquery.Selection) {
		msgs := []*mvcrawler.Item{}
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

			//更新状态
			var status string
			status = sele2.Find("a i").Text()

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
	mvcrawler.Register(mvcrawler.Silisili, func(l *log.Logger) mvcrawler.Module {
		return &Silisili{
			name:    "silisili",
			baseUrl: "http://www.silisili.me",
			logger:  l,
		}
	})
}
