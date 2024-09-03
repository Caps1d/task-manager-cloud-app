package sessions

import (
	"strconv"

	"github.com/google/uuid"
)

type Store interface {
	Put(string, int32) error
	Get(string) (string, error)
	Delete(string) error
}

func Create(kv Store, userID int32) (string, error) {
	// generate sessionID -> uuid
	sessionID := uuid.New().String()
	err := kv.Put(sessionID, userID)
	if err != nil {
		return "", err
	}

	return sessionID, nil
}

func Get(kv Store, sessionID string) (int32, error) {
	userID, err := kv.Get(sessionID)
	if err != nil {
		return 0, err
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		return 0, err
	}
	return int32(id), nil
}

func Destroy(kv Store, sessionID string) error {
	err := kv.Delete(sessionID)
	if err != nil {
		return err
	}

	return nil
}
