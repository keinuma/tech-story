package store

import (
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type Store struct {
	Client redis.UniversalClient
	TTL    time.Duration
}

func NewRedisClient() (*Store, error) {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})
	err := client.Ping(nil).Err()
	if err != nil {
		return nil, err
	}
	defaultTTL := 24 * time.Hour
	return &Store{Client: client, TTL: defaultTTL}, err
}
