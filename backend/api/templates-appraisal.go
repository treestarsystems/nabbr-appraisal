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
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Balance",
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Musculature",
				Value:  4,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Gender Authenticity",
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Impressive",
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Bearing",
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
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Broad",
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Deep",
				Value:  4,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Square",
				Value:  3,
				Factor: 7,
			},

			{
				Name:   "Filled",
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Characteristics Typically Boerboel",
				Value:  3,
				Factor: 7,
			},
		},
		{
			{
				Name:        "Large in circumference",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Skull"},
			},
			{
				Name:        "Flat",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Skull"},
			},

			{
				Name:        "Muscular",
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
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Stop",
				Value:  3,
				Factor: 7,
			}, {
				Name:   "Well Filled Between Eyes",
				Value:  3,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Lips",
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Teeth",
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Jaws",
				Value:  3,
				Factor: 7,
			},
		},
		{
			{
				Name:        "Straight And Parallel",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Nasal Bone"},
			},
			{
				Name:        "Deep And Broad",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Nasal Bone"},
			},
			{
				Name:        "Length 8-10 Cm",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Nasal Bone"},
			},
		},
		{
			{
				Name:        "Eyes-Setting",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Eyes"},
			},
			{
				Name:        "Eyelids",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Eyes"},
			},
			{
				Name:        "Colour And Pigmentation",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Eyes"},
			},
		},
		{
			{
				Name:        "Earflaps-Setting",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Earflaps"},
			},
			{
				Name:        "Shape",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Earflaps"},
			},
			{
				Name:        "Propotion",
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
				Value:  5,
				Factor: 7,
			},
			{
				Name:   "Length",
				Value:  5,
				Factor: 7,
			},
			{
				Name:   "Dewlap",
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
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Shoulders"},
			},
			{
				Name:        "Angulation",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Shoulders"},
			},
			{
				Name:        "Elbow",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Shoulders"},
			},
		},
		{
			{
				Name:        "Thick And Strong",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forelegs"},
			},
			{
				Name:        "Musculature",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forelegs"},
			},
			{
				Name:        "Vertical",
				Value:       3,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forelegs"},
			},
		}, {
			{
				Name:        "Length",
				Value:       2,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Pasterns"},
			},
			{
				Name:        "Thickness",
				Value:       2,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Pasterns"},
			},
			{
				Name:        "Position",
				Value:       2,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Pasterns"},
			},
		},
		{
			{
				Name:        "Size",
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forepaws"},
			},
			{
				Name:        "Shape",
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Forepaws"},
			},
			{
				Name:        "Tread",
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
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Loin",
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Tail",
				Value:  3,
				Factor: 7,
			},
		},
		{
			{
				Name:   "Back",
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Croup",
				Value:  3,
				Factor: 7,
			},
			{
				Name:   "Musculature",
				Value:  3,
				Factor: 7,
			},
		},
		{
			{
				Name:        "Broad",
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Chest"},
			},
			{
				Name:        "Deep",
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Chest"},
			},
			{
				Name:        "Ribcage",
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
				Value:  11,
				Factor: 7,
			},
			{
				Name:   "Angulation",
				Value:  11,
				Factor: 7,
			},
			{
				Name:   "Stance",
				Value:  11,
				Factor: 7,
			},
		},
		{
			{
				Name:        "Shape And Size",
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Hind Paws"},
			},
			{
				Name:        "Hind Pasterns, Hocks",
				Value:       4,
				Factor:      7,
				SubDivision: utils.NabbrAppraisalChartScoreCharacteristicsSubDivision{Name: "Hind Paws"},
			},
			{
				Name:        "Tread",
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
				Value:  5,
				Factor: 7,
			},
			{
				Name:   "Short And Thick Hair",
				Value:  5,
				Factor: 7,
			},
			{
				Name:   "Pigmentation",
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
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Genitals",
				Value:  4,
				Factor: 7,
			},
			{
				Name:   "Condition",
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
				Value:  8,
				Factor: 7,
			},
			{
				Name:   "Reliable",
				Value:  8,
				Factor: 7,
			},
			{
				Name:   "Self-Confident",
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
				Value:  8,
				Factor: 7,
			},
			{
				Name:   "Parallel",
				Value:  8,
				Factor: 7,
			},
			{
				Name:   "Topline",
				Value:  8,
				Factor: 7,
			},
		},
	}
	return responseData
}
