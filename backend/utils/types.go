package utils

import "gorm.io/gorm"

// APIResponse represents a generic API response.
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoadDbInsertGorm struct {
	ID        uint           `gorm:"primarykey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type NabbrAppraisalChart struct {
	NabbrAppraisalChartMemberInfo                `json:"memberInformation"`
	NabbrAppraisalChartPetInfo                   `json:"petInformation"`
	NabbrAppraisalChartScoreAppraisalInformation `json:"appraisalInformation"`
}
