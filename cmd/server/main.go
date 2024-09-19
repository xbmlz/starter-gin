package main

import (
	"flag"

	"github.com/xbmlz/starter-gin/internal/config"
	"github.com/xbmlz/starter-gin/internal/data"
	"github.com/xbmlz/starter-gin/internal/logger"
	"github.com/xbmlz/starter-gin/internal/server"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the application.
	Version string

	// configFile is the path to the configuration file.
	configFile string
)

func init() {
	flag.StringVar(&configFile, "c", "config.yaml", "path to the configuration file, e.g. -c config.yaml")
}

func main() {

	// load config
	c, err := config.Load(configFile)
	if err != nil {
		panic(err)
	}

	// init logger
	logger := logger.NewZapLogger()
	defer logger.Sync()

	// init data
	data, cleanup, err := data.NewData(&c.Data, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}

	// start http server
	srv := server.NewHTTPServer(&c.Server, logger, data)

	srv.ListenAndServe()
}
