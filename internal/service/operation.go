package service

import (
	"errors"
	"wallet/internal/repository"

	"gorm.io/gorm"
)

var (
	ErrSenderWalletNotFound   = errors.New("sender wallet not found")
	ErrReceiverWalletNotFound = errors.New("receiver wallet not found")
	ErrInsufficientBalance    = errors.New("insufficient balance")
)

type OperationService struct {
	walletRepo    repository.WalletRepo
	operationRepo repository.OperationRepo
}

func NewOperationService(or repository.OperationRepo, wr repository.WalletRepo) *OperationService {
	return &OperationService{operationRepo: or, walletRepo: wr}
}

func (os *OperationService) SendMoney(from, to string, amount float64) error {
	sender, err := os.walletRepo.FindByAddress(from)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrSenderWalletNotFound
		}
		return err
	}

	if sender.Balance < amount {
		return ErrInsufficientBalance
	}

	receiver, err := os.walletRepo.FindByAddress(to)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrReceiverWalletNotFound
		}
		return err
	}

	return os.operationRepo.TransferBalance(sender, receiver, amount)
}
