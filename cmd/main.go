package main

import (
	"easyGO/internal/transport/rest"
	"easyGO/server"
	"log"
)

func main() {
	handler := new(rest.Handler)
	srv := new(server.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
