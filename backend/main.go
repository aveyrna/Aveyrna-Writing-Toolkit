package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"backend/db"
	"backend/routes"

	"github.com/joho/godotenv"
)

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func main() {
	// Charge les variables d'env depuis .env si présent (dev)
	_ = godotenv.Load()

	// Init DB (assure-toi que db.Init() lit DATABASE_URL)
	db.Init()
	// si tu as db.Close(), pense à le defer ici :
	// defer db.Close()

	// Router (inclut CORS si ENABLE_CORS=true)
	r := routes.Router()

	port := getenv("PORT", "8080")
	addr := ":" + port

	srv := &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	fmt.Printf("🚀 Server running on http://localhost:%s\n", port)

	// Lancement serveur dans une goroutine
	errCh := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	// Capture SIGINT/SIGTERM pour un arrêt propre
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	select {
	case <-quit:
		fmt.Println("\n🛑 Shutting down gracefully...")
	case err := <-errCh:
		log.Fatalf("server error: %v", err)
	}

	// Contexte de shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("forced shutdown: %v", err)
	}

	fmt.Println("✅ Server stopped cleanly.")
}
