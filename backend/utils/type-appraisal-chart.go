package utils

type NabbrAppraisalChartMemberInfo struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	MemberNumber string `json:"memberNumber"`
}

type NabbrAppraisalChartPetInfo struct {
	Name               string `json:"name"`
	Type               string `json:"type"`
	Breed              string `json:"breed"`
	Age                string `json:"age"`
	DnaNumber          string `json:"dnaNumber"`
	Weight             string `json:"weight"`
	Color              string `json:"color"`
	Markings           string `json:"markings"`
	Microchip          string `json:"microchip"`
	Sex                string `json:"sex"`
	RegistrationNumber string `json:"registrationNumber"`
}

type NabbrAppraisalChartScoreCharacteristicsSubDivision struct {
	Name string `json:"name"`
}
type NabbrAppraisalChartScoreCharacteristics struct {
	Name        string                                             `json:"name"`
	Score       int                                                `json:"score"`
	Value       int                                                `json:"value"`
	Factor      int                                                `json:"factor"`
	SubDivision NabbrAppraisalChartScoreCharacteristicsSubDivision `json:"subDivision,omitempty"`
}

type NabbrAppraisalChartScoreDivision struct {
	Characteristics  [][]NabbrAppraisalChartScoreCharacteristics `json:"characteristics"`
	Name             string                                      `json:"name"`
	PercentageWeight int                                         `json:"percentageWeight"`
}

type NabbrAppraisalChartScoreDivisions struct {
	Appearance struct {
		NabbrAppraisalChartScoreDivision
	} `json:"appearance"`
	Head struct {
		NabbrAppraisalChartScoreDivision
	} `json:"head"`
	Face struct {
		NabbrAppraisalChartScoreDivision
	} `json:"face"`
	Neck struct {
		NabbrAppraisalChartScoreDivision
	} `json:"neck"`
	Forequarter struct {
		NabbrAppraisalChartScoreDivision
	} `json:"forequarter"`
	CenterPiece struct {
		NabbrAppraisalChartScoreDivision
	} `json:"centerPiece"`
	Hindquarter struct {
		NabbrAppraisalChartScoreDivision
	} `json:"hindquarter"`
	SkinCoat struct {
		NabbrAppraisalChartScoreDivision
	} `json:"skinCoat"`
	Health struct {
		NabbrAppraisalChartScoreDivision
	} `json:"health"`
	Temperament struct {
		NabbrAppraisalChartScoreDivision
	} `json:"temperament"`
	Movement struct {
		NabbrAppraisalChartScoreDivision
	} `json:"movement"`
}

type NabbrAppraisalChartScoreAppraisalInformation struct {
	NabbrAppraisalChartScoreDivisions `json:"mainDivision"`
	Appraiser                         string `json:"appraiser"`
	AppraiserNumber                   string `json:"appraiserNumber"`
	SeniorAppraiser                   string `json:"seniorAppraiser"`
	SeniorAppraiserNumber             string `json:"seniorAppraiserNumber"`
	Date                              string `json:"date"`
	Value                             string `json:"value"`
	Notes                             string `json:"notes"`
	AdditionalComments                string `json:"additionalComments"`
	Place                             string `json:"place"`
}
