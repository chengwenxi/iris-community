package models

type DimApprovalFailedReason struct {
	Id     uint `gorm:"primary_key"`
	Code   string
	Name   string
	NameCn string
}

func (d *DimApprovalFailedReason) QueryByUserId(id uint) (DimApprovalFailedReason, error) {
	var df DimApprovalFailedReason
	err := DB.Where(&DimApprovalFailedReason{Id: id}).First(&df).Error
	return df, err
}

func NewReason() *DimApprovalFailedReason {
	return &DimApprovalFailedReason{}
}
