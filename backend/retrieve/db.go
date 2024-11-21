package retrieve

import (
	"log"
	"nabbr-appraisal/utils"
)

// RetrieveDbDataAll is wrapper for MongoDB and SQLite find methods.
func RetrieveDbDataAll() []utils.NabbrAppraisalChart {
	resultJobPosts, err := retrieveDbFromMongoDbAll()
	if err != nil {
		log.Print(err)
		return []utils.NabbrAppraisalChart{}
	}
	return resultJobPosts
}

func RetrieveDbDataBySearchTerm(appraisalId string) ([]utils.NabbrAppraisalChart, error) {
	resultJobPosts, err := retrieveDbFromMongoDbSearch(appraisalId)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resultJobPosts, nil
}
