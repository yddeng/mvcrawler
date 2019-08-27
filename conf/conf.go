package conf

import "github.com/BurntSushi/toml"

type Config struct {
	HttpAddr string
	TichDur  int
	Log      struct {
		LogPath string
		LogName string
	}
	DownLoad struct {
		OutPath        string
		QueueSize      int
		GoroutineCount int
	}
	Analysis struct {
		QueueSize      int
		GoroutineCount int
	}
}

var config *Config

func LoadConfig(path string) {
	config = &Config{}
	_, err := toml.DecodeFile(path, config)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return config
}
