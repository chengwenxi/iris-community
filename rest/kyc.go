package rest

import ("github.com/gin-gonic/gin"
	"github.com/irisnet/iris-community/models"
	"github.com/irisnet/iris-community/utils"
	"log"
	"net/http"
)

type Kyc struct {
	CertificateTypeId uint
	CertificateNum string
	CountryId uint
	FrontFileKey string
	ReverseFileKey string
	HandFileKey string

	FamilyName string
	Name string
}

func RegisterKyc(g *gin.RouterGroup) {

	//查询国家列表
	g.GET("/country", func(context *gin.Context) {
		country,err:= models.Country().List()
		if err != nil {
			log.Printf("Country is Empty,please init ")
		}
		context.JSON(http.StatusOK, country)
	})

	//查询证件类型列表
	g.GET("/cerType", func(context *gin.Context) {
		cerTypes,err:= models.CerficateType().List()
		if err != nil {
			log.Printf("cerType is Empty,please init ")
		}
		context.JSON(http.StatusOK, cerTypes)
	})

	//用户实名认证
	g.POST("/cerficate", func(context *gin.Context) {
		var kyc Kyc
		if context.ShouldBindJSON(&kyc) == nil {
			authCode := context.Request.Header.Get("Authorization")
			if err := postKyc(kyc,authCode);err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"code":"fail","msg": "user certify fail"})
			}else {
				context.JSON(http.StatusOK, gin.H{"code":"success","msg": "user certify success"})
			}
		}else {
			context.JSON(http.StatusBadRequest, "invalide json")
		}
	})

	//获取临时授权账号(文件上传)
	g.GET("/slsAuth", func(context *gin.Context){
		resp := utils.Auth()
		context.JSON(http.StatusOK, resp)
	})

}

//提交用户实行认证信息
func postKyc(kyc Kyc,authCode string) error{
	//1：保存用户证件信息到数据库
	frontFile := models.Files{
		OssKey:kyc.FrontFileKey,
	}

	reverseFile := models.Files{
		OssKey:kyc.ReverseFileKey,
	}

	handFile := models.Files{
		OssKey:kyc.HandFileKey,
	}

	//开启事物
	tx := models.DB.Begin()

	if err := tx.Create(&frontFile).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&reverseFile).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&handFile).Error; err != nil {
		tx.Rollback()
		return err
	}

	//2:保存用户证件信息
	cerficates := models.Cerficates{
		TypeId:kyc.CertificateTypeId,
		Num:kyc.CertificateNum,
		FrontFileId:frontFile.Id,
		ReverseFileId:reverseFile.Id,
		HandFileId:handFile.Id,
	}

	if err := tx.Create(&cerficates).Error; err != nil {
		tx.Rollback()
		return err
	}

	//3:更新用户信息
	userAuth := models.UserAuth{
		AuthCode:authCode,
	}

	userAuth.FindByAuth()

	userProfile := models.UserProfile{
		UserId:userAuth.UserId,
	}

	userProfile.First()

	err := tx.Model(&userProfile).Updates(models.UserProfile{
		FamilyName:kyc.FamilyName,
		Name:kyc.Name,
		CountryId:kyc.CountryId,
		CerficateId:cerficates.Id,
	}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}
