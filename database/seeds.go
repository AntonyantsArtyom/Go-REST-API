package database

import (
	"log"

	"wallet/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedWallets(db *gorm.DB) {
	var count int64
	err := db.Model(&models.Wallet{}).Count(&count).Error
	if err != nil {
		log.Printf("error checking wallet count: %v", err)
		return
	}

	if count > 0 {
		log.Println("wallets already seeded")
		return
	}

	for i := range 10 {
		wallet := models.Wallet{
			Address: uuid.New().String(),
			Balance: 100.0,
		}
		if err := db.Create(&wallet).Error; err != nil {
			log.Printf("error creating wallet #%d: %v", i+1, err)
		}
	}

	log.Println("10 wallets successfully seeded")
}
