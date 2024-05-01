package server

import (
	"net/http"
)

type DevServer struct {
	ListenAddr string
	Router     *http.ServeMux
}

func NewDevServer(listenAddr string) *DevServer {

	return &DevServer{
		ListenAddr: listenAddr,
		Router:     http.NewServeMux(),
	}
}
