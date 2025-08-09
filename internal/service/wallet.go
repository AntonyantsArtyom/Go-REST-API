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

// Возвращает баланс кошелька по адресу и ошибку (при успехе nil)
//
// Параметры:
//   - address: адрес кошелька
//
// Возможные ошибки:
//   - ErrWalletNotFound, если кошелек не найден
//   - error, если произошла другая ошибка
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
