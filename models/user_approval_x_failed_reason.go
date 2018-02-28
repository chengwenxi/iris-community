package models

import "time"

type UserApprovalXFailedReason struct {
	Id         uint `gorm:"primary_key"`
	UserId     uint
	ReasonId   uint
	Createtime time.Time
	Updatetime time.Time
}

func (R *UserApprovalXFailedReason) QueryByUserId(userId uint) ([]UserApprovalXFailedReason, error) {
	var reasons []UserApprovalXFailedReason
	err := DB.Where(&UserApprovalXFailedReason{UserId: userId}).Find(&reasons).Error
	return reasons, err
}

func NewUserApprovalXFailedReason() *UserApprovalXFailedReason {
	return &UserApprovalXFailedReason{}
}
