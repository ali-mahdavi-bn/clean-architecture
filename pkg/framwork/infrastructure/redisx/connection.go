package redisx

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Connection interface {
	GetValue(ctx context.Context, key string) (string, error)
	SetValue(ctx context.Context, key string, s string, exp time.Duration) error
	DeleteKey(ctx context.Context, key string) error
}
type connection struct {
	client *redis.Client
}

func NewRedisConnection(ctx context.Context, options *redis.Options) (*connection, error) {
	client := redis.NewClient(options)
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &connection{client: client}, nil
}

func (r *connection) GetValue(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *connection) ExistsKey(ctx context.Context, key string) (bool, error) {
	exists, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func (r *connection) SetValue(ctx context.Context, key string, value string, exp time.Duration) error {
	return r.client.Set(ctx, key, value, exp).Err()
}

func (r *connection) DeleteKey(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *connection) PushValues(ctx context.Context, key string, values []string, exp time.Duration) error {
	if err := r.client.LPush(ctx, key, values).Err(); err != nil {
		return err
	}
	if exp > 0 {
		if err := r.client.Expire(ctx, key, exp).Err(); err != nil {
			return err
		}
	}
	return nil
}

func (r *connection) GetAllListValues(ctx context.Context, key string) ([]string, error) {
	values, err := r.client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (r *connection) HealthCheck(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}
