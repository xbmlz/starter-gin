package main

import (
	"context"
	"flag"

	"github.com/xbmlz/starter-gin/internal/config"
	"github.com/xbmlz/starter-gin/internal/server"
)

func main() {
	configPath := flag.String("config", "config/config.yaml", "path to config file, e.g. config/config.yaml")
	flag.Parse()

	err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	srv := server.NewHTTPServer()

	if err := srv.Run(context.Background()); err != nil {
		panic(err)
	}
}
