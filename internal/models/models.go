package models

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Address string    `gorm:"uniqueIndex;not null"`
	Balance float64   `gorm:"type:numeric(10,2);not null"`
}

type Transaction struct {
	ID         uint    `gorm:"primaryKey"`
	FromWallet string  `gorm:"not null"`
	ToWallet   string  `gorm:"not null"`
	Amount     float64 `gorm:"not null;check:amount > 0"`
	CreatedAt  time.Time
}
