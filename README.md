### 动漫网爬虫

有时候想看番剧，由于更新时间的不一致、资源不通，需要在切换到各类网站上搜索。

尝试写一个爬取各类网站当日更新的网站；搜索在所有网站上搜索。

#### 收录网站

- www.silisili.me
- www.bimibimi.tv
- www.5dm.tv (由于服务器在境外，访问该网站有权限问题)

#### 前端展示

![index](https://github.com/tagDong/mvcrawler/blob/master/assets/image/index.jpg)

#### 安装

服务器部署
```
1. go get github.com/tagDong/mvcrawler
2. 配置 conf/conf.toml
3. make 
4. ./start.sh  
```

网页部署

[网站部署](./fileSev/README.md)

#### 目录结构
```
mvcrawler/
├── app      前端展示页   
│   └── index.html     
├── conf     配置文件
│   |── conf.go       
│   └── conf.toml.template   
├── core     组件
│   ├── analysis    分析器
│   └── downloader  下载器
├── fileSev   文件服务器   
│   ├── conf.toml.template 
│   └── httpSev.go      
├── dhttp     http简单封装
│   ├── hclient.go    
│   └── hserver.go  
├── main      程序入口
│   └── crawler.go 
├── module    爬取模块
│   ├── bimibimi.go
│   └── silisili.go
├── util      工具包
│  
└── README.md
```

#### 低配版冷热数据库

```
set数据时
1.将数据保存在缓存中。
2.以一定时间间隔将缓存中的脏数据保存到本地json文件。
3.缓存容量满，根据淘汰算法删除缓存数据并保存到本地。

get数据时
1.在缓存中搜索(没有到2步）
2.查询并读取本地文件(有数据执行set）
3.返回结果

```

#### Modele

收录网站：添加网站方法

[网站模块注册](./module/README.md)

#### 更新日志

[更新日志](./UPDATE.md)

#### TODO

[TODO](./TODO.md)

#### 交流反馈

网站收录或建议 提交Issues:[Issues](https://github.com/tagDong/mvcrawler/issues)


