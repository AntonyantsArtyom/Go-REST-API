package service

import (
	"fmt"
	"wallet/internal/repository"
	"wallet/models"
)

type TransactionService struct {
	transactionRepo repository.TransactionRepo
}

func NewTransactionService(r repository.TransactionRepo) *TransactionService {
	return &TransactionService{transactionRepo: r}
}

func (ts *TransactionService) GetRecentTransactions(count int) ([]models.Transaction, error) {
	if count <= 0 {
		return nil, fmt.Errorf("filter count must be greater than 0")
	}
	if count > 100 {
		return nil, fmt.Errorf("filter count must be less than 100")
	}
	return ts.transactionRepo.FindRecent(count)
}
