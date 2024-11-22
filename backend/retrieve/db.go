package retrieve

import (
	"log"
	"nabbr-appraisal/utils"
)

// RetrieveDbDataAll is wrapper for MongoDB and SQLite find methods.
func RetrieveDbDataAll() []utils.NabbrAppraisalChart {
	resultAppraisalCharts, err := retrieveDbFromMongoDbAll()
	if err != nil {
		log.Print(err)
		return []utils.NabbrAppraisalChart{}
	}
	return resultAppraisalCharts
}

func RetrieveDbDataBySearchTerm(appraisalId string) ([]utils.NabbrAppraisalChart, error) {
	resultAppraisalCharts, err := retrieveDbFromMongoDbSearch(appraisalId)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resultAppraisalCharts, nil
}
