package api

import (
	"net/http"
	"strconv"

	"wallet/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	WalletService      *service.WalletService
	TransactionService *service.TransactionService
	OperationSerice    *service.OperationService
}

func (h *Handler) sendHandler(ctx *gin.Context) {
	var req SendRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	err := h.OperationSerice.SendMoney(req.From, req.To, req.Amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, SendResponse{Message: "transaction successful"})
}

func (h *Handler) transactionsHandler(ctx *gin.Context) {
	countParam := ctx.DefaultQuery("count", "10")
	count, err := strconv.Atoi(countParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "count must be a number"})
		return
	}

	transactions, err := h.TransactionService.GetRecentTransactions(count)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, TransactionsResponse{Transactions: transactions})
}

func (handler *Handler) balanceHandler(ctx *gin.Context) {
	address := ctx.Param("address")

	balance, err := handler.WalletService.GetBalance(address)
	if err != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, BalanceResponse{
		Balance: balance,
	})
}
