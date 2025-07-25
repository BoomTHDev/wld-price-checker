package handler

import (
	"net/http"

	"github.com/boomthdev/wld-price-cheker/config"
	"github.com/boomthdev/wld-price-cheker/server"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	conf := config.ConfigGetting()
	server := server.NewFiberServer(conf)
	server.Start()
}
