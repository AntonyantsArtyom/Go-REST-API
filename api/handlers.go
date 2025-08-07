package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/api/send", sendHandler)
	r.GET("/api/transactions", transactionsHandler)
	r.GET("/api/wallet/:address/balance", balanceHandler)
}

func sendHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "placeholder",
	})
}

func transactionsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "placeholder",
	})
}

func balanceHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "placeholder",
	})
}
