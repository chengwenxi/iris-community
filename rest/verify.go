package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
	myredis "github.com/irisnet/iris-community/models/redis"
	"github.com/irisnet/iris-community/config"
	"github.com/garyburd/redigo/redis"
	"strings"
	"math/rand"
	"strconv"
)

func VerifyRegisterAll(g *gin.RouterGroup) {
	g.GET("", CreateCode)
}

func CreateCode(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	var configD = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	}
	var seed string
	for i := 0; i < configD.CaptchaLen; i++ {
		seed += strconv.Itoa(rand.Intn(10))
	}
	dig := base64Captcha.EngineDigitsCreate(seed, configD)
	verifyValue := dig.VerifyValue
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(dig)
	con := myredis.Pool.Get()
	_, err := con.Do("SET", "verc_"+email, verifyValue)
	_, err = con.Do("EXPIRE", "verc_"+email, config.Config.Redis.VercTimeOut) //20 seconds expired
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"code": base64stringC})
		return
	}
	c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func VerifyCode(email string, code string) bool {
	con := myredis.Pool.Get()
	v, _ := redis.String(con.Do("GET", "verc_"+email))
	return strings.ToLower(v) == strings.ToLower(code)
}
