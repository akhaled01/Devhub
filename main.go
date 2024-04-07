package main

import (
	"log"

	serv "RTF/api/server"
)

func main() {
	server := serv.NewDevServer(":7000")
	log.Fatal(server.Boot())
}
