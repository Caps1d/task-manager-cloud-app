package kv

import (
	"context"

	"github.com/Caps1d/task-manager-cloud-app/auth/config"
	"github.com/redis/go-redis/v9"
)

type kv struct {
	conn *redis.Client
}

func New(cfg config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.KVAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (r *kv) Get(key string) (string, error) {
	val, err := r.conn.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *kv) Put(key string, value int) error {
	err := r.conn.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *kv) Delete(key string) error {
	err := r.conn.Del(context.Background(), key).Err()
	if err != nil {
		return err
	}
	return nil
}
