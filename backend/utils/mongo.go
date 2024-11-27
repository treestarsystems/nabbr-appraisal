package utils

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB variables
var ClientMongo *mongo.Client
var CollectionMongoUsers *mongo.Collection
var CollectionMongoAppraisals *mongo.Collection
var CtxMongo = context.TODO()

func LoadDbConnectToMongoDb() {
	mongoDbUri := os.Getenv("DB_MONGODB_URI")
	mongoDbName := os.Getenv("DB_NAME")
	mongoDbCollectionNameUsers := "users"
	mongoDbCollectionNameAppraisals := os.Getenv("DB_TABLE_NAME_APPRAISALS")
	clientOptions := options.Client().ApplyURI(mongoDbUri)
	ClientMongo, err := mongo.Connect(CtxMongo, clientOptions)
	if err != nil {
		log.Fatalf("error - MongoDB: Unable to establish database connection: %s", err)
	}
	// This reduces the codes resilience to failure. So we may want to remove this.
	err = ClientMongo.Ping(CtxMongo, nil)
	if err != nil {
		log.Fatalf("error - MongoDB: Unable to ping database instance: %s", err)
	}

	// Users collection
	CollectionMongoUsers = ClientMongo.Database(mongoDbName).Collection(mongoDbCollectionNameUsers)
	// Appraisal collection
	CollectionMongoAppraisals = ClientMongo.Database(mongoDbName).Collection(mongoDbCollectionNameAppraisals)
}
