package redis

import (
	"time"
	"flag"
	"github.com/garyburd/redigo/redis"
)

func InitRedis()  {
	Pool = NewPool(*redisServer)
}

func NewPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		Dial: func () (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

var (
	Pool *redis.Pool
	redisServer = flag.String("redisServer", "localhost:6379", "")
)
