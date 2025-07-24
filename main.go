package main

import (
	"github.com/boomthdev/wld-price-cheker/config"
	"github.com/boomthdev/wld-price-cheker/server"
)

func main() {
	conf := config.ConfigGetting()
	server := server.NewFiberServer(conf)
	server.Start()
}
