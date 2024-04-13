package main

import (
	"RTF/api/server"
	"RTF/utils"
)

func main() {
	utils.ErrorConsoleLog(server.NewDevServer(":7000").Boot().Error())
}
