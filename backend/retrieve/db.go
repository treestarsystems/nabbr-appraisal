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

// RetrieveDbDataSearch is wrapper for MongoDB and SQLite search methods that retrieves job posts based on the search term.
// func RetrieveDbDataSearch(appraisalId string) []utils.NabbrAppraisalChart {

// resultJobPosts, err := retrieveDbFromMongoDbQuery(appraisalId)
// if err != nil {
// 	log.Print(err)
// 	return []utils.NabbrAppraisalChart{}
// }
// return resultJobPosts
// if os.Getenv("DB_MONGODB_ENABLE") == "true" {
// }

// if os.Getenv("DB_SQLITE_ENABLE") == "true" {
// 	resultJobPosts, err := retrieveDbFromSqliteSearch(searchTerm)
// 	if err != nil {
// 		log.Print(err)
// 		return []utils.JobPost{}
// 	}
// 	return resultJobPosts
// }
// return []utils.JobPost{}
// }

func RetrieveDbDataByAppraisalId(appraisalId string) ([]utils.NabbrAppraisalChart, error) {

	resultJobPosts, err := retrieveDbFromMongoDbSearch(appraisalId)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resultJobPosts, nil
}
