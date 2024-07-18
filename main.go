package main

import (
	"github.com/joho/godotenv"
	"github.com/xbmlz/starter-gin/internal/db"
	"github.com/xbmlz/starter-gin/internal/server"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db.MustInit()

	server.Run()
}
