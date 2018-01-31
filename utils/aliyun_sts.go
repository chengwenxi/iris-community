package utils

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/iris-community/config"
	myredis "github.com/irisnet/iris-community/models/redis"
	"github.com/tobyzxj/uuid"
	"time"
	"log"
	"github.com/garyburd/redigo/redis"
)

const (
	Format           = "JSON"
	SignatureMethod  = "HMAC-SHA1"
	SignatureVersion = "1.0"
	Timestamp        = "2006-01-02T15:04:05Z"
	STS_Action       = "AssumeRole"
	STS_ACCOUNT_KEY  = "STS_ACCOUNT_KEY"
)

type Sts struct {
	c *Client
}

func NewSls(client *Client) *Sts {
	return &Sts{
		c: client,
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
	req.Put("RoleSessionName", fmt.Sprintf("Role%d", time.Now().Unix()))
	req.Put("DurationSeconds", s.c.DurationSeconds)
	return req
}

//获取用户临时权限
func AssumeRole() *StsResponse {
	var aliYun = config.Config.AliYun

	redisCli := myredis.Pool.Get()

	if v, err := redisCli.Do("GET", STS_ACCOUNT_KEY); v != nil{
		var resp StsResponse
		re,_:= redis.Bytes(v,err)
		json.Unmarshal(re, &resp)
		return &resp
	}

	acsClient := New(aliYun.AccessKeyId, aliYun.AccessKeySecret)
	acsClient.SetArn(aliYun.Sts.Arn)
	acsClient.SetEndPoint(aliYun.Sts.Endpoint)
	acsClient.SetVersion(aliYun.Sts.Version)
	acsClient.SetDurationSeconds(aliYun.Sts.DurationSeconds)

	req := NewSls(acsClient)
	resp, _, err := acsClient.send(req.newRequset())
	log.Println(err)


	var stsResp StsResponse
	json.Unmarshal(resp, &stsResp)

	//存入redis
	redisCli.Do("SET", STS_ACCOUNT_KEY, resp)
	redisCli.Do("EXPIRE", STS_ACCOUNT_KEY, aliYun.Sts.DurationSeconds) //10 seconds expired

	return &stsResp
}

type Credentials struct {
	AccessKeyId     string `json:"AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret"`
	Expiration      string `json:"Expiration"`
	SecurityToken   string `json:"SecurityToken"`
}

type AssumedRoleUser struct {
	Arn               string `json:"arn"`
	AssumedRoleUserId string `json:"AssumedRoleUserId"`
}

type StsResponse struct {
	Credentials     Credentials     `json:"Credentials"`
	AssumedRoleUser AssumedRoleUser `json:"AssumedRoleUser"`
	RequestId       string          `json:"RequestId"`
}
