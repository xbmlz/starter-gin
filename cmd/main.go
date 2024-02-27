package main

import (
	"context"
	"flag"

	"github.com/xbmlz/starter-gin/internal/conf"
	"github.com/xbmlz/starter-gin/internal/log"
	"github.com/xbmlz/starter-gin/internal/model"
	"github.com/xbmlz/starter-gin/internal/server"
)

// @title           Example API
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io
// @license.name    MIT
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host            localhost:8080
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	configPath := flag.String("config", "config.yaml", "path to config file, e.g. config/config.yaml")
	flag.Parse()

	log.InitLogger()

	err := conf.Load(*configPath)
	if err != nil {
		panic(err)
	}

	err = model.InitDB()
	if err != nil {
		panic(err)
	}

	srv := server.NewHTTPServer()

	if err := srv.Run(context.Background()); err != nil {
		panic(err)
	}
}
