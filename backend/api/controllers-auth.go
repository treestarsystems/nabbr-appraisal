package api

import (
	"context"
	"fmt"
	"nabbr-appraisal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

// CreateUser is the api used to tget a single user
func SignUp(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user utils.User

	if err := c.BindJSON(&user); err != nil {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusBadRequest,
			fmt.Sprintf("error - SignUp: (%v)", err.Error()),
			[]string{},
		)
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	// validate the userPrivilegeLevel
	if !utils.ValidateUserPrivilegeLevel(c, user.UserPrivilegeLevel) {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusBadRequest,
			"error - SignUp: User privilege level is invalid",
			[]string{},
		)
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusBadRequest,
			fmt.Sprintf("error - SignUp: (%v)", validationErr.Error()),
			[]string{},
		)
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	count, err := utils.CollectionMongoUsers.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusInternalServerError,
			fmt.Sprintf("error - SignUp: Email (%v)", err.Error()),
			[]string{},
		)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	if count > 0 {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusInternalServerError,
			"error - SignUp: Email already exits",
			[]string{},
		)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	count, err = utils.CollectionMongoUsers.CountDocuments(ctx, bson.M{"phone": user.Phone})
	if err != nil {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusInternalServerError,
			fmt.Sprintf("error - SignUp: Phone number (%v)", err.Error()),
			[]string{},
		)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	if count > 0 {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusInternalServerError,
			"error - SignUp: Phone number already exits",
			[]string{},
		)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	password := utils.HashPassword(*user.Password)
	user.Password = &password
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.UserId = user.ID.Hex()
	token, refreshToken, err := utils.GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, user.UserId, user.UserPrivilegeLevel)
	if err != nil {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusInternalServerError,
			fmt.Sprintf("error - SignUp: User was not created (%v)", err.Error()),
			[]string{},
		)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}
	user.Token = &token
	user.RefreshToken = &refreshToken
	resultInsertionNumber, insertErr := utils.CollectionMongoUsers.InsertOne(ctx, user)

	if insertErr != nil {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusInternalServerError,
			fmt.Sprintf("error - SignUp: User was not created (%v)", insertErr.Error()),
			[]string{},
		)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}
	apiResponse := utils.NewAPIResponse(
		"success",
		http.StatusOK,
		"User created successfully",
		[]map[string]interface{}{
			{
				"userId": resultInsertionNumber.InsertedID,
			},
		},
	)
	c.JSON(http.StatusOK, apiResponse)
}

// Login is the api used to tget a single user
func Login(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user utils.User
	var foundUser utils.User

	if err := c.BindJSON(&user); err != nil {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusInternalServerError,
			fmt.Sprintf("error - Login: (%v)", err.Error()),
			[]string{},
		)
		c.JSON(http.StatusBadRequest, apiResponse)
	}

	err := utils.CollectionMongoUsers.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusInternalServerError,
			fmt.Sprintf("error - Login: (%v)", err.Error()),
			[]string{},
		)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	passwordIsValid, msg := utils.VerifyPassword(*user.Password, *foundUser.Password)
	if !passwordIsValid {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusInternalServerError,
			fmt.Sprintf("error - Login: Verify Password (%v)", msg),
			[]string{},
		)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	token, refreshToken, err := utils.GenerateAllTokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, foundUser.UserId, foundUser.UserPrivilegeLevel)
	if err != nil {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusInternalServerError,
			fmt.Sprintf("error - Login: GenerateAllTokens (%v)", err.Error()),
			[]string{},
		)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}
	utils.UpdateAllTokens(token, refreshToken, foundUser.UserId)
	apiResponse := utils.NewAPIResponse(
		"success",
		http.StatusOK,
		"User logged in successfully",
		[]utils.UserLoginResponse{
			{
				FirstName:          foundUser.FirstName,
				LastName:           foundUser.LastName,
				Email:              foundUser.Email,
				Phone:              foundUser.Phone,
				Token:              &token,
				RefreshToken:       &refreshToken,
				UserId:             foundUser.UserId,
				UserPrivilegeLevel: foundUser.UserPrivilegeLevel,
			},
		},
	)
	c.JSON(http.StatusOK, apiResponse)
}
