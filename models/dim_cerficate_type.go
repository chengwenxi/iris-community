package models

type DimCerficateType struct {
	Id     uint `gorm:"primary_key"`
	Code   string
	Name   string
	NameEn string
}

func (*DimCerficateType) List() ([]DimCerficateType, error) {
	var cerficateType []DimCerficateType
	err := DB.Find(&cerficateType).Error
	return cerficateType, err
}

func (d *DimCerficateType) query(id uint) error {
	return DB.Where(&DimCerficateType{Id: id}).Find(d).Error
}

func CerficateType() (*DimCerficateType) {
	return &DimCerficateType{}
}
