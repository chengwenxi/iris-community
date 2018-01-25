package models

import (
	"time"
	"github.com/irisnet/iris-community/utils"
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

func (user *Users) BeforeCreate() error {
	now := time.Now()
	user.Createtime = now
	user.Updatetime = now
	return nil
}

func (user *Users) Create() error {
	tx := DB.Begin()
	//写入user
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	//写入邀请码
	invitationCode := utils.IntTo52(6, int(user.Id))
	if err := tx.Omit("CountryId", "CerficateId").Create(&UserProfile{UserId: user.Id, InvitationCode: invitationCode}).Error;
		err != nil {
		tx.Rollback()
		return err
	}
	//写入被邀请信息
	if err := tx.Omit().Create(&UserInvitation{InviteeId: user.Id, InvitationCode: invitationCode}).Error;
		err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (user *Users) Delete() error {
	return DB.Delete(user).Error
}

func (user *Users) First() error {
	return DB.First(user).Error
}

func (user *Users) Update() error {
	return DB.Omit("Createtime", "Updatetime").Updates(user).Error
}

func FindUserByEmail(email string) (Users, error) {
	var user Users
	err := DB.Where("Email = ?", email).First(&user).Error
	return user, err
}

func UserList(skip int, limit int) ([]Users, error) {
	var users []Users
	err := DB.Limit(limit).Offset(skip).Find(&users).Error
	return users, err
}
