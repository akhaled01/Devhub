package server

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"RTF/api/auth"
	"RTF/api/posts"
	"RTF/storage"
)

const (
	brightGreen = "\033[32;1m"
	reset       = "\033[0m"
)

// starts the server after registering all the endpoints
// and kickstarting the DB
func (d *DevServer) Boot() error {
	if err := storage.Init(); err != nil {
		return err
	}

	posts.RegisterPostRoutes(d.Router)
	auth.RegisterAuthRoutes(d.Router)

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGTERM)
	go d.GracefulShutdown(shutdownSignal)

	fmt.Println(brightGreen + "DevHub API is Live on http://127.0.0.1:7000" + reset)
	return http.ListenAndServe(":7000", d.Router)
}
