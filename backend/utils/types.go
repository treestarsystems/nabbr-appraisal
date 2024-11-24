package utils

// APIResponse represents a generic API response.
type APIResponse struct {
	Status     string        `json:"status"`
	HttpStatus int           `json:"httpStatus"`
	Message    string        `json:"message"`
	Payload    []interface{} `json:"payload"`
}

type NabbrAppraisalChart struct {
	AppraisalId          string                                       `bson:"appraisalId" json:"appraisalId"`
	MemberInformation    NabbrAppraisalChartMemberInfo                `bson:"memberInformation" json:"memberInformation" binding:"required"`
	PetInformation       NabbrAppraisalChartPetInfo                   `bson:"petInformation" json:"petInformation" binding:"required"`
	AppraisalInformation NabbrAppraisalChartScoreAppraisalInformation `bson:"appraisalInformation" json:"appraisalInformation" binding:"required"`
}
