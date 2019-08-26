package dhttp

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var defTimeout = 5 * time.Second

//发送GET请求
//url:请求地址; timeout:超时时间,小于等于0不设置超时
//response:请求返回的内容
func Get(url string, timeout time.Duration) (*http.Response, error) {
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, rerr := client.Do(req)
	if resp.StatusCode != http.StatusOK || rerr != nil {
		return nil, fmt.Errorf("http url %s get StatusCode %d err %s", url, resp.StatusCode, rerr)
	}

	return resp, nil
}

//发送POST请求
//url:请求地址; timeout:超时时间,小于等于0不设置超时; data:POST请求提交的数据
//Response:请求返回的内容，err，请求错误
func Post(url string, timeout time.Duration, data url.Values) (*http.Response, error) {
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode())) //("bar=baz&foo=quux")
	if err != nil {
		return nil, err
	}
	//contentType:请求体格式，如：application/json
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, rerr := client.Do(req)
	if resp.StatusCode != http.StatusOK || rerr != nil {
		return nil, fmt.Errorf("http url %s post StatusCode %d err %s", url, resp.StatusCode, rerr)
	}

	return resp, nil
}
