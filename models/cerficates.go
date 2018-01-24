package models

type Cerficates struct{
	Id         	uint `gorm:"primary_key"`
	TypeId	   	uint
	Num        	string
	FrontFileId uint
	ReverseFileId uint
	HandFileId uint
}

func (cerficates *Cerficates) Create() error {
	return DB.Create(cerficates).Error
}

