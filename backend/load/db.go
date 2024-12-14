package load

import (
	"nabbr-appraisal/utils"
)

func LoadDbDataAppraisals(data utils.NabbrAppraisalChart, appraisalId string) (string, error) {
	if appraisalId == "" || appraisalId == " " {
		appraisalId = utils.RandomAplhaNumericString(20)
	}

	// Upsert job posts to the database
	appraisalId, err := loadDbDataToMongoDbAppraisals(data, appraisalId)
	if err != nil {
		return "", err
	}

	return appraisalId, nil
}

func LoadDbDataAppraisalDelete(appraisalId string) error {
	// Delete job post from the database
	err := loadDbDataToMongoDbAppraisalDelete(appraisalId)
	if err != nil {
		return err
	}
	return nil
}
