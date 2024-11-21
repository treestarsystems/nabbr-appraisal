package load

import (
	"nabbr-appraisal/utils"
)

func LoadDbData(data utils.NabbrAppraisalChart) error {
	appraisalId := utils.RandomAplhaNumericString(20)

	// Upsert job posts to the database
	err := loadDbDataToMongoDb(data, appraisalId)
	if err != nil {
		return err
	}

	return nil
}
