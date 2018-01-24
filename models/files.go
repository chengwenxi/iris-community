package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Files struct {
	Id         uint `gorm:"primary_key"`
	OssKey     string
	Createtime time.Time
	Updatetime time.Time
}

func (file *Files) BeforeCreate(scope *gorm.Scope) error {
	now := time.Now()
	file.Createtime = now
	file.Updatetime = now
	return nil
}

func (file *Files) Create() error {
	return DB.Create(file).Error
}

