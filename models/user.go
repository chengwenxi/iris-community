package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Users struct {
	Id         uint `gorm:"primary_key"`
	Email      string
	Salt       string
	Password   string
	IsActived  bool
	IsBlocked  bool
	Createtime time.Time
	Updatetime time.Time
}

func (user *Users) BeforeCreate(scope *gorm.Scope) error {
	now := time.Now()
	user.Createtime = now
	user.Updatetime = now
	return nil
}

func (user *Users) BeforeUpdate(scope *gorm.Scope) error {
	user.Updatetime = time.Now()
	return nil
}

func (user *Users) Create() error {
	return DB.Create(user).Error
}

func (user *Users) Delete() error {
	return DB.Delete(user).Error
}

func (user *Users) First() error {
	return DB.First(user).Error
}

func (user *Users) Update() error {
	return DB.Save(user).Error
}


func AuthUser(email string, password string) (Users, error) {
	var users Users
	err := DB.Where("email = ? AND password = ?", email, password).Find(&users).Error
	return users, err
}

func UserList(skip int, limit int) ([]Users, error) {
	var users []Users
	err := DB.Limit(limit).Offset(skip).Find(&users).Error
	return users, err
}
