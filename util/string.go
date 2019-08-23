package util

import "fmt"

var ts = []string{
	"b",
	"Kb",
	"Mb",
	"Gb",
}

func SiezToString(b int64) string {
	n := float64(b)
	i := 0
	for n > 1024 {
		n /= 1024
		i++
	}

	return fmt.Sprintf("%.2f%s", n, ts[i])
}
