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
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Face.Name = "Face"
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Face.PercentageWeight = 15
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Face.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Fusion (Skull And Mouth)",
				Score:  0,
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Lips",
				Score:  0,
				Value:  3,
				Factor: 7,
			},
			{
				Name:        "Straight And Parallel",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Nasal Bone"},
			},
			{
				Name:        "Eyes-Setting",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Eyes"},
			},
			{
				Name:        "Earflaps-Setting",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Earflaps"},
			},
		},
		{
			{
				Name:   "Stop",
				Score:  0,
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Teeth",
				Score:  0,
				Value:  3,
				Factor: 7,
			},
			{
				Name:        "Deep And Broad",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Nasal Bone"},
			},
			{
				Name:        "Eyelids",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Eyes"},
			},
			{
				Name:        "Shape",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Earflaps"},
			},
		},
		{
			{
				Name:   "Well Filled Between Eyes",
				Score:  0,
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Jaws",
				Score:  0,
				Value:  3,
				Factor: 7,
			},
			{
				Name:        "Length 8-10 Cm",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Nasal Bone"},
			},
			{
				Name:        "Colour And Pigmentation",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Eyes"},
			},
			{
				Name:        "Propotion",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Earflaps"},
			},
		},
	}
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Neck.Name = "Neck"
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Neck.PercentageWeight = 5
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Neck.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Shape",
				Score:  0,
				Value:  5,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Length",
				Score:  0,
				Value:  5,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Dewlap",
				Score:  0,
				Value:  5,
				Factor: 7,
			},
		},
	}
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Forequarter.Name = "Forequarter"
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Forequarter.PercentageWeight = 12
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Forequarter.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:        "Attachment",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Shoulders"},
			},
			{
				Name:        "Thick And Strong",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forelegs"},
			},
			{
				Name:        "Length",
				Score:       0,
				Value:       2,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Pasterns"},
			},
			{
				Name:        "Size",
				Score:       0,
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forepaws"},
			},
		},
		{
			{
				Name:        "Angulation",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Shoulders"},
			},
			{
				Name:        "Musculature",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forelegs"},
			},
			{
				Name:        "Thickness",
				Score:       0,
				Value:       2,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Pasterns"},
			},
			{
				Name:        "Shape",
				Score:       0,
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forepaws"},
			},
		},
		{
			{
				Name:        "Elbow",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Shoulders"},
			},
			{
				Name:        "Vertical",
				Score:       0,
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forelegs"},
			},
			{
				Name:        "Position",
				Score:       0,
				Value:       2,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Pasterns"},
			},
			{
				Name:        "Tread",
				Score:       0,
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forepaws"},
			},
		},
	}
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.CenterPiece.Name = "Center Piece"
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.CenterPiece.PercentageWeight = 10
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.CenterPiece.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Topline",
				Score:  0,
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Back",
				Score:  0,
				Value:  3,
				Factor: 7,
			},
			{
				Name:        "Broad",
				Score:       0,
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Chest"},
			},
		},
		{
			{
				Name:   "Loin",
				Score:  0,
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Crop",
				Score:  0,
				Value:  3,
				Factor: 7,
			},
			{
				Name:        "Deep",
				Score:       0,
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Chest"},
			},
		},
		{
			{
				Name:   "Tail",
				Score:  0,
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Musculature",
				Score:  0,
				Value:  3,
				Factor: 7,
			},
			{
				Name:        "Ribcage",
				Score:       0,
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Chest"},
			},
		},
	}
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Hindquarter.Name = "Hindquarter"
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Hindquarter.PercentageWeight = 15
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Hindquarter.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Strong And Sturdy",
				Score:  0,
				Value:  11,
				Factor: 7,
			},
			{
				Name:        "Shape And Size",
				Score:       0,
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Hind Paws"},
			},
		},
		{
			{
				Name:   "Angulation",
				Score:  0,
				Value:  11,
				Factor: 7,
			},
			{
				Name:        "Hind Pasterns, Hocks",
				Score:       0,
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Hind Paws"},
			},
		},
		{
			{
				Name:   "Stance",
				Score:  0,
				Value:  11,
				Factor: 7,
			},
			{
				Name:        "Tread",
				Score:       0,
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Hind Paws"},
			},
		},
	}
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.SkinCoat.Name = "Skin/Coat"
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.SkinCoat.PercentageWeight = 5
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.SkinCoat.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Thick And Loose Skin",
				Score:  0,
				Value:  5,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Short And Thick Hair",
				Score:  0,
				Value:  5,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Pigmentation",
				Score:  0,
				Value:  5,
				Factor: 7,
			},
		},
	}
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Health.Name = "Health"
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Health.PercentageWeight = 4
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Health.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Condition Versus Age",
				Score:  0,
				Value:  4,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Genitals",
				Score:  0,
				Value:  4,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Condition",
				Score:  0,
				Value:  4,
				Factor: 7,
			},
		},
	}

	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Temperament.Name = "Temperament"
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Temperament.PercentageWeight = 8
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Temperament.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Obedient And Manageable",
				Score:  0,
				Value:  8,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Reliable",
				Score:  0,
				Value:  8,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Self-Confident",
				Score:  0,
				Value:  8,
				Factor: 7,
			},
		},
	}
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Movement.Name = "Movement"
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Movement.PercentageWeight = 8
	responseData.NabbrAppraisalChartScoreAppraisalInformation.NabbrAppraisalChartScoreDivisions.Movement.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Buoyant",
				Score:  0,
				Value:  8,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Parallel",
				Score:  0,
				Value:  8,
				Factor: 7,
			},
		},
		{

			{
				Name:   "Topline",
				Score:  0,
				Value:  8,
				Factor: 7,
			},
		},
	}
	return responseData
}
