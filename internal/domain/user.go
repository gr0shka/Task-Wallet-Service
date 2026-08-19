package domain

import (
	"context"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")
var ErrAmountMustZero = errors.New("amount must be greater than zero")

type User struct {
	ID      int64
	Balance int64
}

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (User, error)
	UpdateBalance(ctx context.Context, id int64, balance int64) error
}

type UserService interface {
	Deposit(ctx context.Context, userID int64, amount int64) (User, error)
}
