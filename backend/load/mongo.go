package load

import (
	"log"
	"nabbr-appraisal/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func loadDbDataToMongoDb(data utils.NabbrAppraisalChart, appraisalId string) error {
	filter := bson.M{"appraisalId": appraisalId}
	update := bson.M{
		"$set": bson.M{
			"memberInformation":    data.NabbrAppraisalChartMemberInfo,
			"petInformation":       data.NabbrAppraisalChartPetInfo,
			"appraisalInformation": data.NabbrAppraisalChartScoreAppraisalInformation,
		},
		"$setOnInsert": bson.M{
			"appraisalId": appraisalId,
		},
	}
	opts := options.Update().SetUpsert(true)

	_, err := utils.CollectionMongo.UpdateOne(utils.CtxMongo, filter, update, opts)
	if err != nil {
		log.Printf("error - MongoDB: Database write failure: %s\n", err)
		return err
	}
	return nil
}
