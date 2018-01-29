package utils

import (
	"testing"
	"fmt"
	"github.com/irisnet/iris-community/config"
)


func TestAssumeRole(t *testing.T) {
	//init config
	if err := config.LoadConfiguration("../config.yml"); err!=nil{
		fmt.Print("config error")
		return
	}

	var aliYun = config.Config.AliYun;

	acsClient := New(aliYun.AccessKeyId,aliYun.AccessKeySecret)
	acsClient.SetArn(aliYun.Sls.Arn)
	acsClient.SetEndPoint(aliYun.Sls.Endpoint)

	acsClient.SetVersion(aliYun.Sls.Version)

	req := NewSls(acsClient)
	resp,httpCode,err := acsClient.send(req.newRequset())


	fmt.Println(string(resp))
	fmt.Println(string(httpCode))
	fmt.Println(err)



}