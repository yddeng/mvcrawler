/*
 *
 */
package main

import (
	"encoding/json"
	"fmt"
	"github.com/tagDong/dutil/dstring"
	"github.com/tagDong/dutil/io"
	"strings"
)

type Item struct {
	Info map[string]string `json:"info"`
}

func main() {
	fmt.Println("process start")

	var m = map[string]*Item{}

	bt, err := io.ReadFile("./item.txt")
	if err == nil {
		s := string(bt)
		lines := strings.Split(s, "\n")
		for _, line := range lines {
			v := strings.Split(line, "@")
			if len(v) == 3 {
				name := v[0]
				title := v[1]
				url := v[2]
				if dstring.IsEmpty(name) || dstring.IsEmpty(title) || dstring.IsEmpty(url) {
					fmt.Println("err line", line)
					continue
				}

				item, ok := m[name]
				if !ok {
					m[name] = &Item{Info: map[string]string{}}
					item = m[name]
				}
				item.Info[title] = url
			} else {
				fmt.Println("err len", line)
			}
		}

		jsonBt, er := json.Marshal(m)
		if er == nil {
			io.WriteByte("./", "item.json", jsonBt)
		}
	}

	fmt.Println("process end")
}
