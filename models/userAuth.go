package models

import (
	"time"
)

type UserAuth struct {
	Id         uint `gorm:"primary_key"`
	UserId     uint
	AuthCode   string
	ExpiresIn  string
	Createtime time.Time
	Updatetime time.Time
}

func (userAuth *UserAuth) Create() error {
	if err := DB.Omit("AuthCode", "ExpiresIn", "Createtime", "Updatetime").Create(userAuth).Error; err != nil {
		return err
	}
	return DB.First(userAuth).Error
}

func (userAuth *UserAuth) First() error {
	return DB.First(userAuth).Error
}

func (userAuth *UserAuth) FindByAuth() error{
	return DB.Where("AuthCode = ?", userAuth.AuthCode).First(userAuth).Error
}
