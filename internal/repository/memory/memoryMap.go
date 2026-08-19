package memory

import (
	"Task-Wallet-Service/internal/domain"
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

func (u *InMemoryUserRepository) GetByID(id int64) (*domain.User, error) {
	u.mut.RLock()
	defer u.mut.RUnlock()

	if v, ok := u.mp[id]; ok {
		return &v, nil
	}

	return nil, domain.ErrUserNotFound
}

func (u *InMemoryUserRepository) UpdateBalance(id int64, balance int64) error {
	u.mut.Lock()
	defer u.mut.Unlock()

	if _, ok := u.mp[id]; ok {
		return nil
	}

	return domain.ErrUserNotFound
}
