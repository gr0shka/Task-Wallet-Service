package memory

import (
	"Task-Wallet-Service/internal/domain"
	"context"
	"sync"
)

type InMemoryUserRepository struct {
	mp  map[int64]*domain.User
	mut sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		mp:  make(map[int64]*domain.User),
		mut: sync.RWMutex{},
	}
}

func (u *InMemoryUserRepository) AddUser(ctx context.Context, user *domain.User) error {
	u.mut.Lock()
	defer u.mut.Unlock()

	u.mp[user.ID] = user
	return nil
}

func (u *InMemoryUserRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	u.mut.RLock()
	defer u.mut.RUnlock()

	if v, ok := u.mp[id]; ok {
		return *v, nil
	}

	return domain.User{}, domain.ErrUserNotFound
}

func (u *InMemoryUserRepository) UpdateBalance(ctx context.Context, id int64, balance int64) error {
	u.mut.Lock()
	defer u.mut.Unlock()

	if _, ok := u.mp[id]; !ok {
		return domain.ErrUserNotFound
	}
	u.mp[id].Balance = balance

	return nil
}
