package api

import "nabbr-appraisal/utils"

func NewAppraisalChart() utils.NabbrAppraisalChart {
	var responseData utils.NabbrAppraisalChart

	responseData.AppraisalInformation.MainDivision.Appearance.Name = "General Appearance"
	responseData.AppraisalInformation.MainDivision.Appearance.PercentageWeight = 8
	responseData.AppraisalInformation.MainDivision.Appearance.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Big And Strong",
				Score:  "nil",
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Balance",
				Score:  "nil",
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Musculature",
				Score:  "nil",
				Value:  4,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Gender Authenticity",
				Score:  "nil",
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Impressive",
				Score:  "nil",
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Bearing",
				Score:  "nil",
				Value:  4,
				Factor: 7,
			},
		},
	}
	responseData.AppraisalInformation.MainDivision.Head.Name = "Head"
	responseData.AppraisalInformation.MainDivision.Head.PercentageWeight = 7
	responseData.AppraisalInformation.MainDivision.Head.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Short",
				Score:  "nil",
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Broad",
				Score:  "nil",
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Deep",
				Score:  "nil",
				Value:  4,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Square",
				Score:  "nil",
				Value:  3,
				Factor: 7,
			},

			{
				Name:   "Filled",
				Score:  "nil",
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Characteristics Typically Boerboel",
				Score:  "nil",
				Value:  3,
				Factor: 7,
			},
		},
		{
			{
				Name:        "Large in circumference",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Skull"},
			},
			{
				Name:        "Flat",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Skull"},
			},

			{
				Name:        "Muscular",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Skull"},
			},
		},
	}
	responseData.AppraisalInformation.MainDivision.Face.Name = "Face"
	responseData.AppraisalInformation.MainDivision.Face.PercentageWeight = 15
	responseData.AppraisalInformation.MainDivision.Face.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Fusion (Skull And Mouth)",
				Score:  "nil",
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Stop",
				Score:  "nil",
				Value:  3,
				Factor: 7,
			}, {
				Name:   "Well Filled Between Eyes",
				Score:  "nil",
				Value:  3,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Lips",
				Score:  "nil",
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Teeth",
				Score:  "nil",
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Jaws",
				Score:  "nil",
				Value:  3,
				Factor: 7,
			},
		},
		{
			{
				Name:        "Straight And Parallel",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Nasal Bone"},
			},
			{
				Name:        "Deep And Broad",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Nasal Bone"},
			},
			{
				Name:        "Length 8-10 Cm",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Nasal Bone"},
			},
		},
		{
			{
				Name:        "Eyes-Setting",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Eyes"},
			},
			{
				Name:        "Eyelids",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Eyes"},
			},
			{
				Name:        "Colour And Pigmentation",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Eyes"},
			},
		},
		{
			{
				Name:        "Earflaps-Setting",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Earflaps"},
			},
			{
				Name:        "Shape",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Earflaps"},
			},
			{
				Name:        "Propotion",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Earflaps"},
			},
		},
	}
	responseData.AppraisalInformation.MainDivision.Neck.Name = "Neck"
	responseData.AppraisalInformation.MainDivision.Neck.PercentageWeight = 5
	responseData.AppraisalInformation.MainDivision.Neck.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Shape",
				Score:  "nil",
				Value:  5,
				Factor: 7,
			},
			{
				Name:   "Length",
				Score:  "nil",
				Value:  5,
				Factor: 7,
			},
			{
				Name:   "Dewlap",
				Score:  "nil",
				Value:  5,
				Factor: 7,
			},
		},
	}
	responseData.AppraisalInformation.MainDivision.Forequarter.Name = "Forequarter"
	responseData.AppraisalInformation.MainDivision.Forequarter.PercentageWeight = 12
	responseData.AppraisalInformation.MainDivision.Forequarter.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:        "Attachment",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Shoulders"},
			},
			{
				Name:        "Angulation",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Shoulders"},
			},
			{
				Name:        "Elbow",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Shoulders"},
			},
		},
		{
			{
				Name:        "Thick And Strong",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forelegs"},
			},
			{
				Name:        "Musculature",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forelegs"},
			},
			{
				Name:        "Vertical",
				Score:       "nil",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forelegs"},
			},
		}, {
			{
				Name:        "Length",
				Score:       "nil",
				Value:       2,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Pasterns"},
			},
			{
				Name:        "Thickness",
				Score:       "nil",
				Value:       2,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Pasterns"},
			},
			{
				Name:        "Position",
				Score:       "nil",
				Value:       2,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Pasterns"},
			},
		},
		{
			{
				Name:        "Size",
				Score:       "nil",
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forepaws"},
			},
			{
				Name:        "Shape",
				Score:       "nil",
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forepaws"},
			},
			{
				Name:        "Tread",
				Score:       "nil",
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forepaws"},
			},
		},
	}
	responseData.AppraisalInformation.MainDivision.CenterPiece.Name = "Center Piece"
	responseData.AppraisalInformation.MainDivision.CenterPiece.PercentageWeight = 10
	responseData.AppraisalInformation.MainDivision.CenterPiece.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Topline",
				Score:  "nil",
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Loin",
				Score:  "nil",
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Tail",
				Score:  "nil",
				Value:  3,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Back",
				Score:  "nil",
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Croup",
				Score:  "nil",
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Musculature",
				Score:  "nil",
				Value:  3,
				Factor: 7,
			},
		},
		{
			{
				Name:        "Broad",
				Score:       "nil",
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Chest"},
			},
			{
				Name:        "Deep",
				Score:       "nil",
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Chest"},
			},
			{
				Name:        "Ribcage",
				Score:       "nil",
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Chest"},
			},
		},
	}
	responseData.AppraisalInformation.MainDivision.Hindquarter.Name = "Hindquarter"
	responseData.AppraisalInformation.MainDivision.Hindquarter.PercentageWeight = 15
	responseData.AppraisalInformation.MainDivision.Hindquarter.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Strong And Sturdy",
				Score:  "nil",
				Value:  11,
				Factor: 7,
			},
			{
				Name:   "Angulation",
				Score:  "nil",
				Value:  11,
				Factor: 7,
			},
			{
				Name:   "Stance",
				Score:  "nil",
				Value:  11,
				Factor: 7,
			},
		},
		{
			{
				Name:        "Shape And Size",
				Score:       "nil",
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Hind Paws"},
			},
			{
				Name:        "Hind Pasterns, Hocks",
				Score:       "nil",
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Hind Paws"},
			},
			{
				Name:        "Tread",
				Score:       "nil",
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Hind Paws"},
			},
		},
	}
	responseData.AppraisalInformation.MainDivision.SkinCoat.Name = "Skin/Coat"
	responseData.AppraisalInformation.MainDivision.SkinCoat.PercentageWeight = 5
	responseData.AppraisalInformation.MainDivision.SkinCoat.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Thick And Loose Skin",
				Score:  "nil",
				Value:  5,
				Factor: 7,
			},
			{
				Name:   "Short And Thick Hair",
				Score:  "nil",
				Value:  5,
				Factor: 7,
			},
			{
				Name:   "Pigmentation",
				Score:  "nil",
				Value:  5,
				Factor: 7,
			},
		},
	}
	responseData.AppraisalInformation.MainDivision.Health.Name = "Health"
	responseData.AppraisalInformation.MainDivision.Health.PercentageWeight = 4
	responseData.AppraisalInformation.MainDivision.Health.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Condition Versus Age",
				Score:  "nil",
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Genitals",
				Score:  "nil",
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Condition",
				Score:  "nil",
				Value:  4,
				Factor: 7,
			},
		},
	}

	responseData.AppraisalInformation.MainDivision.Temperament.Name = "Temperament"
	responseData.AppraisalInformation.MainDivision.Temperament.PercentageWeight = 8
	responseData.AppraisalInformation.MainDivision.Temperament.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Obedient And Manageable",
				Score:  "nil",
				Value:  8,
				Factor: 7,
			},
			{
				Name:   "Reliable",
				Score:  "nil",
				Value:  8,
				Factor: 7,
			},
			{
				Name:   "Self-Confident",
				Score:  "nil",
				Value:  8,
				Factor: 7,
			},
		},
	}
	responseData.AppraisalInformation.MainDivision.Movement.Name = "Movement"
	responseData.AppraisalInformation.MainDivision.Movement.PercentageWeight = 8
	responseData.AppraisalInformation.MainDivision.Movement.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
		{
			{
				Name:   "Buoyant",
				Score:  "nil",
				Value:  8,
				Factor: 7,
			},
			{
				Name:   "Parallel",
				Score:  "nil",
				Value:  8,
				Factor: 7,
			},
			{
				Name:   "Topline",
				Score:  "nil",
				Value:  8,
				Factor: 7,
			},
		},
	}
	return responseData
}
