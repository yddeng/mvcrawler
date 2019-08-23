package mvcrawler

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestHttp(t *testing.T) {
	postV := url.Values{
		"keyboard": {"海"},
	}
	resp, err := http.PostForm("http://www.silisili.me/e/search/index.php", postV)
	fmt.Println(resp, err)
}
