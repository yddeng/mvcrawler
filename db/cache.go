/*
 * cache
 */
package db

import (
	"fmt"
	"github.com/tagDong/dutil/heap"
	"reflect"
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
	c       *Client
	size    int //存储量
	mu      sync.Mutex
	valType reflect.Type
	data    map[string]*val
	minHeap *heap.Heap
}

type val struct {
	key          string
	value        interface{}
	lastUsedTime time.Time //最近使用时间，LRU算法
	usedTimes    int       //使用次数，LFU算法
	dbDirty      bool      //脏数据，存储
}

var (
	cacheTT = LFU
)

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

func New(c *Client, size int, vT interface{}) *Cache {
	cache := &Cache{c: c, size: size, valType: reflect.TypeOf(vT), data: map[string]*val{}, minHeap: heap.NewHeap()}
	return cache
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
	_val.dbDirty = true
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
		data, err := c.c.get(key, c.valType)
		if err != nil {
			fmt.Println("dbget ", err)
			return "", false
		}

		c.checkAndRemove()
		_val = &val{
			key:          key,
			value:        data,
			lastUsedTime: time.Now(),
			usedTimes:    1,
		}
		c.data[key] = _val
		c.minHeap.Push(_val)
		_val.dbDirty = true

		return data, true
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
		c.c.save(min.(*val).key, min.(*val).value)
	}
}

func (c *Cache) PackDirty(fields map[string]interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, v := range c.data {
		if v.dbDirty {
			fields[v.key] = v.value
		}
	}
}

func (c *Cache) ClearDirty(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_val, ok := c.data[key]
	if ok {
		_val.dbDirty = false
	}
}
