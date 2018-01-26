package models

type DimCountry struct {
	Id         uint `gorm:"primary_key"`
	Code       string
	Name       string
	NameEn     string
}

func (*DimCountry)List() ([]DimCountry,error) {
	var countrys []DimCountry
	err := DB.Find(&countrys).Error
	return countrys, err
}

func Country()(*DimCountry){
	return &DimCountry{}
}
