package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	Client *redis.Client
}

func CreateClient() Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return Redis{Client: client}
}

func (redis Redis) SetNX(key string, value string, timeout time.Duration) {
	redis.Client.SetNX(key, value, timeout)
	redis.Client.Save()
}

func (redis Redis) Get(key string) (string, error) {
	val := redis.Client.Get(key).Val()
	if val == "" {
		return "", errors.New("no session")
	}
	return redis.Client.Get(key).Val(), nil
}

func (redis Redis) Delete(key string) {
	redis.Client.Del(key)
}
