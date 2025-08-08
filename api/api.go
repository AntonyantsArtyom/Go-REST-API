package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"wallet/internal/repository"
	"wallet/internal/service"
)

// Подключает в маршрутизатор Gin логику обработки роутов.
//
// Параметры:
//   - router: маршрутизатор Gin
//   - db: подключение к базе
//
// Роуты приложения:
//   - POST /api/send — перевод средств между кошельками
//   - GET /api/transactions — получение последних транзакций
//   - GET /api/wallet/:address/balance — получение баланса кошелька
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	var walletRepo = repository.NewWalletRepo(db)
	var WalletService = service.NewWalletService(*walletRepo)

	var transactionRepo = repository.NewTransactionRepo(db)
	var TransactionService = service.NewTransactionService(*transactionRepo)

	handler := &Handler{WalletService, TransactionService}

	//router.POST("/api/send", handler.sendHandler)
	router.GET("/api/transactions", handler.transactionsHandler)
	router.GET("/api/wallet/:address/balance", handler.balanceHandler)
}
