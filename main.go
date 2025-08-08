package main

import (
	"fmt"
	"log"

	"wallet/api"
	"wallet/database"

	"github.com/gin-gonic/gin"
)

func main() {
	databaseConnection, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}

	database, err := databaseConnection.DB()

	if err != nil {
		log.Fatalf("get database error: %v", err)
	}

	defer database.Close()

	fmt.Println("database connected")

	router := gin.Default()
	api.RegisterRoutes(router, databaseConnection)
	router.Run(":8080")
}
