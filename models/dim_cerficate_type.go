package models


type DimCerficateType struct {
	Id         uint `gorm:"primary_key"`
	Code       string
	Name       string
	NameEn     string
}

// set CerficateType's table name to be `dim_cerficate_type`
func (DimCerficateType) TableName() string {
	return "dim_cerficate_type"
}

func CenficateTypeList() ([]DimCerficateType,error) {
	var cerficateType []DimCerficateType
	err := DB.Find(&cerficateType).Error
	return cerficateType, err
}
