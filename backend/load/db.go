package load

import (
	"nabbr-appraisal/utils"
	"os"
)

func LoadDbData(data utils.NabbrAppraisalChart) error {
	appraisalId := utils.RandomAplhaNumericString(20)

	// Upsert job posts to the database
	if os.Getenv("DB_MONGODB_ENABLE") == "true" {
		err := loadDbDataToMongoDb(data, appraisalId)
		if err != nil {
			return err
		}
	}

	return nil
}
