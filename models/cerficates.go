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

func (d *Cerficates) Query(id uint) (error){
	return DB.Where(&Cerficates{Id:id}).Find(d).Error
}

func NewCerficates() *Cerficates{
	return &Cerficates{}
}

