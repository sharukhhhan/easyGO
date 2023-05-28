package main

import (
	"easyGO/server"
	"log"
)

func main() {
	srv := new(server.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
