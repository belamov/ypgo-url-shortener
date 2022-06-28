package random

import (
	"crypto/rand"
	"github.com/google/uuid"
)

func GenerateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateUserId() string {
	return uuid.NewString()
}