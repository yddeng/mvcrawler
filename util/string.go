package util

import (
	"github.com/yddeng/dutil/dstring"
)

// 判读s串中是否有str子串
// 有返回，没有在头部添加再返回
func CheckAndInsertHead(s, substr, headStr string) string {
	if dstring.CheckString(s, substr) {
		return s
	}
	return dstring.MergeString(headStr, s)
}
