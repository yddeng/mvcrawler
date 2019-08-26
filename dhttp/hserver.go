package dhttp

import "net/http"

type HttpServer struct {
	handler    *http.ServeMux
	listenAddr string
}

func NewHttpServer(addr string) *HttpServer {
	s := new(HttpServer)
	s.handler = http.NewServeMux()
	s.listenAddr = addr

	return s
}

func (s *HttpServer) Register(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.handler.HandleFunc(pattern, handler)
}

func (s *HttpServer) Listen() error {
	return http.ListenAndServe(s.listenAddr, s.handler)
}
