package models

import (
	"time"
	"github.com/irisnet/iris-community/utils"
)

type Users struct {
	Id         uint   `gorm:"primary_key"`
	Email      string
	Salt       string `json:"-"`
	Password   string `json:"-"`
	IsActived  bool
	IsBlocked  bool
	Createtime time.Time
	Updatetime time.Time
}

func (user *Users) BeforeCreate() error {
	now := time.Now()
	user.Createtime = now
	user.Updatetime = now
	return nil
}

func (user *Users) Create(invitationCode string) error {
	tx := DB.Begin()
	//写入user
	if err := tx.Omit("InvitationCode", "VerifyCode").Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	//写入邀请码
	selfCode := utils.IntTo52(6, int(user.Id))
	if err := tx.Omit("CountryId", "CerficateId").Create(&UserProfile{UserId: user.Id, InvitationCode: selfCode}).Error;
		err != nil {
		tx.Rollback()
		return err
	}
	//写入被邀请信息
	if invitationCode != "" {
		if err := tx.Omit().Create(&UserInvitation{InviteeId: user.Id, InvitationCode: invitationCode}).Error;
			err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (user *Users) First() error {
	return DB.First(user).Error
}

func (user *Users) ActivateUser() error {
	return DB.Model(&user).Update("IsActived", true).Error
}

func (user *Users) UpdatePwd(salt string, password string) error {
	return DB.Model(&user).Update(map[string]interface{}{"salt": salt, "password": password}).Error
}

func FindUserByEmail(email string) (Users, error) {
	var user Users
	err := DB.Where("email = ?", email).First(&user).Error
	return user, err
}
