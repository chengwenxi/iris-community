package rest

import (
	"testing"
	"github.com/irisnet/iris-community/models/redis"
)

func TestVerifyCode(t *testing.T) {
	redis.InitRedis()
	code := "81541"
	println(VerifyCode("760329367@qq.com", code))
}
