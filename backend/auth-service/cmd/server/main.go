package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hosterizer/auth-service/internal/handler"
	"github.com/hosterizer/auth-service/internal/repository"
	"github.com/hosterizer/auth-service/internal/service"
	"github.com/hosterizer/shared/database"
)

func main() {
	log.Println("Auth Service starting...")

	// Load configuration from environment
	dbConfig := loadDBConfig()
	jwtSecret := getEnv("JWT_SECRET", "your-secret-key-change-in-production")
	redisAddr := getEnv("REDIS_ADDR", "localhost:6379")
	redisPassword := getEnv("REDIS_PASSWORD", "")
	port := getEnv("PORT", "8001")

	// Initialize database connection
	db, err := database.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Database connection established")

	// Initialize repositories
	userRepo := repository.NewPostgresUserRepository(db.DB)

	// Initialize services
	passwordSvc := service.NewPasswordService()
	jwtSvc := service.NewJWTService(service.JWTConfig{
		SecretKey:            jwtSecret,
		AccessTokenDuration:  15 * time.Minute,
		RefreshTokenDuration: 7 * 24 * time.Hour,
	})
	mfaSvc := service.NewMFAService("Hosterizer")
	lockoutSvc := service.NewLockoutService(userRepo, service.LockoutConfig{
		MaxFailedAttempts: 3,
		LockoutDuration:   15 * time.Minute,
	})
	sessionSvc, err := service.NewSessionService(service.SessionConfig{
		RedisAddr:      redisAddr,
		RedisPassword:  redisPassword,
		RedisDB:        0,
		SessionTimeout: 30 * time.Minute,
	})
	if err != nil {
		log.Fatalf("Failed to initialize session service: %v", err)
	}
	defer sessionSvc.Close()

	authSvc := service.NewAuthService(service.AuthServiceConfig{
		UserRepo:    userRepo,
		PasswordSvc: passwordSvc,
		JWTSvc:      jwtSvc,
		MFASvc:      mfaSvc,
		LockoutSvc:  lockoutSvc,
		SessionSvc:  sessionSvc,
	})

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authSvc, jwtSvc)

	// Setup HTTP server
	mux := http.NewServeMux()
	authHandler.RegisterRoutes(mux)

	// Add health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Auth Service listening on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}

func loadDBConfig() database.Config {
	return database.Config{
		Host:            getEnv("DB_HOST", "localhost"),
		Port:            getEnvAsInt("DB_PORT", 5432),
		User:            getEnv("DB_USER", "postgres"),
		Password:        getEnv("DB_PASSWORD", "postgres"),
		Database:        getEnv("DB_NAME", "hosterizer"),
		SSLMode:         getEnv("DB_SSLMODE", "disable"),
		MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
		MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 5),
		ConnMaxLifetime: time.Duration(getEnvAsInt("DB_CONN_MAX_LIFETIME", 5)) * time.Minute,
		ConnMaxIdleTime: time.Duration(getEnvAsInt("DB_CONN_MAX_IDLE_TIME", 10)) * time.Minute,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		var intValue int
		if _, err := fmt.Sscanf(value, "%d", &intValue); err == nil {
			return intValue
		}
	}
	return defaultValue
}
