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
	walletRepo := *repository.NewWalletRepo(db)
	transactionRepo := *repository.NewTransactionRepo(db)
	operationRepo := *repository.NewOperationsRepo(db)

	handler := &Handler{
		service.NewWalletService(walletRepo),
		service.NewTransactionService(transactionRepo),
		service.NewOperationService(operationRepo, walletRepo),
	}

	router.POST("/api/send", handler.sendHandler)
	router.GET("/api/transactions", handler.transactionsHandler)
	router.GET("/api/wallet/:address/balance", handler.balanceHandler)
}
