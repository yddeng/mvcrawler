#### 网站模块

需实现对外的两个的接口

```
//搜索接口
Search(context string) []*Message
//获取日更新接口
Update() []*Message
```

实现过程

1) 定义搜索选择器

```
/*
 例 Search：
 1.在silisili搜索栏搜索"海贼王"。
 2.获取搜索的返回结果，保存或打印以便分析。
 3.根据内容定位我们需要title,img路径,url链接等
*/

 ps：截取的返回结果需要分析的部分
 <div class="anime_list">
    <dl>
        <dt><a href="/anime/1522.html"><img src="http://wxt.sinaimg.cn/small/00747wQSgy1fw2ubg9olzj305u08rjrp.jpg"/></a></dt>
        <dd>
            <h3><a href="/anime/1522.html">海贼王</a></h3>
            <p><div class="d_label"><b>地区：</b>日本</div><div class="d_label"><b>年代：</b>1997年</div></p>
            <p><div class="d_label"><b>标签：</b>奇幻,冒险,动作</div><div class="d_label"><b>播放：</b>1288974</div></p>
            <p><b>看点：</b>OP在哪呢</p>
            <p><b>简介：</b>有个男人他拥有世界上一切财富、名望和权势，他就是「海盗王」高路德•罗杰。在临死前说过这样一句话：让全世界的人都奔向大海「想要我的财宝吗？想要的话全就拿去吧……！你们去找吧！我把世界上的一切都</p>
            <p><b style="color:#F00">状态:</b> 连载 </p>
        </dd>
    </dl>
    <dl>
        <dt><a href="/anime/343.html"><img src="http://wxt.sinaimg.cn/small/00747wQSgy1fw2j026g9bj307g09yjsp.jpg"/></a></dt>
        <dd>
            <h3><a href="/anime/343.html">海贼王剧场版</a></h3>                    
            <p><div class="d_label"><b>地区：</b>日本</div><div class="d_label"><b>年代：</b>1997年</div></p>
            <p><div class="d_label"><b>标签：</b>热血,冒险,奇幻</div><div class="d_label"><b>播放：</b>88724</div></p>
            <p><b>看点：</b>热血少年追寻梦想的故事！！！</p>
            <p><b>简介：</b>《ONE PIECE》（海贼王、航海王）简称“OP”，是日本漫画家尾田荣一郎作画的少年漫画作品。总有一天，我会聚集一群不输给这些人的伙伴，并找到世界第一的财宝，我要当海贼王！！！</p>
            <p><b style="color:#F00">状态:</b> 完结 </p>
        </dd>
     </dl>			
 </div>			

/* 
 确定item选择器
 Dom: ".anime_list dl",
 确定具体选择器
 {Dom: "dd h3 a", Attr: ""},     //title
 {Dom: "dt img", Attr: "src"},   //img src
 {Dom: "dd h3 a", Attr: "href"}, //url
*/
```

2）更新选择器

```
1.选取网站上每日更新的部分代码分析。
2.定位每一个选取内容的标签，制作选择器。(步骤同上）
3.日更新会有时间的变化，要选择具体的标签，需与时间计算。

如：item选择器的定义
dom := fmt.Sprintf(".xfswiper%d li", siliWeek[n])
```

3）注册到模块

```
mvcrawler.Register(mvcrawler.Silisili, func(
	anal *mvcrawler.Analysis, down *mvcrawler.Downloader, l *util.Logger) mvcrawler.Module {

	return &Silisili{
		...
	}
})
```


#### 收录的网站

- www.silisili.me
- www.dimidimi.tv