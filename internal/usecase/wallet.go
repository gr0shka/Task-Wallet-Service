package usecase

import (
	"Task-Wallet-Service/internal/domain"
	"context"
)

type walletUsecase struct {
	repository domain.UserRepository
}

func NewWalletUsecase(repository domain.UserRepository) domain.WalletUsecase {
	return &walletUsecase{repository: repository}
}

func (w *walletUsecase) Deposit(ctx context.Context, userID int64, amount int64) (domain.User, error) {
	if amount <= 0 {
		return domain.User{}, domain.ErrAmountMustZero
	}

	select {
	case <-ctx.Done():
		return domain.User{}, ctx.Err()
	default:
	}

	user, err := w.repository.GetByID(userID)
	if err != nil {
		return domain.User{}, err
	}

	updatedUser, err := w.repository.UpdateBalance(userID, user.Balance+amount)
	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}
