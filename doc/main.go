package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//程序入口
func main() {
	var baseUrl = "http://www.demoweb.com"
	var firstPage = "/html/article/133461.html"
	doProcess(baseUrl, baseUrl+firstPage)
}

//递归处理，爬取下一页的内容
func doProcess(baseUrl string, url string) {
	cc := make(chan string)
	next_url := parseContent(cc, url)
	doProcess(baseUrl, baseUrl+next_url)
}

//解析网页内容
func parseContent(cc chan string, url string) string {
	var innerImg = []string{}
	doc, _ := goquery.NewDocument(url)
	div_body := doc.Find("div.n_bd")
	fmt.Println("div n_bd:", div_body)
	div_body.Find("img").Each(func(i int, s *goquery.Selection) {
		img_path, _ := s.Attr("src")
		innerImg = append(innerImg, img_path)
	})
	fmt.Println("img src:", innerImg)
	go downPic(doc, innerImg, cc)
	for range innerImg {
		fmt.Println("result:", <-cc)
	}

	next := doc.Find("li.next")
	next_src := next.Find("a")
	next_url, _ := next_src.Attr("href")
	fmt.Println("Next URL:", next_url)
	return next_url
}

//下载功能
func downPic(doc *goquery.Document, paths []string, cc chan string) {

	var base_dir = "/home/wang/pics/"

	for _, path := range paths {
		fmt.Println("single path:", path)
		res, _ := http.Get(path)
		file_name, dir := splitPath(path)
		os.Mkdir(base_dir+dir, 0777)

		fmt.Println("save dir:", base_dir+dir)
		save_dir := base_dir + dir + "/"
		fmt.Println("file:", save_dir+file_name)

		file, _ := os.Create(save_dir + file_name)
		io.Copy(file, res.Body)
	}
	cc <- "done"
}

//路径处理
func splitPath(path string) (string, string) {
	ss := strings.Split(path, "/")
	lenth := len(ss)
	file_name := ss[lenth-1]
	dir_name := ss[lenth-3] + ss[lenth-2]
	return file_name, dir_name
}

//goquery测试
func goqueryDemo() {
	doc, err := goquery.NewDocument("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(doc)
	head := doc.Find("head")
	meta := head.Find("meta")
	content, _ := meta.Attr("content")
	fmt.Println("content is:", content)
}

//原生代码
func httpGet() {
	resp, _ := http.Get("http://www.baidu.com")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
