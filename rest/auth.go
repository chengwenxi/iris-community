package rest

import (
	"github.com/irisnet/iris-community/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/irisnet/iris-community/utils"
	myredis "github.com/irisnet/iris-community/models/redis"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

func AuthRegisterAll(g *gin.RouterGroup) {
	g.POST("", Login)
	g.GET("/reset", AuthRest)
	g.GET("/user", AuthUser)
}

func Login(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err == nil {
		if len(user.Email) == 0 || len(user.Password) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
			return
		}
		user1, _ := models.FindUserByEmail(user.Email)
		if user1.Id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email not exist"})
			return
		}
		password := utils.Md5(user.Password)
		salt := user1.Salt
		if user1.Password == utils.Sha1s(salt+password) {
			if user1.IsBlocked {
				c.JSON(http.StatusOK, gin.H{"error": "You've been blacklisted"})
				return
			}
			userAuth := &models.UserAuth{
				UserId: user1.Id,
			}
			userAuth.Create()
			c.Header("Authorization", userAuth.AuthCode)
			c.JSON(http.StatusOK, userAuth)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email or password error"})
		}

	}
}

func AuthUser(c *gin.Context) {
	if authorization := c.Request.Header.Get("Authorization"); authorization == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no authorization"})
	} else {
		userAuth := &models.UserAuth{
			AuthCode: authorization,
		}
		userAuth.FindByAuth()
		if userAuth.Id != 0 {
			user := &models.Users{
				Id: uint(userAuth.UserId),
			}
			user.First()
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization error"})
		}
	}
}

func AuthRest(c *gin.Context) {
	id := c.Query("id")

	code := c.Query("code")
	if id == "" || code == "" {
		c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	con := myredis.Pool.Get()
	v, _ := redis.String(con.Do("GET", "resc_"+code))
	if v == id {
		iid, _ := strconv.Atoi(id)
		user := &models.Users{
			Id: uint(iid),
		}
		user.First()
		if user.IsBlocked {
			c.JSON(http.StatusOK, gin.H{"error": "You've been blacklisted"})
			return
		}
		userAuth := &models.UserAuth{
			UserId: uint(iid),
		}
		userAuth.Create()
		c.Header("Authorization", userAuth.AuthCode)
		c.JSON(http.StatusOK, userAuth)
	} else {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}
}
