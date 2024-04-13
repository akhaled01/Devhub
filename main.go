package main

import (
	"log"

	serv "RTF/api/server"
	"RTF/utils"
)

func main() {
	defer utils.RecoverFromPanic()
	server := serv.NewDevServer(":7000")
	log.Fatal(server.Boot())
}
