package conf

import "github.com/tagDong/mvcrawler/util"

type Config struct {
	Common struct {
		Log struct {
			LogPath string
			LogName string
		}
		DownLoad struct {
			OutPath        string
			ChanSize       int
			GoroutineCount int
		}
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
