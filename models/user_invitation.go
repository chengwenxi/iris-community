package models

import (
	"time"
)

type UserInvitation struct {
	InviteeId uint 	`gorm:"primary_key"`
	InvitationCode string
	Createtime 		time.Time
	Updatetime 		time.Time
}

func (userInvitation *UserInvitation) BeforeCreate() error {
	now := time.Now()
	userInvitation.Createtime = now
	userInvitation.Updatetime = now
	return nil
}

func (userInvitation *UserInvitation) Create() error {
	return DB.Create(userInvitation).Error
}