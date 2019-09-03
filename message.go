package mvcrawler

type Message struct {
	Title  string `json:"title"`
	From   string `json:"from"`
	Img    string `json:"img"`
	Status string `json:"status"`
	Url    string `json:"url"`
}

type SearchRespone struct {
	Code     int        `json:"code"`
	MsgNum   int        `json:"msg_num"`
	Messages []*Message `json:"messages"`
}

type UpdateRespone struct {
	Code     int          `json:"code"`
	Messages [][]*Message `json:"messages"`
}

//去重
func Process(data *[]*Message) {
	tmp := []*Message{}
	m := map[string]struct{}{}

	for _, v := range *data {
		if _, ok := m[v.Title]; !ok {
			tmp = append(tmp, v)
			m[v.Title] = struct{}{}
		}
	}

	*data = tmp
}
