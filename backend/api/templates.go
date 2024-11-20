package api

import "nabbr-appraisal/utils"

func GenerateAppraisalEmptyChart() utils.NabbrAppraisalChart {
	var responseData utils.NabbrAppraisalChart

	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Appearance.Name = "General Appearance"
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Appearance.PercentageWeight = 8
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Appearance.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Big And Strong",
				Score:  0,
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Gender Authenticity",
				Score:  0,
				Value:  4,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Balance",
				Score:  0,
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Impressive",
				Score:  0,
				Value:  4,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Musculature",
				Score:  0,
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Bearing",
				Score:  0,
				Value:  4,
				Factor: 7,
			},
		},
	}
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Head.Name = "Head"
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Head.PercentageWeight = 7
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Head.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Short",
				Score:  0,
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Square",
				Score:  0,
				Value:  3,
				Factor: 7,
			},
			{
				Name:        "Large in circumference",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Skull"},
			},
		},
		{
			{
				Name:   "Broad",
				Score:  0,
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Filled",
				Score:  0,
				Value:  3,
				Factor: 7,
			},
			{
				Name:        "Flat",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Skull"},
			},
		},
		{
			{
				Name:   "Deep",
				Score:  0,
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Characteristics Typically Boerboel",
				Score:  0,
				Value:  3,
				Factor: 7,
			},
			{
				Name:        "Muscular",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Skull"},
			},
		},
	}
	return responseData
}
