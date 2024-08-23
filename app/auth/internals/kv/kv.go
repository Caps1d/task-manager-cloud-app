package kv

import (
	"context"
	"errors"

	"github.com/Caps1d/task-manager-cloud-app/auth/config"
	"github.com/redis/go-redis/v9"
)

type KV struct {
	conn *redis.Client
}

func New(cfg config.Config) *KV {
	return &KV{
		conn: redis.NewClient(&redis.Options{
			Addr:     cfg.KVAddr,
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

func (r *KV) Get(key string) (string, error) {
	val, err := r.conn.Get(context.Background(), key).Result()
	if errors.Is(err, redis.Nil) {
		return "", err
	}
	return val, nil
}

func (r *KV) Put(key string, value int) error {
	err := r.conn.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *KV) Delete(key string) error {
	err := r.conn.Del(context.Background(), key).Err()
	if err != nil {
		return err
	}
	return nil
}
