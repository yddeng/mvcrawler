package util

import (
	"bytes"
	"fmt"
)

var ts = []string{
	"b",
	"Kb",
	"Mb",
	"Gb",
}

// 字节长度格式化输出
// 例：2566b -> 2.50Kb
func SiezToString(b int64) string {
	n := float64(b)
	i := 0
	for n > 1024 {
		n /= 1024
		i++
	}

	return fmt.Sprintf("%.2f%s", n, ts[i])
}

// 拼接字符串
func MergeString(args ...string) string {
	buffer := bytes.Buffer{}
	for _, str := range args {
		buffer.WriteString(str)
	}
	return buffer.String()
}
