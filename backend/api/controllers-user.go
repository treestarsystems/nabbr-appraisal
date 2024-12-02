package api

import (
	"context"
	"fmt"
	"log"
	"nabbr-appraisal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func createMatchStage() bson.D {
	return bson.D{{Key: "$match", Value: bson.D{}}}
}

func createGroupStage() bson.D {
	return bson.D{{Key: "$group", Value: bson.D{
		{Key: "_id", Value: "null"},
		{Key: "totalCount", Value: bson.D{{Key: "$sum", Value: 1}}},
		{Key: "data", Value: bson.D{{Key: "$push", Value: bson.D{
			{Key: "email", Value: "$email"},
			{Key: "firstName", Value: "$firstName"},
			{Key: "lastName", Value: "$lastName"},
			{Key: "phone", Value: "$phone"},
			{Key: "userId", Value: "$userId"},
			{Key: "userPrivilegeLevel", Value: "$userPrivilegeLevel"},
		}}}},
	}}}
}

func createProjectStage(startIndex, recordPerPage int) bson.D {
	return bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "totalCount", Value: 1},
			{Key: "userItems", Value: bson.D{{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}}}},
		}},
	}
}

func getAggregationPipeline(startIndex, recordPerPage int) mongo.Pipeline {
	return mongo.Pipeline{
		createMatchStage(),
		createGroupStage(),
		createProjectStage(startIndex, recordPerPage),
	}
}

func GetUsersAll(c *gin.Context) {
	if err := utils.CheckUserType(c, "ADMIN"); err != nil {
		apiResponse := utils.NewAPIResponse(
			"failuire",
			http.StatusUnauthorized,
			err.Error(),
			[]string{},
		)
		c.JSON(http.StatusUnauthorized, apiResponse)
		return
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 100
	}

	page, err1 := strconv.Atoi(c.Query("page"))
	if err1 != nil || page < 1 {
		page = 1
	}

	startIndex, err := strconv.Atoi(c.Query("startIndex"))
	if err != nil {
		startIndex = (page - 1) * recordPerPage
	}

	// Get the aggregation pipeline
	pipeline := getAggregationPipeline(startIndex, recordPerPage)
	result, err := utils.CollectionMongoUsers.Aggregate(ctx, pipeline)
	if err != nil {
		apiResponse := utils.NewAPIResponse(
			"failuire",
			http.StatusInternalServerError,
			"error - Get Users All: error occured while listing user items",
			[]string{},
		)
		c.JSON(http.StatusInternalServerError, apiResponse)
	}
	var allusers []bson.M
	if err = result.All(ctx, &allusers); err != nil {
		log.Fatal(err)
	}
	userItems := allusers[0]["userItems"]
	apiResponse := utils.NewAPIResponse(
		"success",
		http.StatusOK,
		"All users retrieved successfully",
		userItems,
	)
	c.JSON(http.StatusOK, apiResponse)
}

// GetUser is the api used to tget a single user
func GetUser(c *gin.Context) {
	userId := c.Param("userId")

	if err := utils.MatchUserTypeToUid(c, userId); err != nil {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusBadRequest,
			fmt.Sprintf("error - GetUser: (%v)", err.Error()),
			[]string{},
		)
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user utils.User

	err := utils.CollectionMongoUsers.FindOne(ctx, bson.M{"userId": userId}).Decode(&user)
	if err != nil {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusInternalServerError,
			fmt.Sprintf("error - GetUser: (%v)", err.Error()),
			[]string{},
		)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}
	apiResponse := utils.NewAPIResponse(
		"failure",
		http.StatusOK,
		"User retrieved successfully",
		[]interface{}{user},
	)
	c.JSON(http.StatusOK, apiResponse)
}
