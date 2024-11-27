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
	responseData.AppraisalInformation.MainDivision.Head.Name = "Head"
	responseData.AppraisalInformation.MainDivision.Head.PercentageWeight = 7
	responseData.AppraisalInformation.MainDivision.Head.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
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
	responseData.AppraisalInformation.MainDivision.Face.Name = "Face"
	responseData.AppraisalInformation.MainDivision.Face.PercentageWeight = 15
	responseData.AppraisalInformation.MainDivision.Face.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
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
	responseData.AppraisalInformation.MainDivision.Neck.Name = "Neck"
	responseData.AppraisalInformation.MainDivision.Neck.PercentageWeight = 5
	responseData.AppraisalInformation.MainDivision.Neck.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
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
	responseData.AppraisalInformation.MainDivision.Forequarter.Name = "Forequarter"
	responseData.AppraisalInformation.MainDivision.Forequarter.PercentageWeight = 12
	responseData.AppraisalInformation.MainDivision.Forequarter.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
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
	responseData.AppraisalInformation.MainDivision.CenterPiece.Name = "Center Piece"
	responseData.AppraisalInformation.MainDivision.CenterPiece.PercentageWeight = 10
	responseData.AppraisalInformation.MainDivision.CenterPiece.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
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
	responseData.AppraisalInformation.MainDivision.Hindquarter.Name = "Hindquarter"
	responseData.AppraisalInformation.MainDivision.Hindquarter.PercentageWeight = 15
	responseData.AppraisalInformation.MainDivision.Hindquarter.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
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
	responseData.AppraisalInformation.MainDivision.SkinCoat.Name = "Skin/Coat"
	responseData.AppraisalInformation.MainDivision.SkinCoat.PercentageWeight = 5
	responseData.AppraisalInformation.MainDivision.SkinCoat.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
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
	responseData.AppraisalInformation.MainDivision.Health.Name = "Health"
	responseData.AppraisalInformation.MainDivision.Health.PercentageWeight = 4
	responseData.AppraisalInformation.MainDivision.Health.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
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

	responseData.AppraisalInformation.MainDivision.Temperament.Name = "Temperament"
	responseData.AppraisalInformation.MainDivision.Temperament.PercentageWeight = 8
	responseData.AppraisalInformation.MainDivision.Temperament.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
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
	responseData.AppraisalInformation.MainDivision.Movement.Name = "Movement"
	responseData.AppraisalInformation.MainDivision.Movement.PercentageWeight = 8
	responseData.AppraisalInformation.MainDivision.Movement.Characteristics = [][]utils.NabbrAppraisalChartScoreCharacteristics{
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
