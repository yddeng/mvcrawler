package mvcrawler

import (
	"encoding/json"
	"net/http"
)

type SearchReq struct {
	Txt     string `json:"txt"`
	Modules []int  `json:"modules"`
}

func (s *Service) search(w http.ResponseWriter, r *http.Request) {
	logger.Infoln("http search request")

	var req SearchReq
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Errorf("search read err: %s", err)
		return
	}

	for _, m := range s.modules {
		ret := m.Search(req.Txt)
		if err := json.NewEncoder(w).Encode(ret); err != nil {
			logger.Errorf("search write err: %s", err)
		}
	}
}
