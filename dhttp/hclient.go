package dhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//发送GET请求
//url:请求地址; timeout:超时时间,小于等于0不设置超时
//response:请求返回的内容
func Get(url string, timeout time.Duration) (*http.Response, error) {
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	//Http的短连接，防止链接重用。解决 Connection reset by peer
	req.Close = true

	resp, rerr := client.Do(req)
	if rerr != nil {
		return nil, rerr
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http url %s get status code %d", url, resp.StatusCode)
	}
	return resp, nil
}

//发送Urlencoded POST请求
//url:请求地址; data:POST请求提交的数据; timeout:超时时间,小于等于0不设置超时;
//Response:请求返回的内容，err，请求错误
func PostUrlencoded(url string, data url.Values, timeout time.Duration) (*http.Response, error) {
	return Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()), timeout)
}

//发送Json POST请求
//url:请求地址; data:POST请求提交的json数据; timeout:超时时间,小于等于0不设置超时;
//Response:请求返回的内容，err，请求错误
func PostJson(url string, req interface{}, timeout time.Duration) (*http.Response, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	return Post(url, "application/json", bytes.NewReader(data), timeout)
}

//发送POST请求
//url:请求地址; data:POST请求提交的json数据; timeout:超时时间,小于等于0不设置超时;
//Response:请求返回的内容，err，请求错误
func Post(url string, contentType string, reader io.Reader, timeout time.Duration) (*http.Response, error) {
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	//Http的短连接，防止链接重用。解决 Connection reset by peer
	req.Close = true

	resp, rerr := client.Do(req)
	if rerr != nil {
		return nil, rerr
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http url %s post status code %d", url, resp.StatusCode)
	}

	return resp, nil
}
