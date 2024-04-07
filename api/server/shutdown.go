package server

import (
	"log"
	"os"

	"RTF/storage"
)

// listens for shutdown signal and gracefully
// shuts down the server
func (d *DevServer) GracefulShutdown(signal chan os.Signal) {
	<-signal
	if err := storage.DB_Conn.Close(); err != nil {
		log.Fatal("ERROR CLOSING DATABASE", err)
	}
	os.Exit(0)
}
