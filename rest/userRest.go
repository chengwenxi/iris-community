package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/irisnet/iris-community/models"
	"strconv"
	"net/http"
	"github.com/irisnet/iris-community/utils"
)

func UserRegisterAll(g *gin.RouterGroup) {
	g.GET("", ListUser)
	g.GET("/:id", FindUser)
	g.POST("", CreateUser)
	g.PUT("", UpdateUser)
	g.DELETE("/:id", DeleteUser)
}

func ListUser(c *gin.Context) {
	skip, err := strconv.Atoi(c.DefaultQuery("skip", "0"))
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	users, err := models.UserList(skip, limit)
	if err != nil {

	}
	c.JSON(http.StatusOK, users)
}

func FindUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user := &models.Users{
		Id: uint(id),
	}
	user.First()
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err == nil {
		if len(user.Email) == 0 || len(user.Password) == 0{
			c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		user.Password = utils.Md5(user.Password)
		if dbErr := user.Create(); dbErr == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": dbErr.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user := &models.Users{
		Id: uint(id),
	}
	if err := user.Delete(); err == nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func UpdateUser(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err == nil {
		if len(user.Email) == 0 {
			c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if dbErr := user.Update(); dbErr == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": dbErr.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
