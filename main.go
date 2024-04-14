package main

import (
	"RTF/api/server"
	"RTF/utils"
)

func main() {
	utils.ErrorConsoleLog(server.NewDevServer(":8080").Boot().Error())
}
