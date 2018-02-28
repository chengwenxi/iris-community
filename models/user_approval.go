package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type UserApproval struct {
	UserId         uint `gorm:"primary_key"`
	ApprovalStatus string
	Createtime     time.Time
	Updatetime     time.Time
}

func (u *UserApproval) BeforeCreate(scope *gorm.Scope) error {
	now := time.Now()
	u.Createtime = now
	u.Updatetime = now
	return nil
}

func (u *UserApproval) QueryById() error {
	return DB.First(u).Error
}

func (u *UserApproval) Create() error {
	return DB.Create(u).Error
}

func NewUserApproval(userId uint) *UserApproval {
	return &UserApproval{UserId: userId}
}
