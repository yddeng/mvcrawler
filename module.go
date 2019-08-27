package mvcrawler

import "github.com/tagDong/mvcrawler/util"

type Module interface {
	Search(context string) []*Message
	Update() []*Message
}

type ModuleType int32

const (
	Invaild  = ModuleType(0)
	Silisili = ModuleType(1)
	Bimibimi = ModuleType(2)
	End      = ModuleType(3)
)

var (
	mt2String = []string{
		"",
		"silisili",
		"bimibimi",
	}

	moduleFunc = map[ModuleType]func(anal *Analysis, down *Downloader, l *util.Logger) Module{}
)

//非安全的注册，需启动时完成注册
func Register(mt ModuleType, fn func(anal *Analysis, down *Downloader, l *util.Logger) Module) {
	if _, ok := moduleFunc[mt]; !ok {
		moduleFunc[mt] = fn
	}
}

func MT2String(mt ModuleType) string {
	if int(mt) > len(mt2String) {
		return ""
	}
	return mt2String[mt]
}
