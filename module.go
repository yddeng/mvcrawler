package mvcrawler

type Module interface {
	Search(context string) []*Message
	Update()
}

type ModuleType int32

const (
	Invaild  = ModuleType(0)
	Silisili = ModuleType(1)
	End      = ModuleType(2)
)

var (
	mt2String = []string{
		"",
		"silisili",
	}

	moduleFunc = map[ModuleType]func() Module{}
)

//非安全的注册，需启动时完成注册
func Register(mt ModuleType, fn func() Module) {
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
