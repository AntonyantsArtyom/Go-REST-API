package service

import (
	"wallet/internal/repository"
)

type WalletService struct {
	walletRepo repository.WalletRepo
}

func NewWalletService(r repository.WalletRepo) *WalletService {
	return &WalletService{walletRepo: r}
}

func (ws *WalletService) GetBalance(address string) (float64, error) {
	wallet, err := ws.walletRepo.FindByAddress(address)

	if err != nil {
		return 0, err
	}

	return wallet.Balance, err
}
