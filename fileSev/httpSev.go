package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"net/http"
	"os"
)

type Config struct {
	Addr     string
	LoadPath string
}

func getConfig(path string) *Config {
	cfg := &Config{}
	_, err := toml.DecodeFile(path, cfg)
	if err != nil {
		panic("./conf.toml not found")
	}
	return cfg
}

func main() {
	if len(os.Args) < 1 {
		fmt.Printf("usage config\n")
		return
	}

	conf := getConfig(os.Args[1])
	fmt.Printf("HttpServer On %s, LoadPath On %s\n", conf.Addr, conf.LoadPath)

	http.Handle("/", http.FileServer(http.Dir(conf.LoadPath)))
	err := http.ListenAndServe(conf.Addr, nil)
	if err != nil {
		panic(err)
	}
}
