package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/irisnet/iris-community/utils"
	"net/http"
)


func RegisterAliyun(g *gin.RouterGroup) {
	//获取临时授权账号(文件上传)[acs:ram::1768586477174862:role/aliyunosswriteandreadrole 角色具有oss的读写权限]
	g.GET("/stsAuth", func(context *gin.Context) {
		resp := utils.AssumeRole()
		context.JSON(http.StatusOK, resp)
	})
}
