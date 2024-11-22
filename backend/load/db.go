package load

import (
	"nabbr-appraisal/utils"
)

func LoadDbData(data utils.NabbrAppraisalChart, appraisalId string) (string, error) {
	if appraisalId == "" {
		appraisalId = utils.RandomAplhaNumericString(20)
	}

	// Upsert job posts to the database
	appraisalId, err := loadDbDataToMongoDb(data, appraisalId)
	if err != nil {
		return "", err
	}

	return appraisalId, nil
}
