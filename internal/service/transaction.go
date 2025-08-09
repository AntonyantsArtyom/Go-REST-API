package service

import (
	"wallet/internal/models"
	"wallet/internal/repository"
)

type TransactionService struct {
	transactionRepo repository.TransactionRepo
}

func NewTransactionService(r repository.TransactionRepo) *TransactionService {
	return &TransactionService{transactionRepo: r}
}

func (ts *TransactionService) GetRecentTransactions(count int) ([]models.Transaction, error) {
	return ts.transactionRepo.FindRecent(count)
}
