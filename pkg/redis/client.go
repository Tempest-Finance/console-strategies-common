package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
)

var redisClient redis.UniversalClient

func InitClient(config Config) error {
	if redisClient == nil {
		redisAddresses := config.InitAddress
		if len(redisAddresses) == 0 {
			return errors.New("redis host is empty")
		}

		redisClient = newRedisClient(redisAddresses, config.Username, config.Password)
		if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
			return err
		}
	}
	return nil
}

func ClientInstance() redis.UniversalClient {
	return redisClient
}

func newRedisClient(redisAddresses []string, masterName string, pwd string) redis.UniversalClient {
	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      redisAddresses,
		MasterName: masterName,
		Password:   pwd,
	})
}
