package models

import "time"

type UserProfile struct{
	UserId         	uint 	`gorm:"primary_key"`
	FamilyName    	string
	Name         	string
	CountryId      	uint
	CerficateId    	uint
	InvitationCode 	string
	Createtime 		time.Time
	Updatetime 		time.Time
}

func (userProfile *UserProfile) BeforeCreate() error {
	now := time.Now()
	userProfile.Createtime = now
	userProfile.Updatetime = now
	return nil
}

func (userProfile *UserProfile) Create() error {
	return DB.Create(userProfile).Error
}

func (userProfile *UserProfile) First() error {
	return DB.First(userProfile).Error
}
