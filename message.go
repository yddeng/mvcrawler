package mvcrawler

type Message struct {
	Title  string `json:"title"`
	From   string `json:"from"`
	Img    string `json:"img"`
	Status string `json:"status"`
	Url    string `json:"url"`
}
