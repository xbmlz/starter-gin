package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xbmlz/starter-gin/global"
	"github.com/xbmlz/starter-gin/initialize"
	"github.com/xbmlz/starter-gin/router"
)

func main() {
	// init config
	initialize.InitConfig()

	// init logger
	initialize.InitLogger()

	// create router
	router := router.CreateRouter()

	address := fmt.Sprintf("%s:%d", global.Config.Server.Address, global.Config.Server.Port)

	server := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
