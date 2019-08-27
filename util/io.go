package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func ReadFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

//读取json文件并反序列化
func DecodeJsonFile(filePath string, i interface{}) error {
	data, err := ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, i)
}

func WriteFile(filePath, fileName string, reader io.Reader) (n int64, err error) {
	os.MkdirAll(filePath, os.ModePerm)

	f, err := os.Create(path.Join(filePath, fileName))
	defer f.Close()
	n, err = io.Copy(f, reader)
	return
}
