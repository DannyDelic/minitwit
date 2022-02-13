package main

import (
	"minitwit/internal/cli"
	"minitwit/internal/conf"
	"minitwit/internal/server"
)

func main() {
	cli.Parse()
	server.Start(conf.NewConfig())
}
