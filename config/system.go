package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

type Configuration struct {
	Postgres        *Postgres `yaml:"postgres"`
	Server          string    `yaml:"server"`
	Mail            *Mail     `yaml:"mail"`
	StaticPath      string    `yaml:"staticPath"`
	AliYun      	*AliYun    `yaml:"aliYun"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Post     int    `yaml:"post"`
	Dbname   string `yaml:"dbname"`
	Password string `yaml:"password"`
	Sslmode  string `yaml:"sslmode"`
}

type Mail struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type AliYun struct{
	Endpoint string `yaml:"endpoint"`
	AccessKeyId string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	Bucket string `yaml:"bucket"`
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
