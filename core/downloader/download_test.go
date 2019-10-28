package downloader_test

import (
	"github.com/yddeng/mvcrawler/core/downloader"
	"github.com/yddeng/mvcrawler/util"
	"testing"
)

func TestNewDownLoader(t *testing.T) {
	logger := util.NewLogger("log", "download")
	d := downloader.NewDownLoader("download", 10, 1, logger)

	_ = d.Push(&downloader.DownloadReq{Url: "http://www.silisili.me", Name: "www.silisili.me.html"})
	_ = d.Push(&downloader.DownloadReq{Url: "http://www.bimibimi.tv", Name: "www.bimibimi.tv.html"})
	_ = d.Push(&downloader.DownloadReq{Url: "http://www.halihali.in", Name: "www.halihali.in.html"})
	_ = d.Push(&downloader.DownloadReq{Url: "http://www.dililitv.com", Name: "www.dililitv.com.html"})
	//_ = d.Push(&downloader.DownloadReq{Url: "http://www.jbhua.com/uploads/150617/1-15061g01u5933.jpg", Name: "1-15061g01u5933.jpg"})
	//_ = d.Push(&downloader.DownloadReq{Url: "http://b-ssl.duitang.com/uploads/blog/201308/06/20130806222801_Tr38Z.jpeg", Name: "20130806222801_Tr38Z.jpeg"})
	//_ = d.Push(&downloader.DownloadReq{Url: "http://photocdn.sohu.com/20150724/mp24129102_1437711995584_2.gif"})
	//_ = d.Push(&downloader.DownloadReq{Url: "http://b-ssl.duitang.com/uploads/item/201411/08/20141108074440_3dFfP.jpeg"})
	//_ = d.Push(&downloader.DownloadReq{Url: "https://gss3.baidu.com/6LZ0ej3k1Qd3ote6lo7D0j9wehsv/tieba-smallvideo/3_cda6388ad4f3a1d9db9fd0f942af406d.mp4",Name: "龙珠英雄1.mp4"})

	//_ = d.Push(&downloader.DownloadReq{Url: "http://www.silisili.me/e/search/result/?searchid=1977", Name: "s.html"})
	//_ = d.Push(&downloader.DownloadReq{Url: "https://tophub.fun/#/?id=1", Name: "tophub.html"})
	select {}
}
