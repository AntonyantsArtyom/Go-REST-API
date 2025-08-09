package service

import (
	"fmt"
	"wallet/internal/repository"
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
		return err
	}

	if sender.Balance < amount {
		return fmt.Errorf("not enough balance")
	}

	receiver, err := os.walletRepo.FindByAddress(to)
	if err != nil {
		return err
	}

	return os.operationRepo.TransferBalance(sender, receiver, amount)
}
