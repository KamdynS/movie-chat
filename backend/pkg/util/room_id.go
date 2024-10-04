package util

import (
	"crypto/rand"

	"github.com/google/uuid"
)

func GenerateRoomID() (uuid.UUID, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return uuid.Nil, err
	}
	return uuid.New(), nil
}
