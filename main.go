package main

import (
	"github.com/xbmlz/starter-gin/core"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	// setup
	core.Setup()
	// run server
	core.RunServer()
}
