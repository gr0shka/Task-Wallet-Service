package service

import (
	"Task-Wallet-Service/internal/domain"
	"context"
)

type UserService struct {
	rep domain.UserRepository
}

func NewUserService(rep domain.UserRepository) *UserService {
	return &UserService{rep}
}

func (u *UserService) Deposit(ctx context.Context, userID int64, amount int64) (domain.User, error) {
	if amount <= 0 {
		return domain.User{}, domain.ErrAmountMustZero
	}

	select {
	case <-ctx.Done():
		return domain.User{}, domain.ContextErr
	default:
		err := u.rep.UpdateBalance(userID, amount)
		if err != nil {
			return domain.User{}, err
		}

		v, err := u.rep.GetByID(userID)
		if err != nil {
			return domain.User{}, err
		}
		
		return *v, nil
	}
}
