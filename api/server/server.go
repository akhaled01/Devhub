package server

import (
	"net"
	"net/http"
)

type DevServer struct {
	ListenAddr string
	Router     *http.ServeMux
}

func NewDevServer(listenAddr string) *DevServer {
	conn, _ := net.Dial("udp", "8.8.8.8:80")

	host := conn.LocalAddr().(*net.UDPAddr)

	return &DevServer{
		ListenAddr: host.IP.String() + listenAddr,
		Router:     http.NewServeMux(),
	}
}
