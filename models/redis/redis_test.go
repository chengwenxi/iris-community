package redis

import (
	"testing"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func TestRedis(t *testing.T) {
	InitRedis()
	v, err := Pool.Get().Do("SET", "name", "red")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	v, err = Pool.Get().Do("EXPIRE", "name", 3) //10 seconds expired
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	v, err = Pool.Get().Do("GET", "name")
	if err != nil {
		fmt.Println(err)
		return
	}
	if s, err := redis.String(v, err); err == nil {
		fmt.Println(s)
	}

	time.Sleep(3 * time.Second)
	v, err = Pool.Get().Do("GET", "name")
	if err != nil {
		fmt.Println(err)
		return
	}
	if s, err := redis.String(v, err); err == nil {
		fmt.Println(s)
	} else {
		fmt.Println(err)
	}
}
