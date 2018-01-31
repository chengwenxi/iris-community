package utils

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pborman/uuid"
	"bytes"
	"github.com/irisnet/iris-community/config"
	"os"
)



func UploadByLocal(filepath string)  (string,error) {
	var aliYun = config.Config.AliYun;
	client, err := oss.New(aliYun.Oss.Endpoint, aliYun.AccessKeyId, aliYun.AccessKeySecret)
	if err != nil {
		return "",err
	}

	fd, err := os.Open(filepath)

	bucket, err := client.Bucket(aliYun.Oss.Bucket)
	if err != nil {
		return "",err
	}

	ossKey := uuid.New()

	err = bucket.PutObject(ossKey, fd)
	if err != nil {
		return "",err
	}
	return ossKey,nil
}

func UploadByBytes(content []byte) (string,error) {
	var aliYun = config.Config.AliYun;
	client, err := oss.New(aliYun.Oss.Endpoint, aliYun.AccessKeyId, aliYun.AccessKeySecret)
	if err != nil {
		return "",err
	}

	bucket, err := client.Bucket(aliYun.Oss.Bucket)
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
