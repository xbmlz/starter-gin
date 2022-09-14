package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xbmlz/starter-gin/routers"
)

func main() {
	router := routers.InitRouter()
	// TODO read for config.yaml
	address := fmt.Sprintf("%s:%d", "0.0.0.0", 8000)

	server := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
