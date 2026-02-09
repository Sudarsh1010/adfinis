package di

import (
	"context"
	"fmt"
	go_http "net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type Server struct {
	server *go_http.Server
	logger *Logger
}

type ServerParams struct {
	fx.In

	Config    *Config
	Logger    *Logger
	Router    *chi.Mux
	Lifecycle fx.Lifecycle
}

func NewServer(p ServerParams) (*Server, error) {
	srv := &go_http.Server{
		Addr:         fmt.Sprintf(":%d", p.Config.Port),
		Handler:      p.Router,
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimout,
		IdleTimeout:  IdleTimeout,
	}

	server := &Server{
		server: srv,
		logger: p.Logger,
	}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			p.Logger.Info(
				"Starting HTTP server",
				"port",
				p.Config.Port,
				"env",
				p.Config.Env,
			)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			p.Logger.Info("Shutting down HTTP server...")
			return srv.Shutdown(ctx)
		},
	})

	return server, nil
}

func (s *Server) Start() error {
	err := s.server.ListenAndServe()
	if err != nil && err != go_http.ErrServerClosed {
		return fmt.Errorf("server failed: %w", err)
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
