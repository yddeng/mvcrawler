#### 网站模块

需实现对外的两个的接口

```
//name 
GetName()string
//url
GetUrl()string
//搜索接口
Search(context string) []*Item
//获取日更新接口
Update() [][]*Item
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

```

2）更新选择器

```
1.选取网站上更新的部分代码分析。
2.定位每一个选取内容的标签，制作选择器。(步骤同上）
3.由于是显示一周内每日的内容，故为一个双层选择器。

<div class="time_con" style="display:none">
            <div class="swiper-container xfswiper0">
              <div class="swiper-wrapper">
                <div class="swiper-slide">
                  <ul class="clear">
                    <li> <a href="/anime/2191.html" title="少女☆寸剧 All Starlight"><img src="https://wxt.sinaimg.cn/orj360/006bnWk0ly1g5a2hzzhx6j30hs0p0n2a.jpg" alt="少女☆寸剧 All Starlight" />
                      <p>少女☆寸剧 All Sta...</p>
                      <i>更新至07话</i><b>new</b>
                      <div class="fc_m"><strong></strong></div>
                      </a> </li>
                    <li> <a href="/anime/2168.html" title="暗芝居第七季"><img src="http://wxt.sinaimg.cn/large/777d58c0gy1g3vvyy9ldlj205006omx5.jpg" alt="暗芝居第七季" />
                      <p>暗芝居第七季</p>
                      <i>更新至08话</i><b>new</b>
                      <div class="fc_m"><strong></strong></div>
                      </a> </li>
                    <li> <a href="/anime/2147.html" title="猎兽神兵"><img src="https://wxt.sinaimg.cn/orj360/006bnWk0gy1g08hhe4qw2j30jg0rk40r.jpg" alt="猎兽神兵" />
                      <p>猎兽神兵</p>
                      <i>更新至08话</i>
                      <div class="fc_m"><strong></strong></div>
                      </a> </li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
```

3）注册到模块

```
mvcrawler.Register(mvcrawler.Silisili, func(l *util.Logger) mvcrawler.Module {
	return &Silisili{
		...
	}
})
```

