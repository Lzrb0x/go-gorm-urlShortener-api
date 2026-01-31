package main

import (
	"net/http"

	"github.com/Lzrb0x/go-gorm-urlShortener-api/config"
	"github.com/Lzrb0x/go-gorm-urlShortener-api/db"
	"github.com/Lzrb0x/go-gorm-urlShortener-api/route"
)

func main() {
	database := db.InitDB()

	handler := route.InitRoutes(database)

	server := &http.Server{
		Addr:    config.Config.AppPort,
		Handler: handler,
	}

	server.ListenAndServe()
}
