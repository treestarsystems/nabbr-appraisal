package utils

type NabbrAppraisalChartMemberInfo struct {
	Name string `bson:"name" json:"name" binding:"required"`
	// Email        string `bson:"email" json:"email" binding:"required"`
	MemberNumber string `bson:"memberNumber" json:"memberNumber" binding:"required"`
}

type NabbrAppraisalChartPetInfo struct {
	Name               string `bson:"name" json:"name" binding:"required"`
	Age                int    `bson:"age" json:"age" binding:"required"`
	DnaNumber          string `bson:"dnaNumber" json:"dnaNumber" binding:"required"`
	Weight             int    `bson:"weight" json:"weight" binding:"required"`
	Color              string `bson:"color" json:"color" binding:"required"`
	Markings           string `bson:"markings" json:"markings" binding:"required"`
	Microchip          string `bson:"microchip" json:"microchip" binding:"required"`
	Sex                string `bson:"sex" json:"sex" binding:"required"`
	RegistrationNumber string `bson:"registrationNumber" json:"registrationNumber" binding:"required"`
}

type NabbrAppraisalChartScoreCharacteristicsSubDivision struct {
	Name string `bson:"name" json:"name"`
}
type NabbrAppraisalChartScoreCharacteristics struct {
	Name        string                                             `bson:"name" json:"name" binding:"required"`
	Score       int                                                `bson:"score" json:"score" binding:"required"`
	Value       int                                                `bson:"value" json:"value" binding:"required"`
	Factor      int                                                `bson:"factor" json:"factor" binding:"required"`
	SubDivision NabbrAppraisalChartScoreCharacteristicsSubDivision `bson:"subDivision" json:"subDivision,omitempty"`
}

type NabbrAppraisalChartScoreDivision struct {
	Characteristics  [][]NabbrAppraisalChartScoreCharacteristics `bson:"characteristics" json:"characteristics" binding:"required"`
	Name             string                                      `bson:"name" json:"name" binding:"required"`
	PercentageWeight int                                         `bson:"percentageWeight" json:"percentageWeight" binding:"required"`
}

type NabbrAppraisalChartScoreDivisions struct {
	Appearance  NabbrAppraisalChartScoreDivision `bson:"appearance" json:"appearance" binding:"required"`
	Head        NabbrAppraisalChartScoreDivision `bson:"head" json:"head" binding:"required"`
	Face        NabbrAppraisalChartScoreDivision `bson:"face" json:"face" binding:"required"`
	Neck        NabbrAppraisalChartScoreDivision `bson:"neck" json:"neck" binding:"required"`
	Forequarter NabbrAppraisalChartScoreDivision `bson:"forequarter" json:"forequarter" binding:"required"`
	CenterPiece NabbrAppraisalChartScoreDivision `bson:"centerPiece" json:"centerPiece" binding:"required"`
	Hindquarter NabbrAppraisalChartScoreDivision `bson:"hindquarter" json:"hindquarter" binding:"required"`
	SkinCoat    NabbrAppraisalChartScoreDivision `bson:"skinCoat" json:"skinCoat" binding:"required"`
	Health      NabbrAppraisalChartScoreDivision `bson:"health" json:"health" binding:"required"`
	Temperament NabbrAppraisalChartScoreDivision `bson:"temperament" json:"temperament" binding:"required"`
	Movement    NabbrAppraisalChartScoreDivision `bson:"movement" json:"movement" binding:"required"`
}

type NabbrAppraisalChartScoreAppraisalInformation struct {
	MainDivision          NabbrAppraisalChartScoreDivisions `bson:"mainDivision" json:"mainDivision" binding:"required"`
	AppraiserName         string                            `bson:"appraiserName" json:"appraiserName" binding:"required"`
	AppraiserNumber       string                            `bson:"appraiserNumber" json:"appraiserNumber" binding:"required"`
	SeniorAppraiserName   string                            `bson:"seniorAppraiserName" json:"seniorAppraiserName" binding:"required"`
	SeniorAppraiserNumber string                            `bson:"seniorAppraiserNumber" json:"seniorAppraiserNumber" binding:"required"`
	Date                  string                            `bson:"date" json:"date" binding:"required"`
	AdditionalComments    string                            `bson:"additionalComments" json:"additionalComments"`
	Place                 string                            `bson:"place" json:"place"`
	AppraisalScore        float32                           `bson:"appraisalScore" json:"appraisalScore" binding:"required"`
}

type NabbrAppraisalChart struct {
	AppraisalId          string                                       `bson:"appraisalId" json:"appraisalId"`
	MemberInformation    NabbrAppraisalChartMemberInfo                `bson:"memberInformation" json:"memberInformation" binding:"required"`
	PetInformation       NabbrAppraisalChartPetInfo                   `bson:"petInformation" json:"petInformation" binding:"required"`
	AppraisalInformation NabbrAppraisalChartScoreAppraisalInformation `bson:"appraisalInformation" json:"appraisalInformation" binding:"required"`
}
