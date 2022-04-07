package server

import "net/http"

func New(addr string) *http.Server {
	InitRoute()
	return &http.Server{
		Addr: addr,
	}

}
