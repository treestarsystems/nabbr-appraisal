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
	var jobPosts []utils.NabbrAppraisalChart

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

	// Iterate through the cursor and decode each document into a JobPost
	for cursor.Next(ctx) {
		var jobPost utils.NabbrAppraisalChart
		if err := cursor.Decode(&jobPost); err != nil {
			return nil, fmt.Errorf("error - MongoDB: Unable to decode job post: %w", err)
		}
		jobPosts = append(jobPosts, jobPost)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error - MongoDB: Cursor error: %w", err)
	}

	return jobPosts, nil
}

// retrieveDbFromMongoDbAll retrieves all job posts from the MongoDB collection.
func retrieveDbFromMongoDbAll() ([]utils.NabbrAppraisalChart, error) {
	resultJobPosts, err := retrieveDbFromMongoDbQuery(bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error - MongoDB: Unable to perform search query: %w", err)
	}
	return resultJobPosts, nil
}

// TODO: Need to review this function to ensure it is working as expected. I think it can be improved.
// SearchJobPostsInMongoDb searches for job posts in MongoDB based on a search term.
func retrieveDbFromMongoDbSearch(searchTerm string) ([]utils.NabbrAppraisalChart, error) {
	// Define the filter for the search query
	filter := bson.M{
		"$or": []bson.M{
			{"appraisalId": bson.M{"$regex": searchTerm}},
		},
	}

	resultJobPosts, err := retrieveDbFromMongoDbQuery(filter)
	if err != nil {
		return nil, fmt.Errorf("error - MongoDB: Unable to perform search query: %w", err)
	}
	return resultJobPosts, nil
}
