package mvcrawler

import (
	"net/http"
	"sync"
	"sync/atomic"
	"unsafe"
)

type Update struct {
	values map[string][]*info
}

func newUpdate() {
	u := &Update{
		values: map[string][]*info{},
	}
	wg := sync.WaitGroup{}
	wg.Add(len(_update.values))

	wg.Wait()
	store(u)
}

type info struct {
	title  string
	status string
	url    string
}

var _update *Update

func infos(msgs [][]string) (ret []*info) {
	for _, v := range msgs {
		ret = append(ret, &info{
			title: v[0],
			url:   v[1],
		})
	}
	return
}

//更新替换，原子操作
func store(u *Update) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&_update)), unsafe.Pointer(u))
}

func update(writer http.ResponseWriter, request *http.Request) {
	up := (*Update)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&_update))))
	logger.Debugln(up.values)
}
