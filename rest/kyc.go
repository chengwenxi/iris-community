package rest

import ("github.com/gin-gonic/gin"
	"github.com/irisnet/iris-community/models"
	"github.com/irisnet/iris-community/utils"
	"log"
	"net/http"
)

type Kyc struct {
	CertificateTypeId uint `form:"certificate_type"`
	CertificateNum string `form:"certificate_num"`
	CountryId uint `form:"country_id"`
	FrontFile []byte `form:"front_file"`
	ReverseFile []byte `form:"reverse_file"`
	HandFile []byte `form:"hand_file"`

	FamilyName string `form:"family_name"`
	Name string `form:"name"`
}

func RegisterKyc(g *gin.RouterGroup) {

	//查询国家列表
	g.GET("/country", func(context *gin.Context) {
		country,err:= models.CountryList()
		if err != nil {
			log.Printf("Country is Empty,please init ")
		}
		context.JSON(http.StatusOK, country)
	})

	//查询证件类型列表
	g.GET("/cerType", func(context *gin.Context) {
		cerTypes,err:= models.CenficateTypeList()
		if err != nil {
			log.Printf("cerType is Empty,please init ")
		}
		context.JSON(http.StatusOK, cerTypes)
	})

	//用户实名认证
	g.POST("/cerficate", func(context *gin.Context) {
		var kyc Kyc
		if context.ShouldBind(&kyc) == nil {
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

	//用户上传文件
	//g.POST("/upload", func(context *gin.Context) {
	//	//得到上传的文件
	//	file, header, err := context.Request.FormFile("image") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	//	if err != nil {
	//		context.String(http.StatusBadRequest, "Bad request")
	//		return
	//	}
	//	t := time.Now().Unix()
	//
	//	//文件的名称
	//	fileId := fmt.Sprintf("%d-%s",t,header.Filename)
	//
	//	//创建文件
	//	out, err := os.Create("static/uploadfile/" + fileId)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	defer out.Close()
	//
	//	_, err = io.Copy(out, file)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	context.String(http.StatusOK, fileId)
	//})


}

//提交用户实行认证信息
func postKyc(kyc Kyc,authCode string) error{
	//1:上传用户证件图片到阿里云服务
	frontFileKey,reverseFileKey,handFileKey,err := uploadImage(kyc)
	if err != nil {
		return err
	}

	//2：保存用户证件信息到数据库
	frontFile := models.Files{
		OssKey:frontFileKey,
	}

	reverseFile := models.Files{
		OssKey:reverseFileKey,
	}

	handFile := models.Files{
		OssKey:handFileKey,
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

	//3:保存用户证件信息
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

	//4:更新用户信息
	userAuth := models.UserAuth{
		AuthCode:authCode,
	}

	userAuth.FindByAuth()

	userProfile := models.UserProfile{
		UserId:userAuth.UserId,
	}

	userProfile.First()

	err = tx.Model(&userProfile).Updates(models.UserProfile{
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

func uploadImage(kyc Kyc) (string,string,string,error) {
	frontFileKey,err := utils.UploadByBytes(kyc.FrontFile)
	if err != nil {
		return "","","",err;
	}

	reverseFileKey,err := utils.UploadByBytes(kyc.ReverseFile)

	if err != nil {
		return "","","",err;
	}

	handFileKey,err := utils.UploadByBytes(kyc.HandFile)

	if err != nil {
		return "","","",err;
	}
	return frontFileKey,reverseFileKey,handFileKey,nil
}
