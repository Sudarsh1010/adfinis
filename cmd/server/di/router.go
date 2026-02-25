package di

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sudarsh1010/adfinis/internal/interface/http/handlers"
	"github.com/sudarsh1010/adfinis/internal/interface/http/response"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type RouterParams struct {
	fx.In

	HealthHandler *handlers.HealthHandler
	SPAHandler    *handlers.SPAHandler
	Logger        *zap.Logger
	// WorkspaceHandler  *handlers.WorkspaceHandler
	// ConnectionHandler *handlers.ConnectionHandler
}

func NewRouter(p RouterParams) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
			"http://localhost:3000",
			"http://localhost:*",
		},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
		},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           HTTPMaxAge,
	}))

	// Routes
	r.Route("/api", func(r chi.Router) {
		// Health check
		r.Get("/health", p.HealthHandler.Handle)

		// 	// Workspace routes
		// 	r.Route("/workspaces", func(r chi.Router) {
		// 		r.Post("/", p.WorkspaceHandler.Create)
		// 		r.Get("/", p.WorkspaceHandler.List)
		// 		r.Get("/{id}", p.WorkspaceHandler.Get)
		// 		r.Put("/{id}/activate", p.WorkspaceHandler.Activate)
		// 	})
		//
		// 	// Connection routes
		// 	r.Route("/connections", func(r chi.Router) {
		// 		r.Post("/", p.ConnectionHandler.Create)
		// 		r.Get("/", p.ConnectionHandler.List)
		// 		r.Post("/{id}/test", p.ConnectionHandler.Test)
		// 	})
	})

	// SPA routing: Serve frontend for non-API routes, 404 for unknown API routes
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			response.NotFound(w, "Route not found")
			return
		}
		p.SPAHandler.ServeHTTP(w, r)
	})

	return r
}
