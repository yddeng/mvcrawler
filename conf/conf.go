package conf

import "github.com/BurntSushi/toml"

type Config struct {
	HttpAddr  string
	UpdateDur int
	SearchDur int
	Log       struct {
		LogPath string
		LogName string
	}
	Web struct {
		PageItemCount int
	}
	DB struct {
		CacheSize int
		DBSaveDur int
		SavePath  string
		MySQLAddr string
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
