package models

type DimCountry struct {
	Id         uint `gorm:"primary_key"`
	Code       string
	Name       string
	NameEn     string
}

func CountryList() ([]DimCountry,error) {
	var country []DimCountry
	err := DB.Find(&country).Error
	return country, err
}

func Country(){

}
