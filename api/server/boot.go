package server

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/template"

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

	// Create a file server to serve static files (CSS, JS, images, etc.)
	fs := http.FileServer(http.Dir("../frontend"))

	// Handle requests for files in the "/static/" path
	http.Handle("/", http.StripPrefix("/", fs))

	mainPage(d.Router)
	posts.RegisterPostRoutes(d.Router)
	auth.RegisterAuthRoutes(d.Router)
	chat.RegisterChatRoutes(d.Router)

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGTERM)
	go d.GracefulShutdown(shutdownSignal)

	utils.InfoConsoleLog(red + "DevHub API is Live on http://127.0.0.1" + d.ListenAddr + reset)
	return http.ListenAndServe(d.ListenAddr, d.Router)
}

func mainPage(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./frontend/index.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	})
}
