package rest

import (
	"github.com/irisnet/iris-community/models"
	"github.com/gin-gonic/gin"
	"encoding/base64"
	"net/http"
)

func AuthRegisterAll(g *gin.RouterGroup) {
	g.POST("", Login)
}

func Login(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err == nil {
		if users, err := models.AuthUser(user.Email, user.Password); err == nil && users.Email == user.Email {
			authorization := "Basic " + base64.StdEncoding.EncodeToString([]byte(user.Email+":"+user.Password))
			c.Header("Authorization", authorization)
			c.Writer.Write([]byte(authorization))
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email or password error"})
		}
	}
}
