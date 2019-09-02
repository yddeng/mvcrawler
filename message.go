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
