package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

type Configuration struct {
	Postgres   *Postgres `yaml:"postgres"`
	Redis      *Redis    `yaml:"redis"`
	Server     string    `yaml:"server"`
	Mail       *Mail     `yaml:"mail"`
	StaticPath string    `yaml:"staticPath"`
	AliYun     *AliYun   `yaml:"aliYun"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Post     int    `yaml:"post"`
	Dbname   string `yaml:"dbname"`
	Password string `yaml:"password"`
	Sslmode  string `yaml:"sslmode"`
}

type Redis struct {
	Url         string `yaml:"host"`
	ActcTimeout int    `yaml:"actcTimeout"`
	RescTimeout int    `yaml:"rescTimeout"`
	VercTimeOut int    `yaml:"vercTimeOut"`
}

type Mail struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type AliYun struct {
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	Oss             *Oss   `yaml:"oss"`
	Sls             *Sls   `yaml:"sls"`
}
type Oss struct {
	Endpoint string `yaml:"endpoint"`
	Bucket   string `yaml:"bucket"`
}
type Sls struct {
	Endpoint string `yaml:"endpoint"`
	Arn      string `yaml:"arn"`
}

var Config *Configuration

func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	//var config Configuration
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	//Config = &config
	log.Printf("config load succeessfully:%v", Config)
	return nil
}
