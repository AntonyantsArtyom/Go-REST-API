package repository

import (
	"wallet/models"

	"gorm.io/gorm"
)

type WalletRepo struct {
	db *gorm.DB
}

func NewWalletRepo(db *gorm.DB) *WalletRepo {
	return &WalletRepo{db: db}
}

func (w *WalletRepo) FindByAddress(address string) (*models.Wallet, error) {
	var wallet models.Wallet
	err := w.db.First(&wallet, "address = ?", address).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}
