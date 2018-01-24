package utils

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pborman/uuid"
	"bytes"
	"github.com/irisnet/iris-community/config"
)

var aliYun = config.Config.AliYun;

func UploadByLocal(filename string)  (string,error) {
	client, err := oss.New(aliYun.Endpoint, aliYun.AccessKeyId, aliYun.AccessKeySecret)
	if err != nil {
		return "",err
	}

	bucket, err := client.Bucket(aliYun.Bucket)
	if err != nil {
		return "",err
	}

	ossKey := uuid.New()

	err = bucket.PutObjectFromFile(ossKey, filename)
	if err != nil {
		return "",err
	}
	return ossKey,nil
}

func UploadByBytes(content []byte) (string,error) {
	client, err := oss.New(aliYun.Endpoint, aliYun.AccessKeyId, aliYun.AccessKeySecret)
	if err != nil {
		return "",err
	}

	bucket, err := client.Bucket(aliYun.Bucket)
	if err != nil {
		return "",err
	}

	ossKey := uuid.New()
	err = bucket.PutObject(ossKey, bytes.NewReader(content))
	if err != nil {
		return "",err
	}
	return ossKey,nil
}
