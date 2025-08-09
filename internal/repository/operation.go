package repository

import (
	"wallet/internal/models"

	"gorm.io/gorm"
)

type OperationRepo struct {
	db *gorm.DB
}

func NewOperationsRepo(db *gorm.DB) *OperationRepo {
	return &OperationRepo{db}
}

func (r *OperationRepo) TransferBalance(sender *models.Wallet, receiver *models.Wallet, amount float64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		sender.Balance -= amount
		if err := tx.Save(sender).Error; err != nil {
			return err
		}

		receiver.Balance += amount
		if err := tx.Save(receiver).Error; err != nil {
			return err
		}

		return tx.Create(&models.Transaction{
			FromWallet: sender.Address,
			ToWallet:   receiver.Address,
			Amount:     amount,
		}).Error
	})
}
