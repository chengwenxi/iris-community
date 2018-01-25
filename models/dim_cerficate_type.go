package models


type DimCerficateType struct {
	Id         uint `gorm:"primary_key"`
	Code       string
	Name       string
	NameEn     string
}

func (DimCerficateType) List() ([]DimCerficateType,error) {
	var cerficateType []DimCerficateType
	err := DB.Find(&cerficateType).Error
	return cerficateType, err
}

func CerficateType()(DimCerficateType)  {
	return DimCerficateType{}
}