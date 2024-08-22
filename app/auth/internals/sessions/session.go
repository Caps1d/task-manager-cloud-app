package sessions

import (
	"strconv"

	"github.com/Caps1d/task-manager-cloud-app/auth/internals/kv"
	"github.com/google/uuid"
)

func Create(kv *kv.KV, userID int) (string, error) {
	// generate sessionID -> uuid
	sessionID := uuid.New().String()
	err := kv.Put(sessionID, userID)
	if err != nil {
		return "", err
	}

	return sessionID, nil
}

func Get(kv *kv.KV, sessionID string) (int64, error) {
	userID, err := kv.Get(sessionID)
	if err != nil {
		return 0, err
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}

func Destroy(kv *kv.KV, sessionID string) error {
	err := kv.Delete(sessionID)
	if err != nil {
		return err
	}

	return nil
}
