package repo

import (
	"fmt"
	"log"
	"sync"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/configs"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/model"
)

var (
	userStore UserStore
	onceInit  sync.Once
)

type UserStore interface {
	Save(user *model.User) error
	Find(username string) (*model.User, error)
}

func init() {
	onceInit.Do(func() {
		userStore = NewInMemoryUserStore()
		err := configs.SeedUsers(userStore)
		if err != nil {
			log.Fatal("could not create seed users")
		}
	})
}

type inMemoryUserStore struct {
	mutex sync.RWMutex
	users map[string]*model.User
}

func NewInMemoryUserStore() UserStore {
	return userStore
}

func (store *inMemoryUserStore) Save(user *model.User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.users[user.Username] != nil {
		return fmt.Errorf("user(%s) already exists", user.Username)
	}

	store.users[user.Username] = user.Clone()
	return nil
}

func (store *inMemoryUserStore) Find(username string) (*model.User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user := store.users[username]
	if user == nil {
		return nil, nil
	}

	return user.Clone(), nil
}
