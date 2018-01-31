package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Files struct {
	Id         uint `gorm:"primary_key"`
	OssKey     string	`gorm:"column:oos_key"`
	Createtime time.Time
	Updatetime time.Time
}

func (f *Files) BeforeCreate(scope *gorm.Scope) error {
	now := time.Now()
	f.Createtime = now
	f.Updatetime = now
	return nil
}

func (f *Files) Create() error {
	return DB.Create(f).Error
}

func (f *Files) BatchQuery(ids []uint)([] Files,error){
	var files []Files
	err :=DB.Where(ids).Find(&files).Error
	return files,err
}

func (f *Files) QueryById(id uint)(Files,error){
	file := Files{Id:id}
	err := DB.First(&file).Error
	return file,err
}

func NewFiles() *Files{
	return &Files{}
}
