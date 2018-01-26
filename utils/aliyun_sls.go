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
		Version  = "2015-04-01"
		SignatureMethod  = "HMAC-SHA1"
		SignatureVersion  = "1.0"
		Timestamp  = "2006-01-02T15:04:05Z"
		SLS_Action  = "AssumeRole"
		)

type Sls struct {
	c *Client
}

func NewSls(client *Client) *Sls{
	return &Sls{
		c:client,
	}
}

// 创建一个新的请求参数
func (s Sls) newRequset() *Request {

	req := &Request{Param: make(map[string]string)}

	req.Put("Format", Format)
	req.Put("Version", Version)
	req.Put("AccessKeyId", s.c.AccessID)
	req.Put("SignatureMethod", SignatureMethod)
	req.Put("SignatureNonce", uuid.New())
	req.Put("SignatureVersion", SignatureVersion)
	req.Put("Timestamp", time.Now().UTC().Format(Timestamp))

	// 2. 业务API参数
	req.Put("Action", SLS_Action)
	req.Put("RoleArn", s.c.Arn)
	req.Put("RoleSessionName", fmt.Sprintf("Role%d",time.Now().Unix()))
	req.Put("DurationSeconds", "3600")
	return req
}

//获取用户临时权限
func Auth()(*SlsResponse){
	var aliYun = config.Config.AliYun;

	acsClient := New(aliYun.AccessKeyId,aliYun.AccessKeySecret)
	acsClient.SetArn(aliYun.Sls.Arn)
	acsClient.SetEndPoint(aliYun.Sls.Endpoint)

	req := NewSls(acsClient)
	//resp,httpCode,err := acsClient.send(req.newRequset())
	resp,_,_ := acsClient.send(req.newRequset())


	var slsResponse SlsResponse
	json.Unmarshal(resp,&slsResponse)
	return &slsResponse
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

type SlsResponse struct {
	Credentials 	Credentials		`json:"Credentials"`
	AssumedRoleUser AssumedRoleUser	`json:"AssumedRoleUser"`
	RequestId string				`json:"RequestId"`
}
