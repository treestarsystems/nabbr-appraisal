package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GenerateAllTokens generates both the detailed token and refresh token
func GenerateAllTokens(email string, firstName string, lastName string, userId string, userPrivilegeLevel string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:              email,
		FirstName:          firstName,
		LastName:           lastName,
		UserId:             userId,
		UserPrivilegeLevel: userPrivilegeLevel,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(AuthSecretKey))
	if err != nil {
		errMsg := fmt.Errorf("error - GenerateAllTokens: (%s)", err.Error())
		log.Print(errMsg)
		return "", "", errMsg
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(AuthSecretKey))
	if err != nil {
		errMsg := fmt.Errorf("error - GenerateAllTokens: (%s)", err.Error())
		log.Print(errMsg)
		return "", "", errMsg
	}

	return token, refreshToken, err
}

// UpdateAllTokens renews the user tokens when they login
func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
	updateObj = append(updateObj, bson.E{Key: "refreshToken", Value: signedRefreshToken})

	UpdatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updatedAt", Value: UpdatedAt})

	upsert := true
	filter := bson.M{"userId": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := CollectionMongoUsers.UpdateOne(
		ctx,
		filter,
		bson.D{
			{Key: "$set", Value: updateObj},
		},
		&opt,
	)
	defer cancel()

	if err != nil {
		log.Fatalf("error - UpdateAllTokens: (%s)", err)
		return
	}
}

// ValidateToken validates the jwt token
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(AuthSecretKey), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return claims, "error - Invalid token"
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return claims, "error - Expired token"
	}

	return claims, msg
}
