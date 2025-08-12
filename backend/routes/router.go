package routes

import (
	"net/http"
	"os"
	"strings"
	"time"

	"backend/routes/characters"
	"backend/routes/projects"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	// Middlewares utiles
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(httprate.LimitByIP(100, 1*time.Minute)) // petit rate-limit safe

	// CORS activable via env si besoin (prod/staging)
	if enableCORS() {
		origins := allowedOrigins()
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   origins,
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}))
	}

	// (Optionnel) healthcheck rapide
	r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	r.Route("/api", func(api chi.Router) {
		api.Mount("/projects", projects.Routes())
		api.Mount("/characters", characters.Routes())
	})

	return r
}

func enableCORS() bool {
	return strings.ToLower(os.Getenv("ENABLE_CORS")) == "true"
}

func allowedOrigins() []string {
	val := os.Getenv("CORS_ORIGINS")
	if val == "" {
		// valeur par défaut raisonnable
		return []string{"http://localhost:5173", "http://127.0.0.1:5173"}
	}
	parts := strings.Split(val, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}
