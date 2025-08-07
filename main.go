package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"wallet/api"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}
	defer db.Close()
	fmt.Println("connected")

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}

func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "wallet_db"),
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func getEnv(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
