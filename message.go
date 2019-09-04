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
	Code    int     `json:"code"`
	PageNum int     `json:"page_num"`
	Items   []*Item `json:"items"`
}

//更新返回结构
type UpdateRespone struct {
	Code  int       `json:"code"`
	Items [][]*Item `json:"items"`
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
