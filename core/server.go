package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/api"
	"github.com/xbmlz/starter-gin/global"
)

// 启动服务
func RunServer() {
	// 注册路由
	engine := api.RegisterRouter()
	// 创建HTTP服务
	httpServer := createHttpServer(engine)
	// 打印服务信息
	printServerInfo()
	// 启动服务
	_ = httpServer.ListenAndServe()
}

// 创建http服务
func createHttpServer(engine *gin.Engine) *http.Server {
	return &http.Server{
		// ip和端口
		Addr: fmt.Sprintf("%s:%d", global.Config.Server.Address, global.Config.Server.Port),
		// 调用的处理器，如为nil会调用http.DefaultServeMux
		Handler: engine,
		// 计算从成功建立连接到request body(或header)完全被读取的时间
		ReadTimeout: time.Second * 10,
		// 计算从request body(或header)读取结束到 response write结束的时间
		WriteTimeout: time.Second * 10,
		// 请求头的最大长度，如为0则用DefaultMaxHeaderBytes
		MaxHeaderBytes: 1 << 20,
	}
}

// 打印服务信息
func printServerInfo() {
	serverConfig := global.Config.Server
	global.Log.Sugar().Infof("Server started on https://%s:%d", serverConfig.Address, serverConfig.Port)
}
