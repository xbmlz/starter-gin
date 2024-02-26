package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/internal/conf"
	"github.com/xbmlz/starter-gin/internal/log"
	"github.com/xbmlz/starter-gin/internal/middleware"
)

type Server struct {
	srv *http.Server
}

func NewHTTPServer() *Server {
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	r.Use(middleware.Logger())

	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port),
		Handler: r,
	}

	return &Server{srv: srv}
}

func (s *Server) Run(ctx context.Context) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	errCh := make(chan error, 1)

	go func() {
		log.Sugar.Infof("server is running at http://%s:%d", conf.Server.Host, conf.Server.Port)
		if err := s.srv.ListenAndServe(); err != nil {
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := s.srv.Shutdown(ctx); err != nil {
			log.Sugar.Errorf("failed to shutdown server: %v", err)
		}
	}()

	// wait all server graceful shutdown
	wg.Wait()
	return nil
}
