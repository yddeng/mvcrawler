package conf_test

import (
	"fmt"
	"github.com/tagDong/mvcrawler"
	"github.com/tagDong/mvcrawler/conf"
	"testing"
)

func TestNewConfig(t *testing.T) {
	conf.LoadConfig("conf.json")
	c := conf.GetConfig()
	fmt.Println(c)

	mvcrawler.InitLogger()
}
