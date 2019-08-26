package mvcrawler

import (
	"fmt"
	"github.com/tagDong/mvcrawler/dhttp"
	"github.com/tagDong/mvcrawler/util"
	"net/http"
	"path"
	"strings"
)

type Downloader struct {
	downloadPath   string
	downloadQueue  chan *DownloadReq
	downloadSize   int
	goroutineCount int
	logger         *util.Logger
}

type DownloadReq struct {
	Name string
	Url  string
}

func (req *DownloadReq) GetName() string {
	name := req.Name
	if name == "" {
		s := strings.Split(req.Url, "/")
		name = s[len(s)-1]
	}
	return name
}

// NewDownLoader
// outPath:文件输出目录
// size:队列的容量，goroutineCount:队列消费者数量
// logger:日志
func NewDownLoader(outPath string, size, goroutineCount int, logger *util.Logger) *Downloader {
	d := &Downloader{
		downloadPath:   outPath,
		downloadQueue:  make(chan *DownloadReq, size),
		downloadSize:   size,
		goroutineCount: goroutineCount,
		logger:         logger,
	}

	for i := 0; i < goroutineCount; i++ {
		go d.run()
	}
	return d
}

//非阻塞投递，队列满丢弃
func (d *Downloader) Push(node *DownloadReq) error {
	if len(d.downloadQueue) == d.downloadSize {
		return fmt.Errorf("downloadQueue is full, discard %s %s", node.Name, node.Url)
	}
	d.downloadQueue <- node
	return nil
}

func (d *Downloader) run() {
	for {
		node := <-d.downloadQueue
		n, err := d.download(node)
		if err != nil {
			d.logger.Errorf("download err:%s", err)
		} else {
			d.logger.Debugf("download url:%s size:%s\n", node.Url, util.SiezToString(n))
		}
	}
}

/*
* 作用：下载资源
* 参数：资源的URL地址
* 返回值：大小，错误
 */
func (d *Downloader) download(req *DownloadReq) (n int64, err error) {
	var resp *http.Response
	resp, err = dhttp.Get(req.Url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	//文件的类型，作为存储目录
	context := resp.Header.Get("Content-Type")
	nextpath := strings.Split(context, ";")[0]

	return util.WriteFile(path.Join(d.downloadPath, nextpath), req.GetName(), resp.Body)
}
