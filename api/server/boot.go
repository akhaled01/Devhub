package server

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"RTF/api/auth"
	"RTF/api/chat"
	"RTF/api/posts"
	"RTF/storage"
	"RTF/utils"
)

const (
	red   = "\033[31;1m"
	reset = "\033[0m"
)

// starts the server after registering all the endpoints
// and kickstarting the DB
func (d *DevServer) Boot() error {
	if err := storage.Init(); err != nil {
		return err
	}

	posts.RegisterPostRoutes(d.Router)
	auth.RegisterAuthRoutes(d.Router)
	chat.RegisterChatRoutes(d.Router)

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGTERM)
	go d.GracefulShutdown(shutdownSignal)

	conn, _ := net.Dial("udp", "8.8.8.8:80")
	host := conn.LocalAddr().(*net.UDPAddr)

	utils.InfoConsoleLog(red + "DevHub API is Live on http://" + host.IP.String() + d.ListenAddr + reset)
	return http.ListenAndServe(d.ListenAddr, d.Router)
}
