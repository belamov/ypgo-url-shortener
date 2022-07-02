package storage

import (
	"errors"
	"sync"

	"github.com/belamov/ypgo-url-shortener/internal/app/models"
)

type InMemoryRepository struct {
	storage map[string]models.ShortURL
	mutex   sync.RWMutex
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		storage: make(map[string]models.ShortURL),
		mutex:   sync.RWMutex{},
	}
}

func (repo *InMemoryRepository) Save(shortURL models.ShortURL) error {
	repo.mutex.RLock()
	_, ok := repo.storage[shortURL.ID]
	repo.mutex.RUnlock()

	if ok {
		return errors.New("not unique id")
	}

	repo.mutex.Lock()
	repo.storage[shortURL.ID] = shortURL
	repo.mutex.Unlock()

	return nil
}

func (repo *InMemoryRepository) GetByID(id string) (models.ShortURL, error) {
	repo.mutex.RLock()
	url, ok := repo.storage[id]
	repo.mutex.RUnlock()

	if !ok {
		return models.ShortURL{}, errors.New("can't find full url by id")
	}

	return url, nil
}
