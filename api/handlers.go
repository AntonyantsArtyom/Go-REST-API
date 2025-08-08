package api

import (
	"net/http"
	"strconv"

	"wallet/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	databaseConnection *gorm.DB
}

func (handler *Handler) sendHandler(ctx *gin.Context) {
	var request SendRequest

	parsingError := ctx.ShouldBindJSON(&request)
	if parsingError != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error: parsingError.Error(),
		})
		return
	}

	var sender models.Wallet
	senderFindError := handler.databaseConnection.First(&sender, "address = ?", request.From).Error

	switch {
	case senderFindError != nil:
		ctx.JSON(http.StatusNotFound, ErrorResponse{
			Error: "database error: " + senderFindError.Error(),
		})
		return
	case sender.Balance < request.Amount:
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "not enough balance",
		})
		return
	}

	var receiver models.Wallet
	receiverFindError := handler.databaseConnection.First(&receiver, "address = ?", request.To).Error

	if receiverFindError != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{
			Error: "database error: " + receiverFindError.Error(),
		})
		return
	}

	transactionError := handler.databaseConnection.Transaction(func(tx *gorm.DB) error {
		sender.Balance -= request.Amount
		err := tx.Save(&sender).Error
		if err != nil {
			return err
		}

		receiver.Balance += request.Amount
		err = tx.Save(&receiver).Error
		if err != nil {
			return err
		}

		return tx.Create(&models.Transaction{
			FromWallet: request.From,
			ToWallet:   request.To,
			Amount:     request.Amount,
		}).Error
	})

	if transactionError != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "database error: " + transactionError.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, SendResponse{
		Message: "transaction successful",
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
