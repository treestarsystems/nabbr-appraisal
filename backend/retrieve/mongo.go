package retrieve

import (
	"context"
	"fmt"
	"nabbr-appraisal/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// retrieveDbFromMongoDbQuery performs a query on the MongoDB collection and returns the results.
func retrieveDbFromMongoDbQuery(filter interface{}) ([]utils.NabbrAppraisalChart, error) {
	var appraisalCharts []utils.NabbrAppraisalChart

	// Set a context with a timeout for the query
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var cursor *mongo.Cursor
	var err error

	// Perform the search query based on the type of filter
	switch f := filter.(type) {
	case bson.M:
		cursor, err = utils.CollectionMongo.Find(ctx, f)
	case mongo.Pipeline:
		cursor, err = utils.CollectionMongo.Aggregate(ctx, f)
	default:
		return nil, fmt.Errorf("error - MongoDB: Unsupported filter type: %T", filter)
	}

	if err != nil {
		return nil, fmt.Errorf("error - MongoDB: Unable to perform search query: %w", err)
	}
	defer cursor.Close(ctx)

	// Iterate through the cursor and decode each document into a AppraisalChart struct
	for cursor.Next(ctx) {
		var appraisalChart utils.NabbrAppraisalChart
		if err := cursor.Decode(&appraisalChart); err != nil {
			return nil, fmt.Errorf("error - MongoDB: Unable to decode job post: %w", err)
		}
		appraisalCharts = append(appraisalCharts, appraisalChart)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error - MongoDB: Cursor error: %w", err)
	}

	return appraisalCharts, nil
}

// retrieveDbFromMongoDbAll retrieves all job posts from the MongoDB collection.
func retrieveDbFromMongoDbAll() ([]utils.NabbrAppraisalChart, error) {
	resultAppraisalChart, err := retrieveDbFromMongoDbQuery(bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error - MongoDB: Unable to perform search query: %w", err)
	}
	return resultAppraisalChart, nil
}

// TODO: Need to review this function to ensure it is working as expected. I think it can be improved.
// func retrieveDbFromMongoDbSearch(searchTerm string) ([]utils.NabbrAppraisalChart, error) {
func retrieveDbFromMongoDbSearch(searchTerm string) ([]utils.NabbrAppraisalChart, error) {
	// Define the filter for the search query
	filter := bson.M{
		"$or": []bson.M{
			{"appraisalId": bson.M{"$regex": searchTerm}},
		},
	}

	resultAppraisalChart, err := retrieveDbFromMongoDbQuery(filter)
	if err != nil {
		return []utils.NabbrAppraisalChart{}, fmt.Errorf("error - MongoDB: Unable to perform search query: %w", err)
	}
	return resultAppraisalChart, nil
}
