package conf

import (
	"github.com/tagDong/mvcrawler/util"
)

type Config struct {
	Common struct {
		HttpAddr string
		Log      struct {
			LogPath string
			LogName string
		}
		DownLoad struct {
			OutPath        string
			ChanSize       int
			GoroutineCount int
		}
		Analysis struct {
			ChanSize       int
			GoroutineCount int
		}
	}

	Update struct {
		Route string
		Urls  []struct {
			Index string
			//数组表示深度
			Selectors []*Selector
		}
	}
}

//选择器
type Selector struct {
	Dom  string // DOM元素 选择器条件
	Exec []struct {
		//这一个选择器应该具体到哪一个标签
		Dom string
		//Attr获取指定属性,如果为空则获取Text
		Attr string
	}
}

var config *Config

func LoadConfig(path string) {
	err := util.DecodeJsonFile(path, &config)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return config
}

var siliWeek = []int{
	6, 0, 1, 2, 3, 4, 5,
}

func GetSilisili() *Selector {
	//n := util.GetWeekDay()
	//dom := fmt.Sprintf(".swiper-container xfswiper%d li", siliWeek[n])
	dom := ".time_con li"

	return &Selector{
		Dom: dom,
		Exec: []struct {
			Dom  string
			Attr string
		}{
			{
				Dom:  "p",
				Attr: "",
			},
			{
				Dom:  "img",
				Attr: "src",
			},
		},
	}
}
