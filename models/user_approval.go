package models

import "time"

type UserApproval struct{
	UserId 			uint	`gorm:"primary_key"`
	ApprovalStatus	string
	Createtime 	time.Time
	Updatetime 	time.Time
}

func (u *UserApproval) QueryById() error{
	return DB.First(u).Error
}

func NewUserApproval(userId uint) *UserApproval{
	return &UserApproval{UserId:userId}
}