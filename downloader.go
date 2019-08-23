package mvcrawler

import (
	"github.com/tagDong/mvcrawler/conf"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
	"path"
	"strings"
)

type Downloader struct {
	downloadPath   string
	goroutineCount int
	downloadQueue  chan *Node
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
		downloadQueue:  make(chan *Node, downConf.ChanSize),
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
func (d *Downloader) Push(node *Node) {
	if len(d.downloadQueue) == d.downloadSize {
		logger.Errorf("downloadQueue is full, discard %s %s", node.Name, node.Url)
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
				logger.Errorf("download url:%s err:%s", node.Url, err)
			} else {
				logger.Debugf("download url:%s size:%s\n", node.Url, util.SiezToString(n))
			}
		}
	}
}

/*
* 作用：下载资源
* 参数：资源的URL地址
* 返回值：大小，错误
 */
func (d *Downloader) download(node *Node) (n int64, err error) {

	resp, err := dhttp.Get(node.Url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	//文件的类型，作为存储目录
	context := resp.Header.Get("Content-Type")
	nextpath := strings.Split(context, ";")[0]

	return util.WriteFile(path.Join(d.downloadPath, nextpath), node.Name, resp.Body)
	//return 0, nil
}

type Node struct {
	Name string
	Url  string
}

func NewNode(url, name string) *Node {

	if name == "" {
		s := strings.Split(url, "/")
		name = s[len(s)-1]
	}

	return &Node{
		Name: name,
		Url:  url,
	}
}
