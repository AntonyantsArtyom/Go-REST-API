package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Устанавливает соединение с базой данных. Настройки берутся из следующих переменных окружения:
//   - DB_HOST: хост базы данных (по умолчанию: localhost)
//   - DB_USER: пользователь базы данных (по умолчанию: postgres)
//   - DB_PASSWORD: пароль пользователя базы данных (по умолчанию: postgres)
//   - DB_NAME: название базы данных (по умолчанию: wallet_db)
//   - DB_PORT: порт подключения к базе данных (по умолчанию: 5432)
func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "wallet_db"),
		getEnv("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
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
