package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo/config"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	echo   *echo.Echo
	config *config.Configuration
	db     *gorm.DB
	logger *logrus.Logger
	ready  chan bool
}

func NewServer(config *config.Configuration, db *gorm.DB, logger *logrus.Logger, ready chan bool) *Server {
	return &Server{echo: echo.New(), config: config, db: db, logger: logger, ready: ready}
}

func (server *Server) Bootstrap() error {
	httpServer := &http.Server{
		Addr:         ":" + server.config.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		server.logger.Logf(logrus.InfoLevel, "Server is listening on PORT: %s", server.config.Port)
		if err := server.echo.StartServer(httpServer); err != nil {
			server.logger.Fatalln("Error starting Server: ", err)
		}
	}()

	if server.ready != nil {
		server.ready <- true
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	server.logger.Fatalln("Server Exited Properly")
	return server.echo.Server.Shutdown(ctx)
}
