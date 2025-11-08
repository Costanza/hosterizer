package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hosterizer/shared/database"
)

//go:embed ../../migrations/*.sql
var migrations embed.FS

func main() {
	var (
		action = flag.String("action", "up", "Migration action: up, down, version")
		host   = flag.String("host", getEnv("DB_HOST", "localhost"), "Database host")
		port   = flag.Int("port", getEnvInt("DB_PORT", 5432), "Database port")
		user   = flag.String("user", getEnv("DB_USER", "postgres"), "Database user")
		pass   = flag.String("password", getEnv("DB_PASSWORD", "postgres"), "Database password")
		dbname = flag.String("database", getEnv("DB_NAME", "hosterizer"), "Database name")
		ssl    = flag.String("sslmode", getEnv("DB_SSLMODE", "disable"), "SSL mode")
	)

	flag.Parse()

	cfg := database.Config{
		Host:     *host,
		Port:     *port,
		User:     *user,
		Password: *pass,
		Database: *dbname,
		SSLMode:  *ssl,
	}

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Printf("Connected to database: %s@%s:%d/%s", cfg.User, cfg.Host, cfg.Port, cfg.Database)

	switch *action {
	case "up":
		log.Println("Running migrations up...")
		if err := database.MigrateUp(db.DB, migrations, "migrations"); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		log.Println("Migrations completed successfully")

	case "down":
		log.Println("Rolling back last migration...")
		if err := database.MigrateDown(db.DB, migrations, "migrations"); err != nil {
			log.Fatalf("Failed to rollback migration: %v", err)
		}
		log.Println("Rollback completed successfully")

	case "version":
		version, dirty, err := database.MigrateVersion(db.DB, migrations, "migrations")
		if err != nil {
			log.Fatalf("Failed to get migration version: %v", err)
		}
		if dirty {
			log.Printf("Current migration version: %d (dirty)", version)
		} else {
			log.Printf("Current migration version: %d", version)
		}

	default:
		log.Fatalf("Unknown action: %s (use: up, down, version)", *action)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		var intValue int
		if _, err := fmt.Sscanf(value, "%d", &intValue); err == nil {
			return intValue
		}
	}
	return defaultValue
}
