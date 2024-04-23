package main

import (
	"github.com/NextCodeX/desktop-application-pedestal/api"
	"github.com/changlie/go-common/a"
)

func main() {
	staticDir := "./"
	//staticDir := a.ProgramDir()

	server := a.HttpServerNew(12345)
	server.Post("/dap/:mod/:action", api.Entry)
	server.Static("/", staticDir)
	server.Start()
}
