package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type rdbClient struct {
	rdb *redis.Client
}
type Config struct {
	Addr         string
	Username     string
	Password     string
	DB           int
	MaxRetries   int
	DialTimeout  time.Duration
	PoolSize     int
	MinIdleConns int
}
type RedisSdk interface {
	RedisGet(string) ([]byte, error)
	RedisSet(string, interface{}, time.Duration) error
}

func RedisSdkI(c Config) RedisSdk {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Username:     c.Username,
		Password:     c.Password, // no password set
		DB:           c.DB,       // use default DB
		MaxRetries:   c.MaxRetries,
		DialTimeout:  c.DialTimeout,
		PoolSize:     c.PoolSize,
		MinIdleConns: c.MinIdleConns,
	})
	return &rdbClient{
		rdb: rdb,
	}
}
func (r rdbClient) RedisGet(key string) ([]byte, error) {
	data, err := r.rdb.Get(context.Background(), key).Bytes()
	if err != nil {
		return nil, err
	}
	return data, err
}
func (r rdbClient) RedisSet(key string, value interface{}, expiry time.Duration) error {
	err := r.rdb.Set(context.Background(), key, value, expiry).Err()
	if err != nil {
		return err
	}
	return nil
}
