package server

import (
	"context"
	"github.com/Worwulew/Songs-library/config"
	"github.com/Worwulew/Songs-library/pkg/logger"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

// Server is server struct
type Server struct {
	cfg    *config.Config
	db     *sqlx.DB
	logger logger.Logger
}

// NewServer is Server constructor
func NewServer(cfg *config.Config, db *sqlx.DB, logger logger.Logger) *Server {
	return &Server{cfg: cfg, db: db, logger: logger}
}

// Run starts server
func (s *Server) Run() error {
	handler, err := s.MapHandlers()
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		Handler:        handler,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil {
			s.logger.Fatalf("Error starting Server: %v", err)
		}
	}()

	/*go func() {
		s.logger.Infof("Starting Debug Server on PORT: %s", s.cfg.Server.PprofPort)
		if err := http.ListenAndServe(s.cfg.Server.PprofPort, http.DefaultServeMux); err != nil {
			s.logger.Errorf("Error PPROF ListenAndServe: %s", err)
		}
	}()*/

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	s.logger.Info("Server Exited Properly")
	return server.Shutdown(ctx)
}
