package rest

import (
	"github.com/irisnet/iris-community/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/irisnet/iris-community/utils"
)

func AuthRegisterAll(g *gin.RouterGroup) {
	g.POST("", Login)
	g.GET("/user", AuthUser)
}

func Login(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err == nil {
		if len(user.Email) == 0 || len(user.Password) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
			return
		}
		user1, dbErr := models.FindUserByEmail(user.Email)
		if dbErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
			return
		}
		if user1.Id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email not exist"})
			return
		}
		password := utils.Md5(user.Password)
		salt := user1.Salt
		if user1.Password == utils.Sha1s(salt+password) {
			userAuth := &models.UserAuth{
				UserId: user1.Id,
			}
			userAuth.Create()
			c.Header("Authorization",userAuth.AuthCode)
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
