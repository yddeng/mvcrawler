/*
 *
 */
package db

import (
	"encoding/json"
	"fmt"
	"github.com/tagDong/mvcrawler/conf"
	"github.com/tagDong/mvcrawler/util"
	"reflect"
	"time"
)

type Client struct {
	name   string
	cache  *Cache
	isSave bool
	fPath  string
}

var clientMap = map[string]*Client{}

func NewClient(name string, isSave bool, vT interface{}) {
	_, ok := clientMap[name]
	if !ok {
		dbConf := conf.GetConfig().DB
		c := &Client{
			name:   name,
			isSave: isSave,
		}
		c.cache = New(c, dbConf.CacheCount, vT)

		if isSave {
			c.fPath = fmt.Sprintf("%s/%s", dbConf.SavePath, c.name)
			go c.loop(time.Duration(dbConf.CacheSaveDur) * time.Second)
		}
		clientMap[name] = c
	}
}

func GetClient(name string) *Cache {
	c, ok := clientMap[name]
	if ok {
		return c.cache
	}
	return nil
}

func (c *Client) loop(dur time.Duration) {
	tick := time.NewTicker(dur)
	for {
		<-tick.C
		fields := map[string]interface{}{}

		c.cache.PackDirty(fields)
		for k, v := range fields {
			err := c.save(k, v)
			if err == nil {
				c.cache.ClearDirty(k)
			}
		}

	}
}

func (c *Client) save(k string, v interface{}) error {
	bt, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = util.WriteByte(c.fPath, fmt.Sprintf("%s.json", k), bt)
	fmt.Println("dbsave ", k, err)
	return err
}

func (c *Client) get(key string, vT reflect.Type) (interface{}, error) {
	bt, err := util.ReadFile(fmt.Sprintf("%s/%s.json", c.fPath, key))
	if err != nil {
		return nil, err
	}

	ret := reflect.New(vT.Elem()).Interface()
	err = json.Unmarshal(bt, ret)
	fmt.Println("dnGet ", key, err)

	return ret, err
}
