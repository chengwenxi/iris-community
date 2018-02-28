package redis

import (
	"time"
	"flag"
	"github.com/garyburd/redigo/redis"
	"github.com/irisnet/iris-community/config"
)

func InitRedis() {
	redisServer := flag.String("redisServer", config.Config.Redis.Url, "")
	Pool = NewPool(*redisServer)
}

func NewPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

var (
	Pool *redis.Pool
)
