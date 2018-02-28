package models

type DimTokenAccessMode struct {
	Id     uint `gorm:"primary_key"`
	Code   string
	Name   string
	NameEn string
	Num    uint
}

func (token *DimTokenAccessMode) Create() error {
	return DB.Create(token).Error
}
