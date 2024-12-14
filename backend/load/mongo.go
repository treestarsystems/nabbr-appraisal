package load

import (
	"context"
	"errors"
	"log"
	"nabbr-appraisal/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func loadDbDataToMongoDbAppraisals(data utils.NabbrAppraisalChart, appraisalId string) (string, error) {
	filter := bson.M{"appraisalId": appraisalId}
	update := bson.M{
		"$set": bson.M{
			"memberInformation":    data.MemberInformation,
			"petInformation":       data.PetInformation,
			"appraisalInformation": data.AppraisalInformation,
		},
		"$setOnInsert": bson.M{
			"appraisalId": appraisalId,
		},
	}
	opts := options.Update().SetUpsert(true)
	_, err := utils.CollectionMongoAppraisals.UpdateOne(utils.CtxMongo, filter, update, opts)
	if err != nil {
		log.Printf("error - MongoDB: Database write failure: %s\n", err)
		return "", err
	}
	return appraisalId, nil
}

func loadDbDataToMongoDbAppraisalDelete(appraisalId string) error {
	filter := bson.M{"appraisalId": appraisalId}

	// Check if the document exists
	var result bson.M
	err := utils.CollectionMongoAppraisals.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("chart not found")
		}
		return errors.New("db read failure")
	}

	// Document found, proceed with deletion
	opts := options.Delete()
	_, err = utils.CollectionMongoAppraisals.DeleteOne(context.TODO(), filter, opts)
	if err != nil {
		return err
	}
	return nil
}
