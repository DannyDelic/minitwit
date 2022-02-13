package main

import (
	"minitwit/internal/conf"
	"minitwit/internal/server"
)

func main() {
	server.Start(conf.NewConfig())
}
