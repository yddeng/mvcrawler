/*
 * cache
 */
package db

import (
	"github.com/tagDong/dutil/heap"
	"sync"
	"time"
)

type TT int16

const (
	LFU TT = iota //（Least Frequently Used）算法根据数据的历史访问频率来淘汰数据，其核心思想是“如果数据过去被访问多次，那么将来被访问的频率也更高”。
	LRU           //（Least recently used，最近最少使用）算法根据数据的历史访问记录来进行淘汰数据，其核心思想是“如果数据最近被访问过，那么将来被访问的几率也更高”。
)

// 字符串类型的缓存
type Cache struct {
	size    int //存储量
	mu      sync.Mutex
	data    map[string]*val
	minHeap *heap.Heap
}

type val struct {
	key          string
	value        interface{}
	lastUsedTime time.Time //最近使用时间，LRU算法
	usedTimes    int       //使用次数，LFU算法
}

var (
	cacheTT = LFU
)

func SetCacheTT(tt TT) { cacheTT = tt }
func GetCacheTT() TT   { return cacheTT }

func (v *val) Less(e heap.Element) bool {
	switch cacheTT {
	case LFU:
		return v.usedTimes < (e.(*val).usedTimes)
	case LRU:
		return v.lastUsedTime.Before(e.(*val).lastUsedTime)
	default:
		return true
	}
}

func New(size int, tt TT) *Cache {
	cacheTT = tt
	c := &Cache{size: size, data: map[string]*val{}, minHeap: heap.NewHeap()}
	return c
}

func (c *Cache) Set(key string, value interface{}) {
	var _val *val
	var ok bool
	c.mu.Lock()
	_val, ok = c.data[key]
	if ok {
		_val.value = value
	} else {
		c.checkAndRemove()
		_val = &val{
			key:          key,
			value:        value,
			lastUsedTime: time.Now(),
			usedTimes:    1,
		}
		c.data[key] = _val
		c.minHeap.Push(_val)
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_val, ok := c.data[key]
	if ok {
		_val.usedTimes++
		_val.lastUsedTime = time.Now()
		c.minHeap.Fix(_val)
		return _val.value, true
	} else {
		//todo 从数据库拉取
	}
	return "", false
}

func (c *Cache) GetAll() map[string]interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	ret := map[string]interface{}{}
	for _, v := range c.data {
		ret[v.key] = v.value
	}
	return ret
}

func (c *Cache) Size() int {
	return len(c.data)
}

func (c *Cache) checkAndRemove() {
	if len(c.data) >= c.size {
		min := c.minHeap.Pop()
		delete(c.data, min.(*val).key)

		//todo 刷新到数据库
	}
}
