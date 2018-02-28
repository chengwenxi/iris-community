package main

import (
	"github.com/gin-gonic/gin"
	"github.com/irisnet/iris-community/rest"
	"github.com/irisnet/iris-community/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io"
	"os"
	"log"
	"github.com/irisnet/iris-community/config"
	"github.com/irisnet/iris-community/models/redis"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/irisnet/iris-community/docs"
)


// @title Swagger IRIS-Community API
// @version 1.0
// @description IRIS-Community API document

func main() {

	//init config
	if err := config.LoadConfiguration("./config.yml"); err != nil {
		log.Print("config error")
		return
	}

	r := gin.New()

	//log
	f, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.Logger())
	log.SetOutput(gin.DefaultWriter) // You may need this

	//authorizer
	//e := casbin.NewEnforcer("./authz/authz_model.conf", "./authz/authz_policy.csv")
	//r.Use(authz.NewAuthorizer(e))

	//init user and role by db
	//e.AddRoleForUser("test","admin")
	//e.AddRoleForUser("test2","user")
	//e.DeleteRoleForUser("test2","user")

	//static source
	r.Static("/static", config.Config.StaticPath)

	//db
	models.InitDB()
	redis.InitRedis()

	//rest
	rest.UserRegisterAll(r.Group("/user")) //user
	rest.AuthRegisterAll(r.Group("/auth")) //auth
	rest.RegisterKyc(r.Group("/kyc"))
	rest.RegisterAliyun(r.Group("/aliyun"))
	rest.QueryRegister(r.Group("/query"))
	rest.VerifyRegisterAll(r.Group("/verify"))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(config.Config.Server) // listen and serve on 0.0.0.0:8080
}
