package models

import "time"

type UserProfile struct{
	UserId         	uint
	family_name    	string
	name         	string
	country_id      uint
	cerficate_id    uint
	invitation_code string
	Createtime 		time.Time
	Updatetime 		time.Time
}

func (userProfile *UserProfile) First() error {
	return DB.First(userProfile).Error
}
