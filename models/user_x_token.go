package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type UserXToken struct {
	Id           uint `gorm:"primary_key"`
	UserId       uint
	AccessModeID uint
	Createtime   time.Time
	Updatetime   time.Time
}

func (user *UserXToken) BeforeCreate(scope *gorm.Scope) error {
	now := time.Now()
	user.Createtime = now
	user.Updatetime = now
	return nil
}

func (user *UserXToken) Create() error {
	return DB.Create(user).Error
}
