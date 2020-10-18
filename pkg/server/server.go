package server

import (
	"context"
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"accounting/pkg/config"
	"accounting/pkg/log"
	"accounting/pkg/routes"
	"accounting/pkg/server/middleware"
	"syscall"
	"time"
)

type Server struct {
	log        log.Logger
	config     *config.Configuration
	routes     *routes.HttpRoutes
	onShutdown func()
}

func NewServer(config *config.Configuration, log log.Logger, routes *routes.HttpRoutes) *Server {
	return &Server{
		log:    log,
		config: config,
		routes: routes,
	}
}

func (s *Server) SetShutdown(onShutdown func()) {
	s.onShutdown = onShutdown
}

func (s *Server) Start() {

	httpServer := &http.Server{
		Addr:      s.config.Server.Port,
		Handler:   s.newGinHandler(s.config),
		TLSConfig: s.newTLSConfig(),
	}

	go s.listen(httpServer)

	s.listenShutdownSignal(httpServer)
}

func (s *Server) newTLSConfig() *tls.Config {
	if s.config.Server.CertFile != nil && s.config.Server.KeyFile != nil {

		pair, err := tls.LoadX509KeyPair(*s.config.Server.CertFile, *s.config.Server.KeyFile)
		if err != nil {
			s.log.Errorf("Failed to load key pair: %v", err)
			panic(err)
		}

		return &tls.Config{Certificates: []tls.Certificate{pair}}
	}

	return nil
}

func (s *Server) newGinHandler(config *config.Configuration) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()

	e.Use(middleware.NewLogRequest(s.log, []string{routes.HealthCheckPath}))

	e.Use(middleware.NewErrorHandler(s.log))

	s.routes.AddHttpRoutes(e)

	return e
}

func (s *Server) listen(httpServer *http.Server) {
	if httpServer.TLSConfig != nil {
		s.log.Infof("Starting Server (HTTP + TLS) on %s", httpServer.Addr)
		if err := httpServer.ListenAndServeTLS("", ""); err != nil {
			s.log.Warnf("Server (HTTP + TLS) stopped [%s]", err)
		}
	} else {
		s.log.Infof("Starting Server (HTTP) on %s", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil {
			s.log.Warnf("Server (HTTP) stopped [%s]", err)
		}
	}
}

func (s *Server) listenShutdownSignal(httpServer *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.log.Info("Shutting down server")

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelFunc()

	if err := httpServer.Shutdown(timeoutCtx); err != nil {
		s.log.Infof("Server shutdown error [%s]", err)
	} else {
		s.log.Info("Server exited")
	}

	c := make(chan struct{})
	go func() {
		defer close(c)
		s.onShutdown()
	}()

	select {
	case <-timeoutCtx.Done():
	case <-c:
	}
}
