package cache

import (
	"context"
	"fmt"
	"magic-bullet/backend/internal/config"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConnector struct {
	config *config.Config
	client *redis.Client
}

func NewRedisConnector(config *config.Config) (*RedisConnector, error) {
	if config == nil {
		return nil, fmt.Errorf("config parameter cannot be nil")
	}

	return &RedisConnector{
		config: config,
	}, nil
}

func (r *RedisConnector) Connect(ctx context.Context) error {
	address := fmt.Sprintf("%s:%d", r.config.Redis.Host, r.config.Redis.Port)

	r.client = redis.NewClient(&redis.Options{
		Addr: address,
		Password: r.config.Database.Password,
		DB: r.config.Redis.DB,
	})
	
	err := r.client.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to ping redis server: %v", err)
	}

	return nil
}

func (r *RedisConnector) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if r.client == nil {
		return fmt.Errorf("redis client is not initialized")
	} 

	err := r.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set key %s: %v", key, err)
	}

	return nil
}

func (r *RedisConnector) Get(ctx context.Context, key string) (string, error) {
	if r.client == nil {
		return "", fmt.Errorf("redis client is not initialized")
	} 

	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		// Key does not exist
		return "", fmt.Errorf("key %s not found", key)
	} else if err != nil {
		return "", fmt.Errorf("faile to get key %s: %v", key, err)
	}

	return val, nil
}

func (r *RedisConnector) Delete(ctx context.Context, key string) error {
		if r.client == nil {
		return fmt.Errorf("redis client is not initialized")
	} 

	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete key %s: %v", key, err)
	}

	return nil
}

func (r *RedisConnector) Close() error {
	if r.client != nil {
		err := r.client.Close()
		if err != nil {
			return fmt.Errorf("error closing redis connection: %v", err)
		}
		r.client = nil
	}

	return nil	
}
