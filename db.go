/*
 * db 存储结构
 */
package mvcrawler

import (
	"encoding/json"
	"fmt"
	"github.com/tagDong/dutil/cache"
	"github.com/tagDong/mvcrawler/conf"
	"github.com/tagDong/mvcrawler/util"
	"time"
)

// 搜索存储结构
type SearchDB struct {
	Name      string    //搜索字
	TotalItem int       //结果数量
	TotalPage int       //页码数量
	PageItems [][]*Item //分页后的项目集合
}

// 更新存储结构
type UpdateDB struct {
	Items [][]*Item //每日的项目集合
}

type dbClient struct {
	name  string
	cache *cache.Cache
	dbs   *dbs
}

var clientMap = map[string]*dbClient{}

// isSave 落地到文件， vT 结构类型
func NewClient(name string, isSave bool, vT interface{}) {
	_, ok := clientMap[name]
	if !ok {
		dbConf := conf.GetConfig().DB
		c := &dbClient{name: name}

		if isSave {
			c.dbs = &dbs{fPath: fmt.Sprintf("%s/%s", dbConf.SavePath, c.name)}
			c.cache = cache.New(c.dbs, dbConf.CacheSize, vT, cache.LFU)
			go c.loop(time.Duration(dbConf.DBSaveDur) * time.Second)
		} else {
			c.cache = cache.New(nil, dbConf.CacheSize, vT, cache.LFU)
		}

		clientMap[name] = c
	}
}

func GetClient(name string) *cache.Cache {
	c, ok := clientMap[name]
	if ok {
		return c.cache
	}
	return nil
}

func (c *dbClient) loop(dur time.Duration) {
	tick := time.NewTicker(dur)
	for {
		<-tick.C
		c.cache.SaveDirty()
	}
}

type dbs struct {
	fPath string
}

func (d *dbs) Save(k string, v interface{}) error {
	bt, err := json.Marshal(v)
	if err != nil {
		logger.Errorln(err)
		return err
	}
	err = util.WriteByte(d.fPath, fmt.Sprintf("%s.json", k), bt)
	logger.Debugln("dbsave ", k, err)
	return err
}

func (d *dbs) Load(k string, v interface{}) error {
	bt, err := util.ReadFile(fmt.Sprintf("%s/%s.json", d.fPath, k))
	if err != nil {
		logger.Errorln(err)
		return err
	}
	err = json.Unmarshal(bt, v)
	logger.Debugln("dbget ", k, err)
	return err
}
