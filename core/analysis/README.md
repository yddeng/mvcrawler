#### analysis

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