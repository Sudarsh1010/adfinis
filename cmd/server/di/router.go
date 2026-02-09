package di

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sudarsh1010/adfinis/internal/interface/http/handlers"
	"github.com/sudarsh1010/adfinis/internal/interface/http/response"
	"go.uber.org/fx"
)

type RouterParams struct {
	fx.In

	HealthHandler *handlers.HealthHandler
	Logger        *Logger
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

	// Catch-all: Return 404 for unknown routes
	r.NotFound(func(w http.ResponseWriter, _ *http.Request) {
		response.NotFound(w, "Route not found")
	})

	return r
}
