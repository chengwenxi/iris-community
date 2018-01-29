// Package acmp Copyright 2016 The GiterLab Authors. All rights reserved.
package utils

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
	"github.com/GiterLab/urllib"
)

// HTTPDebugEnable http调试开关
var HTTPDebugEnable = false

// Request 请求参数设置
type Request struct {
	Param map[string]string
}

// Put 添加请求参数
func (r *Request) Put(key, value string) error {
	if r != nil {
		if r.Param == nil {
			r.Param = make(map[string]string)
		}
		r.Param[key] = value
	}
	return errors.New("requset is nil")
}

// Get 获取请求参数
func (r *Request) Get(key string) string {
	if r != nil && r.Param != nil {
		return r.Param[key]
	}
	return ""
}

// CalcStringToSign 计算签名字符串
func (r *Request) CalcStringToSign(httpMethod string) string {
	if r != nil && r.Param != nil {
		strslice := make([]string, len(r.Param))
		i := 0
		for k, v := range r.Param {
			data := url.Values{}
			data.Add(k, v)
			strslice[i] = data.Encode()
			strslice[i] = percentEncodeBefore(strslice[i])
			i++
		}
		sort.Strings(strslice)
		return httpMethod + "&" + percentEncode("/") + "&" + percentEncode(strings.Join(strslice, "&"))
	}
	return ""
}


//发送http请求
func (c *Client) send(r *Request) (body []byte, httpCode int, err error) {
	if r == nil || r.Param == nil {
		return nil, 0, errors.New("requset is nil")
	}

	signature := sign(c.AccessKey, r.CalcStringToSign("GET"))

	// HTTP requset
	httpReq := urllib.Get(c.EndPoint)
	if HTTPDebugEnable {
		httpReq.Debug(true)
	}
	for k, v := range r.Param {
		httpReq.Param(k, v)
	}
	httpReq.Param("Signature", signature)

	fmt.Sprint("http:%s" ,httpReq)

	resp, err := httpReq.Response()
	if err != nil {
		return nil, 0, err
	}
	if resp.Body == nil {
		return nil, resp.StatusCode, nil
	}
	defer resp.Body.Close()
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, errGzip := gzip.NewReader(resp.Body)
		if errGzip != nil {
			return nil, resp.StatusCode, errGzip
		}
		body, err = ioutil.ReadAll(reader)
	} else {
		body, err = ioutil.ReadAll(resp.Body)
	}
	if err != nil {
		return nil, resp.StatusCode, err
	}
	if HTTPDebugEnable {
		fmt.Println("C-->S:", httpReq.DumpRequestString())
		fmt.Println("S-->C:", string(body))
	}
	return body, resp.StatusCode, nil
}

// Client HTTP请求配置信息
type Client struct {
	// API版本
	Version string
	// SMS服务地域, 默认为cn-hangzhou
	Region string
	// SMS服务的地址，默认为（http://cloudpush.aliyuncs.com/）
	EndPoint string
	// 访问SMS服务的accessid，通过官方网站申请或通过管理员获取
	AccessID string
	// 访问SMS服务的accesskey，通过官方网站申请或通过管理员获取
	AccessKey string
	// 连接池中每个连接的Socket超时，单位为秒，可以为int或float。默认值为30
	SocketTimeout int

	//STS设置 start
	// 全局资源描述符 每个角色都有一个唯一的全局资源描述符，规定格式为 acs:ram::$accountID:role/$roleName
	Arn string
}

// SetVersion API版本
func (c *Client) SetVersion(version string) {
	if c != nil {
		c.Version = version
	}
}

// SetRegion 设置SMS服务地域
func (c *Client) SetRegion(region string) {
	if c != nil {
		c.Region = region
	}
}

// SetEndPoint 设置短信服务器
func (c *Client) SetEndPoint(endPoint string) {
	if c != nil {
		c.EndPoint = endPoint
	}
}

// SetAccessID 设置短信服务的accessid，通过官方网站申请或通过管理员获取
func (c *Client) SetAccessID(accessid string) {
	if c != nil {
		c.AccessID = accessid
	}
}

// SetAccessKey 设置短信服务的accesskey，通过官方网站申请或通过管理员获取
func (c *Client) SetAccessKey(accesskey string) {
	if c != nil {
		c.AccessKey = accesskey
	}
}

// SetAccessKey 设置短信服务的accesskey，通过官方网站申请或通过管理员获取
func (c *Client) SetArn(arn string) {
	if c != nil {
		c.Arn = arn
	}
}

// SetSocketTimeout 设置短信服务的Socket超时，单位为秒，可以为int或float。默认值为30
func (c *Client) SetSocketTimeout(sockettimeout int) {
	if sockettimeout == 0 {
		sockettimeout = 30
	}
	if c != nil {
		c.SocketTimeout = sockettimeout
	}
}

// setACLClient 配置默认的服务权限信息
func setACLClient(accessid, accesskey string) *Client {

	acsClient := Client{
		AccessID:accessid,
		AccessKey:accesskey,
	}

	if urllib.GetDefaultSetting().Transport == nil {
		// set default setting for urllib
		trans := &http.Transport{
			MaxIdleConnsPerHost: 500,
			Dial: (&net.Dialer{
				Timeout: time.Duration(15) * time.Second,
			}).Dial,
		}

		urlSetting := urllib.HttpSettings{
			ShowDebug:        false,            // ShowDebug
			UserAgent:        "GiterLab",       // UserAgent
			ConnectTimeout:   15 * time.Second, // ConnectTimeout
			ReadWriteTimeout: 30 * time.Second, // ReadWriteTimeout
			TlsClientConfig:  nil,              // TlsClientConfig
			Proxy:            nil,              // Proxy
			Transport:        trans,            // Transport
			EnableCookie:     false,            // EnableCookie
			Gzip:             true,             // Gzip
			DumpBody:         true,             // DumpBody
		}
		if acsClient.SocketTimeout != 0 {
			urlSetting.ConnectTimeout = time.Duration(acsClient.SocketTimeout) * time.Second
			urlSetting.ReadWriteTimeout = time.Duration(acsClient.SocketTimeout) * time.Second
		}
		if HTTPDebugEnable {
			urlSetting.ShowDebug = true
		} else {
			urlSetting.ShowDebug = false
		}
		urllib.SetDefaultSetting(urlSetting)
	}

	return &acsClient
}

func New(accessid, accesskey string) *Client {
	return setACLClient(accessid, accesskey)
}
