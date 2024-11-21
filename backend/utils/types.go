package utils

import "gorm.io/gorm"

// APIResponse represents a generic API response.
type APIResponse struct {
	Status     string        `json:"status"`
	HttpStatus int           `json:"httpStatus"`
	Message    string        `json:"message"`
	Payload    []interface{} `json:"payload"`
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
