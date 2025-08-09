package service

import (
	"errors"
	"wallet/internal/models"
	"wallet/internal/repository"
)

var (
	ErrTransactionNotFound = errors.New("transactions empty")
)

type TransactionService struct {
	transactionRepo repository.TransactionRepo
}

func NewTransactionService(r repository.TransactionRepo) *TransactionService {
	return &TransactionService{transactionRepo: r}
}

func (ts *TransactionService) GetRecentTransactions(count int) ([]models.Transaction, error) {
	transactions, err := ts.transactionRepo.FindRecent(count)

	switch {
	case len(transactions) == 0:
		return nil, ErrTransactionNotFound
	case err != nil:
		return nil, err
	}

	return transactions, err
}
