package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL is not set")
	}

	// Parse la config pour pouvoir ajuster les paramètres
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("❌ Error parsing DATABASE_URL: %v", err)
	}

	// ⚙️ Paramètres adaptés à OVH Web Cloud 1GB
	cfg.MaxConns = 20
	cfg.MinConns = 2
	cfg.MaxConnLifetime = 55 * time.Minute
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.HealthCheckPeriod = 30 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Création du pool
	Pool, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("❌ Unable to create connection pool: %v", err)
	}

	// Test de connexion
	if err := Pool.Ping(ctx); err != nil {
		log.Fatalf("❌ Unable to connect to database: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL OVH")
}
