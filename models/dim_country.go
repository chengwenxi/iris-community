package models

type Country struct {
	Id         uint `gorm:"primary_key"`
	Code       string
	Name       string
	NameEn     string
}

func CountryList() ([]Country,error) {
	var country []Country
	err := DB.Find(&country).Error
	return country, err
}
