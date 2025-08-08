package main

import (
	"fmt"
	"log"

	"wallet/api"
	"wallet/database"
	"wallet/models"

	"github.com/gin-gonic/gin"
)

func main() {
	databaseConnection, connectionError := database.ConnectDatabase()
	if connectionError != nil {
		log.Fatalf("connect error: %v", connectionError)
	}

	rawDatabase, getDatabaseError := databaseConnection.DB()
	if getDatabaseError != nil {
		log.Fatalf("get database error: %v", getDatabaseError)
	}

	defer rawDatabase.Close()

	automigrationError := databaseConnection.AutoMigrate(&models.Wallet{}, &models.Transaction{})
	if automigrationError != nil {
		log.Fatalf("automigration error: %v", automigrationError)
	}

	database.SeedWallets(databaseConnection)

	fmt.Println("database connected")

	router := gin.Default()
	api.RegisterRoutes(router, databaseConnection)
	router.Run(":8080")
}
