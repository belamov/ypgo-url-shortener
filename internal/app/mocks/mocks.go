package mocks

import (
	"github.com/belamov/ypgo-url-shortener/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Save(shortURL models.ShortURL) error {
	args := m.Called(shortURL)
	return args.Error(0)
}

func (m *MockRepo) GetByID(id string) (models.ShortURL, error) {
	args := m.Called(id)
	return models.ShortURL{OriginalURL: args.String(0), ID: id}, args.Error(1)
}

type MockGen struct {
	mock.Mock
}

func (m *MockGen) GenerateIDFromString(str string) (string, error) {
	args := m.Called(str)
	return args.String(0), args.Error(1)
}

type MockUserIDGenerator struct {
	mock.Mock
}

func (m *MockUserIDGenerator) GenerateUserId() string {
	args := m.Called()
	return args.String(0)
}
