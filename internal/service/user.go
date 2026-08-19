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
		v, err := u.rep.GetByID(ctx, userID)
		if err != nil {
			return domain.User{}, err
		}

		newBalance := v.Balance + amount

		err = u.rep.UpdateBalance(ctx, userID, newBalance)
		if err != nil {
			return domain.User{}, err
		}

		return domain.User{v.ID, newBalance}, nil
	}
}
