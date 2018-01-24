package models


type CerficateType struct {
	Id         uint `gorm:"primary_key"`
	Code       string
	Name       string
	NameEn     string
}

func CenficateTypeList() ([]CerficateType,error) {
	var cerficateType []CerficateType
	err := DB.Find(&cerficateType).Error
	return cerficateType, err
}
