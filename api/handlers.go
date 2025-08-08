package api

import (
	"net/http"

	"wallet/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, databaseConnection *gorm.DB) {
	handler := &Handler{databaseConnection}

	router.POST("/api/send", handler.sendHandler)
	router.GET("/api/transactions", handler.transactionsHandler)
	router.GET("/api/wallet/:address/balance", handler.balanceHandler)
}

type Handler struct {
	databaseConnection *gorm.DB
}

func (handler *Handler) sendHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "",
	})
}

func (handler *Handler) transactionsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "",
	})
}

func (handler *Handler) balanceHandler(ctx *gin.Context) {
	address := ctx.Param("address")

	var wallet models.Wallet
	var databaseError = handler.databaseConnection.First(&wallet, "address = ?", address).Error
	if databaseError != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{
			Error: "wallet not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, BalanceResponse{
		Balance: wallet.Balance,
	})
}
