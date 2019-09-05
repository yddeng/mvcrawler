package mvcrawler

// 项目结构
type Item struct {
	Title  string `json:"title"`
	From   string `json:"from"`
	Img    string `json:"img"`
	Status string `json:"status"`
	Url    string `json:"url"`
}

// 搜索返回结构
type SearchRespone struct {
	Code      string  `json:"code"`
	Txt       string  `json:"txt"`        //搜索词
	Page      int     `json:"page"`       //请求的页码
	TotalPage int     `json:"total_page"` //总页妈
	TotalItem int     `json:"total_item"` //总页妈
	Items     []*Item `json:"items"`
}

//更新返回结构
type UpdateRespone struct {
	Code  string    `json:"code"`
	Items [][]*Item `json:"items"`
}

var code = map[string]string{
	"ERR": "请求参数错误",
	"OK":  "请求处理成功",
}

//去重
func Process(data *[]*Item) {
	tmp := []*Item{}
	m := map[string]struct{}{}

	for _, v := range *data {
		if _, ok := m[v.Title]; !ok {
			tmp = append(tmp, v)
			m[v.Title] = struct{}{}
		}
	}

	*data = tmp
}
