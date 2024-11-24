package api

import (
	"context"
	"fmt"
	"log"
	"nabbr-appraisal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

// CreateUser is the api used to tget a single user
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user utils.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":     "failure",
				"httpStatus": http.StatusBadRequest,
				"message":    fmt.Sprintf("error - SignUp: %v", err.Error()),
				"payload":    []string{},
			})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":     "failure",
				"httpStatus": http.StatusBadRequest,
				"message":    fmt.Sprintf("error - SignUp: %v", validationErr.Error()),
				"payload":    []string{},
			})
			return
		}

		count, err := utils.CollectionMongoUsers.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":     "failure",
				"httpStatus": http.StatusInternalServerError,
				"message":    fmt.Sprintf("error - SignUp Checking email: %v", err.Error()),
				"payload":    []string{},
			})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":     "failure",
				"httpStatus": http.StatusInternalServerError,
				"message":    "error - SignUp Checking email already exits",
				"payload":    []string{},
			})
			return
		}

		count, err = utils.CollectionMongoUsers.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":     "failure",
				"httpStatus": http.StatusInternalServerError,
				"message":    fmt.Sprintf("error - SignUp Checking phone number: %v", err.Error()),
				"payload":    []string{},
			})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":     "failure",
				"httpStatus": http.StatusInternalServerError,
				"message":    "error - SignUp Checking phone number already exits",
				"payload":    []string{},
			})
			return
		}

		password := utils.HashPassword(*user.Password)
		user.Password = &password
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserId = user.ID.Hex()
		token, refreshToken, _ := utils.GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, user.UserId)
		user.Token = &token
		user.RefreshToken = &refreshToken
		resultInsertionNumber, insertErr := utils.CollectionMongoUsers.InsertOne(ctx, user)

		if insertErr != nil {
			// msg := fmt.Sprintf("User item was not created")
			// c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":     "failure",
				"httpStatus": http.StatusInternalServerError,
				"message":    fmt.Sprintf("error - SignUp user was not created: %v", insertErr.Error()),
				"payload":    []string{},
			})
			return
		}
		defer cancel()

		// c.JSON(http.StatusOK, resultInsertionNumber)
		c.JSON(http.StatusOK, gin.H{
			"status":     "success",
			"httpStatus": http.StatusOK,
			"message":    "User created successfully",
			"payload":    []*mongo.InsertOneResult{resultInsertionNumber},
		})

	}
}

// Login is the api used to tget a single user
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user utils.User
		var foundUser utils.User

		if err := c.BindJSON(&user); err != nil {
			// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.JSON(http.StatusBadRequest, gin.H{
				"status":     "failure",
				"httpStatus": http.StatusInternalServerError,
				"message":    fmt.Sprintf("error - Login : %v", err.Error()),
				"payload":    []string{},
			})
		}

		err := utils.CollectionMongoUsers.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": "login or passowrd is incorrect"})
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":     "failure",
				"httpStatus": http.StatusInternalServerError,
				"message":    fmt.Sprintf("error - Login : %v", err.Error()),
				"payload":    []string{},
			})
			return
		}

		passwordIsValid, msg := utils.VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if !passwordIsValid {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":     "failure",
				"httpStatus": http.StatusInternalServerError,
				"message":    fmt.Sprintf("error - Login Verify Password: %v", msg),
				"payload":    []string{},
			})
			return
		}

		token, refreshToken, _ := utils.GenerateAllTokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, foundUser.UserId)
		utils.UpdateAllTokens(token, refreshToken, foundUser.UserId)
		c.JSON(http.StatusOK, foundUser)
	}
}
