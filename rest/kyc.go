package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/irisnet/iris-community/models"
	"log"
	"net/http"
	"errors"
)

type Kyc struct {
	CertificateTypeId uint
	CertificateNum    string
	CountryId         uint
	FrontFileKey      string
	ReverseFileKey    string
	HandFileKey       string

	FamilyName string
	Name       string
}

type KycInfo struct {
	Kyc    Kyc
	Result KycResult
}

type KycResult struct {
	Status string
	Errors []models.DimApprovalFailedReason
}

func RegisterKyc(g *gin.RouterGroup) {

	//查询国家列表
	g.GET("/country", func(context *gin.Context) {
		country, err := models.Country().List()
		if err != nil {
			log.Printf("Country is Empty,please init ")
		}
		context.JSON(http.StatusOK, country)
	})

	//查询证件类型列表
	g.GET("/cerType", func(context *gin.Context) {
		cerTypes, err := models.CerficateType().List()
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
			if err := postKyc(kyc, authCode); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"code": "fail", "msg": "user certify fail"})
			} else {
				context.JSON(http.StatusOK, gin.H{"code": "success", "msg": "user certify success"})
			}
		} else {
			context.JSON(http.StatusBadRequest, "invalide json")
		}
	})
	//查询用户认证信息
	g.GET("/info", func(context *gin.Context) {
		authCode := context.Request.Header.Get("Authorization")
		kycinfo, err := queryKyc(authCode)
		if err == nil {
			context.JSON(http.StatusOK, kycinfo)
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"code": "fail", "msg": err.Error()})
		}
	})

}

//提交用户实行认证信息
func postKyc(kyc Kyc, authCode string) error {
	//1：保存用户证件信息到数据库
	frontFile := models.Files{
		OssKey: kyc.FrontFileKey,
	}

	reverseFile := models.Files{
		OssKey: kyc.ReverseFileKey,
	}

	handFile := models.Files{
		OssKey: kyc.HandFileKey,
	}

	//开启事物
	tx := models.DB.Begin()

	if err := tx.FirstOrCreate(&frontFile, models.Files{OssKey: kyc.FrontFileKey}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.FirstOrCreate(&reverseFile, models.Files{OssKey: kyc.ReverseFileKey}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.FirstOrCreate(&handFile, models.Files{OssKey: kyc.HandFileKey}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//2:保存用户证件信息
	cer := models.Cerficates{
		TypeId:        kyc.CertificateTypeId,
		Num:           kyc.CertificateNum,
		FrontFileId:   frontFile.Id,
		ReverseFileId: reverseFile.Id,
		HandFileId:    handFile.Id,
	}

	if err := tx.FirstOrCreate(&cer, models.Cerficates{
		TypeId:        cer.TypeId,
		FrontFileId:   cer.FrontFileId,
		ReverseFileId: cer.ReverseFileId,
		HandFileId:    cer.HandFileId,}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//3:更新用户信息
	userAuth := models.UserAuth{
		AuthCode: authCode,
	}

	userAuth.FindByAuth()

	userProfile := models.UserProfile{
		UserId: userAuth.UserId,
	}

	userProfile.First()

	if err := tx.Model(&userProfile).Updates(models.UserProfile{
		FamilyName:  kyc.FamilyName,
		Name:        kyc.Name,
		CountryId:   kyc.CountryId,
		CerficateId: cer.Id,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//4:插入用户审核表
	approval := models.UserApproval{
		UserId:         userAuth.UserId,
		ApprovalStatus: "p",
	}

	if err := tx.FirstOrCreate(&approval, models.UserApproval{UserId: userAuth.UserId}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}

//查询用户实名认证信息
func queryKyc(authCode string) (KycInfo, error) {
	var kycInfo KycInfo

	//查询用户信息
	userAuth := models.UserAuth{
		AuthCode: authCode,
	}

	if err := userAuth.FindByAuth(); err != nil {
		return kycInfo, errors.New("user does not login")
	}

	userProfile := models.UserProfile{
		UserId: userAuth.UserId,
	}

	if err := userProfile.First(); err != nil {
		return kycInfo, errors.New("user does not exist")
	}

	//查询用户提交证件信息
	cer := models.NewCerficates()
	if err := cer.Query(userProfile.CerficateId); err != nil {
		return kycInfo, errors.New("user's cerficates  does not exist")
	}

	frontFile, _ := models.NewFiles().QueryById(cer.FrontFileId)
	reverseFile, _ := models.NewFiles().QueryById(cer.ReverseFileId)
	handFile, _ := models.NewFiles().QueryById(cer.HandFileId)

	kyc := Kyc{
		CertificateTypeId: cer.TypeId,
		CertificateNum:    cer.Num,
		FrontFileKey:      frontFile.OssKey,
		ReverseFileKey:    reverseFile.OssKey,
		HandFileKey:       handFile.OssKey,

		FamilyName: userProfile.FamilyName,
		Name:       userProfile.Name,
		CountryId:  userProfile.CountryId,
	}

	//查询用户审核结果
	approval := models.NewUserApproval(userAuth.UserId)

	if err := approval.QueryById(); err != nil {
		return kycInfo, errors.New("approval is pedding,please wait")
	}

	//如果认证失败，查询失败原因
	reasons := []models.DimApprovalFailedReason{}
	if approval.ApprovalStatus == "f" {
		failedReason, _ := models.NewUserApprovalXFailedReason().QueryByUserId(userAuth.UserId)
		for _, r := range failedReason {
			if rx, err := models.NewReason().QueryByUserId(r.ReasonId); err == nil {
				reasons = append(reasons, rx)
			}
		}
	}

	kycInfo = KycInfo{
		Kyc: kyc,
		Result: KycResult{
			Status: approval.ApprovalStatus,
			Errors: reasons,
		},
	}

	return kycInfo, nil

}
