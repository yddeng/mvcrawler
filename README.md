### 电影爬虫

有时候想看番剧，由于更新时间的不一致、资源不通，需要在切换到各类网站上搜索。

尝试写一个爬取各类网站当日更新的网站；搜索在所有网站上搜索。

#### 收录网站

- www.silisili.me
- www.bimibimi.tv

#### 安装

```
1. go get github.com/tagDong/mvcrawler
2. 配置 conf/conf.toml
3. make 
4. ./start.sh  
```

#### 目录结构
```
mvcrawler/
├── app                       前端展示页   
│   └── index.html     
├── conf                      配置文件
│   |── conf.go       
│   └── conf.toml.template   
├── fileSev                   文件服务器   
│   ├── conf.toml.template 
│   └── httpSev.go      
├── dhttp                     http简单封装
│   ├── hclient.go    
│   └── hserver.go  
├── main                      程序入口
│   └── crawler.go 
├── module                    爬取模块
│   ├── bimibimi.go
│   └── silisili.go
├── util                      工具包
│  
└── README.md
```

#### Analysis分析器

```
// 定义一个选择器
type Selector struct {
	Dom  string // DOM元素 选择器条件
	Exec []struct {
		//这一个Dom应该具体到哪一个标签
		Dom string
		//Attr获取指定属性,如果为空则获取Text
		Attr string
	}
}

// 非阻塞异步投递，队列满丢弃
func (a *Analysis) Post(req *AnalysisReq, callback func(resp *AnalysisReap)) error
// 同步投递
func (a *Analysis) SyncPost(req *AnalysisReq) (resp *AnalysisReap, err error)

// 执行过程
1.实例一个 AnalysisReq
2.同步或异步投递
3.分析器处理线程 exec()
  1）判断HttpResp为空: ture执行2), false执行3)。
  2) 请求url并填入HttpResp。
  3）根据Selector获取返回。
```

#### Modele

收录网站：添加网站方法

[网站模块注册](./module/README.md)

#### 更新

[更新日志](./UPDATE.md)

#### 网页部署

一个简单的目录服务器。配置启动后，加载html文件，可供外网访问。
[网站部署](./fileSev/README.md)