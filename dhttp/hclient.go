package dhttp

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

var defTimeout = 5 * time.Second

//发送GET请求
//url:请求地址
//response:请求返回的内容
func Get(url string) (*http.Response, error) {
	client := &http.Client{Timeout: defTimeout}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

//发送POST请求
//url:请求地址，data:POST请求提交的数据
//Response:请求返回的内容，err，请求错误
func Post(url string, data url.Values) (*http.Response, error) {
	client := &http.Client{Timeout: defTimeout}
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode())) //("bar=baz&foo=quux")
	if err != nil {
		return nil, err
	}
	//contentType:请求体格式，如：application/json
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return client.Do(req)
}
