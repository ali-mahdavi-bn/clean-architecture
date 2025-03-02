package cache

import (
	"clean-hex/pkg/framwork/infrastructure/redisx"
	"clean-hex/pkg/framwork/service_layer/types"
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"strings"
	"time"
)

type Store interface {
	GetValue(ctx context.Context, key string, value interface{}) error
	SetValue(ctx context.Context, key string, value interface{}, exp time.Duration) error
	Cache(ctx context.Context, key string, value interface{}, exp time.Duration, fn types.RedisUseCase) error
	CreateKey(key ...interface{}) string
	DeleteKey(ctx context.Context, key string) error
}

type RedisStore struct {
	store redisx.Connection
}

func NewRedisStore(store redisx.Connection) Store {
	return &RedisStore{store: store}
}
func (r *RedisStore) CreateKey(key ...interface{}) string {
	output := make([]string, len(key))
	for _, v := range key {
		output = append(output, cast.ToString(v))
	}
	return strings.Join(output, "#")
}

func (r *RedisStore) Cache(ctx context.Context, key string, value interface{}, exp time.Duration, fn types.RedisUseCase) error {
	err := r.GetValue(ctx, key, value)
	if !errors.Is(err, redis.Nil) {
		return err
	}
	value, err = fn(ctx)
	if err != nil {
		return err
	}
	return r.SetValue(ctx, key, value, exp)
}

func (r *RedisStore) GetValue(ctx context.Context, key string, value interface{}) error {
	cached, err := r.store.GetValue(ctx, key)
	if err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(cached), value); err != nil {
		return err
	}
	return nil
}

func (r *RedisStore) SetValue(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.store.SetValue(ctx, key, string(jsonData), exp)
}

func (r *RedisStore) DeleteKey(ctx context.Context, key string) error {
	return r.store.DeleteKey(ctx, key)
}
