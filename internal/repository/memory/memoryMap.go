package memory

import (
	"Task-Wallet-Service/internal/domain"
	"context"
	"sync"
)

type InMemoryUserRepository struct {
	mp  map[int64]domain.User
	mut sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		mp:  make(map[int64]domain.User),
		mut: sync.RWMutex{},
	}
}

func NewInMemoryUserRepositoryWithUsers(users []domain.User) *InMemoryUserRepository {
	rep := &InMemoryUserRepository{
		mp:  make(map[int64]domain.User),
		mut: sync.RWMutex{},
	}

	for _, user := range users {
		rep.mp[user.ID] = user
	}

	return rep
}

func (u *InMemoryUserRepository) AddUser(ctx context.Context, user *domain.User) error {
	u.mut.Lock()
	defer u.mut.Unlock()

	u.mp[user.ID] = *user
	return nil
}

func (u *InMemoryUserRepository) GetByID(id int64) (domain.User, error) {
	u.mut.RLock()
	defer u.mut.RUnlock()

	if v, ok := u.mp[id]; ok {
		return v, nil
	}

	return domain.User{}, domain.ErrUserNotFound
}

func (u *InMemoryUserRepository) UpdateBalance(id int64, balance int64) (domain.User, error) {
	u.mut.Lock()
	defer u.mut.Unlock()

	user, ok := u.mp[id]
	if !ok {
		return domain.User{}, domain.ErrUserNotFound
	}

	user.Balance = balance
	u.mp[id] = user

	return user, nil
}
