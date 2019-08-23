package dhttp

import (
	"net/http"
	"time"
)

var defTimeout = 5 * time.Second

func Get(url string) (*http.Response, error) {
	return GetTimeout(url, defTimeout)
}

func GetTimeout(url string, timeout time.Duration) (*http.Response, error) {
	client := &http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)

	return client.Do(request)
}
