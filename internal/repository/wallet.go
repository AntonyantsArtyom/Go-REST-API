package repository

import (
	"wallet/internal/models"

	"gorm.io/gorm"
)

type WalletRepo struct {
	db *gorm.DB
}

func NewWalletRepo(db *gorm.DB) *WalletRepo {
	return &WalletRepo{db: db}
}

func (wr *WalletRepo) FindByAddress(address string) (*models.Wallet, error) {
	var wallet models.Wallet
	err := wr.db.First(&wallet, "address = ?", address).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}
