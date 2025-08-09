package repository

import (
	"wallet/internal/models"

	"gorm.io/gorm"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *TransactionRepo {
	return &TransactionRepo{db}
}

// Получает из базы последние N транзакций
//
// Параметры:
//   - сount: количество необходимых транзакций
func (tr *TransactionRepo) FindRecent(count int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := tr.db.Limit(count).Order("created_at desc").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
