package redis

import "github.com/go-redis/redis"

var RS *redis.Client

func RedisConfig() {

	opt, err := redis.ParseURL("redis://@172.20.0.138:6379")
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	RS = client
}
