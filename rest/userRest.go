package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/irisnet/iris-community/models"
	"strconv"
	"net/http"
	"github.com/irisnet/iris-community/utils"
	myredis "github.com/irisnet/iris-community/models/redis"
	"github.com/garyburd/redigo/redis"
	"github.com/irisnet/iris-community/config"
	"github.com/pborman/uuid"
	"log"
)

func UserRegisterAll(g *gin.RouterGroup) {
	g.POST("", CreateUser)
	g.GET("/activate", ActivateUser)
	g.POST("/resendAct", ResendAct)
	g.PUT("/updatePwd", UpdateUserPwd)
	g.POST("/reset", ResetUser)
}

type RequestUsers struct {
	Email string `binding:"required"`
}

type RequestUpateUsers struct {
	RequestUsers
	Password string `binding:"required"`
}

type RequestCreateUsers struct {
	RequestUsers
	Password       string `binding:"required"`
	InvitationCode string
	VerifyCode     string `binding:"required"`
}

//create user
func CreateUser(c *gin.Context) {
	var req RequestCreateUsers
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !VerifyCode(req.Email, req.VerifyCode) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "verify code error"})
		return
	}
	user1, _ := models.FindUserByEmail(req.Email)
	if user1.Id != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email has been registered."})
		return
	}
	password := utils.Md5(req.Password)
	salt := utils.RandomInfo(6)
	user := &models.Users{
		Email:    req.Email,
		Salt:     salt,
		Password: utils.Sha1s(salt + password),
	}
	if dbErr := user.Create(req.InvitationCode); dbErr == nil {
		con := myredis.Pool.Get()
		uid := uuid.NewUUID().String()
		_, err := con.Do("SET", "actc_"+uid, user.Id)
		_, err = con.Do("EXPIRE", "actc_"+uid, config.Config.Redis.ActcTimeout) //10 seconds expired
		if err != nil {
			log.Println(err)
		}
		utils.RegisterEmail(user.Email, strconv.Itoa(int(user.Id)), uid)
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": dbErr.Error()})
	}

}

//activate user
func ActivateUser(c *gin.Context) {
	id := c.Query("id")
	code := c.Query("code")
	if id == "" || code == "" {
		c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	con := myredis.Pool.Get()
	v, _ := redis.String(con.Do("GET", "actc_"+code))
	if v == id {
		iid, _ := strconv.Atoi(id)
		user := &models.Users{
			Id: uint(iid),
		}
		user.ActivateUser()
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}
}

//resend email to activate user
func ResendAct(c *gin.Context) {
	var req RequestUsers
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := models.FindUserByEmail(req.Email)
	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email does not exist"})
		return
	}
	con := myredis.Pool.Get()
	uid := uuid.NewUUID().String()
	_, err = con.Do("SET", "actc_"+uid, user.Id)
	_, err = con.Do("EXPIRE", "actc_"+uid, config.Config.Redis.RescTimeout) //10 seconds expired
	if err != nil {
		log.Println(err)
		return
	}
	if err = utils.RegisterEmail(user.Email, strconv.Itoa(int(user.Id)), uid); err == nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "send email fail"})
	}

}

//reset user password by email
func ResetUser(c *gin.Context) {
	var req RequestUsers
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(req.Email) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	user, _ := models.FindUserByEmail(req.Email)
	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email does not exist"})
		return
	}
	con := myredis.Pool.Get()
	uid := uuid.NewUUID().String()
	_, err = con.Do("SET", "resc_"+uid, user.Id)
	_, err = con.Do("EXPIRE", "resc_"+uid, config.Config.Redis.RescTimeout) //10 seconds expired
	if err != nil {
		log.Println(err)
		return
	}
	if err = utils.ResetEmail(user.Email, strconv.Itoa(int(user.Id)), uid); err == nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "send email fail"})
	}

}

//update password
func UpdateUserPwd(c *gin.Context) {

	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no authorization"})
		return
	}
	userAuth := &models.UserAuth{
		AuthCode: authorization,
	}
	userAuth.FindByAuth()
	if userAuth.Id == 0 {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}

	var req RequestUpateUsers
	if err := c.ShouldBindJSON(&req); err == nil {
		if len(req.Password) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
			return
		}
		user := &models.Users{
			Id: uint(userAuth.UserId),
		}
		user.First()
		salt := utils.RandomInfo(6)
		password := utils.Sha1s(salt + utils.Md5(req.Password))
		user.UpdatePwd(salt, password)
		c.JSON(http.StatusOK, user)
		return
	}
	c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
}
