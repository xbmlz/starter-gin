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

	db.MustConnect()

	server.RunHTTPServer()
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
