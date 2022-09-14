package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xbmlz/starter-gin/core/config"
	"github.com/xbmlz/starter-gin/core/logger"
	"github.com/xbmlz/starter-gin/routers"
)

func main() {
	// setup config
	config.Setup()

	// setup logger
	logger.Setup()

	// init router
	router := routers.InitRouter()

	address := fmt.Sprintf("%s:%d", config.App.Server.Address, config.App.Server.Port)

	server := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println(config.App.Server.Port)

	server.ListenAndServe()
}
