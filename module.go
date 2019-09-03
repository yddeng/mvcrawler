package mvcrawler

import (
	"github.com/tagDong/dutil/log"
)

type Module interface {
	GetName() string
	// url
	GetUrl() string
	// 周一至周日的集合
	Update() [][]*Message
	// 结果集合
	Search(txt string) []*Message
}

type ModuleType int32

const (
	Invaild  = ModuleType(0)
	Silisili = ModuleType(1)
	Bimibimi = ModuleType(2)
	Dm5      = ModuleType(3)
	End      = ModuleType(4)
)

var (
	moduleFunc = map[ModuleType]func(l *log.Logger) Module{}
)

//非安全的注册，需启动时完成注册
func Register(mt ModuleType, fn func(l *log.Logger) Module) {
	if _, ok := moduleFunc[mt]; !ok {
		moduleFunc[mt] = fn
	}
}
