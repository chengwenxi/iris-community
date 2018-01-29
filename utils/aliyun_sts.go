package utils

import (
	"time"
	"github.com/tobyzxj/uuid"
	"fmt"
	"github.com/irisnet/iris-community/config"
	"encoding/json"
)

const (
		Format   = "JSON"
		SignatureMethod  = "HMAC-SHA1"
		SignatureVersion  = "1.0"
		Timestamp  = "2006-01-02T15:04:05Z"
		STS_Action  = "AssumeRole"
		)

type Sts struct {
	c *Client
}

func NewSls(client *Client) *Sts{
	return &Sts{
		c:client,
	}
}

// 创建一个新的请求参数
func (s Sts) newRequset() *Request {

	req := &Request{Param: make(map[string]string)}

	req.Put("Format", Format)
	req.Put("Version", s.c.Version)
	req.Put("AccessKeyId", s.c.AccessID)
	req.Put("SignatureMethod", SignatureMethod)
	req.Put("SignatureNonce", uuid.New())
	req.Put("SignatureVersion", SignatureVersion)
	req.Put("Timestamp", time.Now().UTC().Format(Timestamp))

	// 2. 业务API参数
	req.Put("Action", STS_Action)
	req.Put("RoleArn", s.c.Arn)
	req.Put("RoleSessionName", fmt.Sprintf("Role%d",time.Now().Unix()))
	req.Put("DurationSeconds", s.c.DurationSeconds)
	return req
}

//获取用户临时权限
func AssumeRole()(*StsResponse){
	var aliYun = config.Config.AliYun;

	acsClient := New(aliYun.AccessKeyId,aliYun.AccessKeySecret)
	acsClient.SetArn(aliYun.Sts.Arn)
	acsClient.SetEndPoint(aliYun.Sts.Endpoint)
	acsClient.SetVersion(aliYun.Sts.Version)
	acsClient.SetDurationSeconds(aliYun.Sts.DurationSeconds)

	req := NewSls(acsClient)
	//resp,httpCode,err := acsClient.send(req.newRequset())
	resp,_,_ := acsClient.send(req.newRequset())


	var stsResp StsResponse
	json.Unmarshal(resp,&stsResp)
	return &stsResp
}

type Credentials struct {
	AccessKeyId 	string	`json:"AccessKeyId"`
	AccessKeySecret string	`json:"AccessKeySecret"`
	Expiration 		string	`json:"Expiration"`
	SecurityToken 	string	`json:"SecurityToken"`
}

type AssumedRoleUser struct {
	Arn 				string	`json:"arn"`
	AssumedRoleUserId 	string	`json:"AssumedRoleUserId"`
}

type StsResponse struct {
	Credentials 	Credentials		`json:"Credentials"`
	AssumedRoleUser AssumedRoleUser	`json:"AssumedRoleUser"`
	RequestId string				`json:"RequestId"`
}
