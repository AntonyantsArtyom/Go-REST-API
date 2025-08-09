package service

import (
	"errors"
	"wallet/internal/repository"

	"gorm.io/gorm"
)

var (
	ErrWalletNotFound = errors.New("wallet empty")
)

type WalletService struct {
	walletRepo repository.WalletRepo
}

func NewWalletService(r repository.WalletRepo) *WalletService {
	return &WalletService{walletRepo: r}
}

func (ws *WalletService) GetBalance(address string) (float64, error) {
	wallet, err := ws.walletRepo.FindByAddress(address)

	switch {
	case err == gorm.ErrRecordNotFound:
		return 0, ErrWalletNotFound
	case err != nil:
		return 0, err
	}

	return wallet.Balance, err
}
