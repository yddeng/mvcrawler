### 动漫网爬虫

有时候想看番剧，由于更新时间的不一致、资源不通，需要在切换到各类网站上搜索。

尝试写一个爬取各类网站当日更新的网站；搜索在所有网站上搜索。

#### 收录网站

- www.silisili.me
- www.bimibimi.tv
- www.5dm.tv (由于服务器在境外，访问该网站有权限问题)

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
#### Modele

收录网站：添加网站方法

[网站模块注册](./module/README.md)

#### 更新日志

[更新日志](./UPDATE.md)

#### TODO

[TODO](./TODO.md)

#### 交流反馈

网站收录或建议 提交Issues:[Issues](https://github.com/tagDong/mvcrawler/issues)


