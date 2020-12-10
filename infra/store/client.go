package store

import (
	"os"

	"github.com/gomodule/redigo/redis"
)

func NewRedisClient() (*redis.Conn, error) {
	pool := &redis.Pool{
		MaxIdle:   50,
		MaxActive: 10000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", os.Getenv("REDIS_PORT"))
			if err != nil {
				os.Exit(1)
			}
			return conn, nil
		},
	}
	conn, err := pool.Dial()
	if err != nil {
		return nil, err
	}
	return &conn, err
}
