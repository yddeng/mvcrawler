package mvcrawler

import (
	"github.com/tagDong/crawler/util"
	"github.com/tagDong/mvcrawler/conf"
	"net/http"
	"path"
	"strings"
)

type Downloader struct {
	downloadPath   string
	goroutineCount int
	downloadQueue  chan Node
	downloadSize   int
}

/*
 一个下载器
 有多个消费者处理队列下载事件
*/
func NewDownLoader() *Downloader {

	downConf := conf.GetConfig().Common.DownLoad

	d := &Downloader{
		downloadPath:   downConf.OutPath,
		goroutineCount: downConf.GoroutineCount,
		downloadQueue:  make(chan Node, downConf.ChanSize),
		downloadSize:   downConf.ChanSize,
	}

	for i := 0; i < downConf.GoroutineCount; i++ {
		go d.run()
	}

	return d
}

/*
 像队列中添加下载事件
 队列满，丢弃
*/
func (d *Downloader) Push(node Node) {
	if len(d.downloadQueue) == d.downloadSize {
		logger.Errorf("downloadQueue is full, discard %s %s", node.name, node.url)
	} else {
		d.downloadQueue <- node
	}
}

func (d *Downloader) run() {
	for {
		select {
		case node := <-d.downloadQueue:
			n, err := d.download(node)
			if err != nil {
				logger.Errorf("url: %s, err: %s", node.url, err)
			} else {
				logger.Debugln("type: %s, size: %d, url: %s", node.GetTTString(), n, node.url)
			}
		}
	}
}

/*
* 作用：下载资源
* 参数：资源的URL地址
* 返回值：大小，错误
 */
func (d *Downloader) download(node Node) (n int64, err error) {

	resp, err := http.Get(node.url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	return util.WriteFile(path.Join(d.downloadPath, node.tt.ToString()), node.name, resp.Body)

}

/****************************************************************************/

//下载文件类型
type DownType int

const (
	ImageType DownType = iota
	VideoType
	HtmlType
)

var TypeString = [...]string{
	"image",
	"video",
	"html",
}

type Node struct {
	tt   DownType
	name string
	url  string
}

func NewNode(url, name string, t DownType) Node {

	if name == "" {
		s := strings.Split(url, "/")
		name = s[len(s)-1]
	}

	return Node{
		tt:   t,
		name: name,
		url:  url,
	}
}

func (t DownType) ToString() string {
	return TypeString[t]
}

func (n Node) GetTTString() string {
	return n.tt.ToString()
}

func (n Node) GetName() string {
	return n.name
}

func (n Node) GetUrl() string {
	return n.url
}
