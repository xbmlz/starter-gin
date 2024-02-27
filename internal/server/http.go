package server

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xbmlz/starter-gin/docs"
	"github.com/xbmlz/starter-gin/internal/conf"
	"github.com/xbmlz/starter-gin/internal/handler"
	"github.com/xbmlz/starter-gin/internal/log"
	"github.com/xbmlz/starter-gin/internal/middleware"
	"github.com/xbmlz/starter-gin/ui"
)

type Server struct {
	srv *http.Server
}

func NewHTTPServer() *Server {
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	t, _ := template.ParseFS(ui.Templates, "templates/*.html")

	r.SetHTMLTemplate(t)

	addr := fmt.Sprintf("%s:%d", conf.Config.Server.Host, conf.Config.Server.Port)
	// swagger doc
	docs.SwaggerInfo.Host = addr
	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	r.Use(
		middleware.LogMiddleware(),
		middleware.CORSMiddleware(),
	)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Starter Gin",
			"content": "Welcome to Starter Gin!",
		})
	})

	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	v1 := r.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/register", handler.UserRegister)
			noAuthRouter.POST("/login", handler.UserLogin)
		}

		// Route group with permission
		authRouter := v1.Group("/").Use(middleware.AuthMiddleware())
		{
			authRouter.GET("/users", handler.GetUsers)
		}
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return &Server{srv: srv}
}

func (s *Server) Run(ctx context.Context) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	errCh := make(chan error, 1)

	go func() {
		log.Sugar.Infof("server is running at http://%s:%d", conf.Config.Server.Host, conf.Config.Server.Port)
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Sugar.Errorf("server is stopped: %v", err)
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		_ = s.Stop()
		return err
	case <-ctx.Done():
		return s.Stop()
	case <-quit:
		return s.Stop()
	}
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		if err := s.srv.Shutdown(ctx); err != nil {
			log.Sugar.Errorf("failed to shutdown server: %v", err)
		}
	}()

	return nil
}
