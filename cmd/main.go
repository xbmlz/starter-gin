package main

import (
	"context"
	"flag"

	"github.com/xbmlz/starter-gin/internal/conf"
	"github.com/xbmlz/starter-gin/internal/log"
	"github.com/xbmlz/starter-gin/internal/repo"
	"github.com/xbmlz/starter-gin/internal/server"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file, e.g. config/config.yaml")
	flag.Parse()

	log.InitLogger()

	err := conf.Load(*configPath)
	if err != nil {
		panic(err)
	}

	err = repo.InitDB()
	if err != nil {
		panic(err)
	}

	err = repo.MigrateDB()
	if err != nil {
		panic(err)
	}

	srv := server.NewHTTPServer()

	if err := srv.Run(context.Background()); err != nil {
		panic(err)
	}
}
