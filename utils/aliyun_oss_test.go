package utils

import (
	"testing"
	"fmt"
	//"github.com/irisnet/iris-community/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"github.com/pborman/uuid"
	"encoding/json"
	"github.com/irisnet/iris-community/models/redis"
	"github.com/irisnet/iris-community/config"
	"log"
)


func TestAssumeRole(t *testing.T) {
	redis.InitRedis()
	//init config
	if err := config.LoadConfiguration("../config.yml"); err != nil {
		log.Print("config error")
		return
	}
	resp := AssumeRole()

	bytes,err:= json.Marshal(resp)
	fmt.Println(string(bytes))
	fmt.Println(err.Error())

}

func TestUploadByBytes(t *testing.T) {
	Endpoint := "oss-cn-shanghai.aliyuncs.com"
	AccessKeyId := "STS.CuVLxDGECfNeZocZEinm2KQ4g"
	AccessKeySecret := "9sXjm4bj3KJruKzFhcp8WubtbpNQdZqDEikyW4T6ioHz"
	StsToken := "CAIShgJ1q6Ft5B2yfSjIorbjB8Lwqppi0YyOWEnSvkU8YuIepJSfhTz2IHhOfHhuB+sWtP40nWpR6v8SlqdJTJtIHBGdMpAutssNrFtR1TpmR4jng4YfgbiJREKxaXeiruKwDsz9SNTCAITPD3nPii50x5bjaDymRCbLGJaViJlhHL91N0vCGlggPtpNIRZ4o8I3LGbYMe3XUiTnmW3NFkFlyGEe4CFdkf3gnJ3BukSC0wKjlbNF+d7LT8L6P5U2DvBWSMyo2eF6TK3F3RNL5gJCnKUM1/QeoWaZ4IrNWQMKuk/fb7OE6KxmKA5oe649ALVbtvngmPR+tvbenojtzBJALUe908llvHrLGoABB6hQLkyYiPTzm+C1JorKyJTyrnc1LJcEN9PSiRa2qEBSwHh/0JVCRIUSIOyBdiLcQVcJo94ofngAxEaN749oOMcnaFOx9TpcOR8PdTZDS+Iz8w9m0Dicu9fNlhjKqrEjSGoapE5bpqFDNjvfmKoWBfracdgPjgiU1fixH95fE6Q="
	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret, oss.SecurityToken(StsToken))
	if err != nil {
		// HandleError(err)
	}

	bucket, err := client.Bucket("bianjie-default")
	if err != nil {
		// HandleError(err)
	}

	fd, err := os.Open("./mail.go")
	if err != nil {
		// HandleError(err)
	}
	defer fd.Close()

	ossKey := uuid.New()

	fmt.Println("ossKey=",ossKey)

	err = bucket.PutObject(ossKey, fd)
	if err != nil {
		fmt.Println(err)
	}
}