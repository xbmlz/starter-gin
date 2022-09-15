package main

import (
	"github.com/xbmlz/starter-gin/core"
	"github.com/xbmlz/starter-gin/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func main() {
	// init config
	initialize.InitConfig()

	// init logger
	initialize.InitLogger()

	// init datasource
	initialize.InitDatasource()

	// run server
	core.RunServer()
}
