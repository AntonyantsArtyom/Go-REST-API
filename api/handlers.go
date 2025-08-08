package api

import (
	"net/http"
	"strconv"

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
	countParam := ctx.DefaultQuery("count", "10")
	count, err := strconv.Atoi(countParam)

	switch {
	case err != nil:
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "count must be a number",
		})
		return
	case count <= 0:
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "count must be greater than 0",
		})
		return
	case count > 100:
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "count must be less than 100",
		})
		return
	}

	var transactions []models.Transaction
	databaseError := handler.databaseConnection.Limit(count).Order("created_at desc").Find(&transactions).Error
	if databaseError != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "database error: " + databaseError.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, TransactionsResponse{
		Transactions: transactions,
	})
}

func (handler *Handler) balanceHandler(ctx *gin.Context) {
	address := ctx.Param("address")

	var wallet models.Wallet
	var databaseError = handler.databaseConnection.First(&wallet, "address = ?", address).Error
	if databaseError != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{
			Error: "database error: " + databaseError.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, BalanceResponse{
		Balance: wallet.Balance,
	})
}
