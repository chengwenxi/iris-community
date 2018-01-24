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
			c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		user.Password = utils.Md5(user.Password)
		if users, err := models.AuthUser(user.Email, user.Password); err == nil && users.Email == user.Email {
			userAuth := &models.UserAuth{
				UserId: users.Id,
			}
			userAuth.Create()
			c.JSON(http.StatusOK, userAuth)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email or password error"})
		}
	}
}

func AuthUser(c *gin.Context) {
	if authorization := c.Request.Header.Get("Authorization"); authorization == "" {
		c.JSON(http.StatusForbidden, gin.H{"error": "no authorization"})
	} else {

	}

}
