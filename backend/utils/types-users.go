package utils

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                 primitive.ObjectID `bson:"_id"`
	FirstName          *string            `bson:"firstName" json:"firstName" validate:"required,min=2,max=100"`
	LastName           *string            `bson:"lastName" json:"lastName" validate:"required,min=2,max=100"`
	Password           *string            `bson:"password" json:"password" validate:"required,min=8,max=64"`
	Email              *string            `bson:"email" json:"email" validate:"email,required"`
	Phone              *string            `bson:"phone" json:"phone" validate:"required"`
	Token              *string            `bson:"token" json:"token"`
	RefreshToken       *string            `bson:"refreshToken" json:"refreshToken"`
	CreatedAt          time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt          time.Time          `bson:"updatedAt" json:"updatedAt"`
	UserId             string             `bson:"userId" json:"userId"`
	UserPrivilegeLevel string             `bson:"userPrivilegeLevel" json:"userPrivilegeLevel" validate:"required"`
	// UserRole           string             `bson:"userRole" json:"userRole"`
	// ServiceLevel       string             `bson:"serviceLevel" json:"serviceLevel"`
}

type UserLoginResponse struct {
	FirstName          *string `json:"firstName"`
	LastName           *string `json:"lastName"`
	Email              *string `json:"email"`
	Phone              *string `json:"phone"`
	Token              *string `json:"token"`
	RefreshToken       *string `json:"refreshToken"`
	UserId             string  `json:"userId"`
	UserPrivilegeLevel string  `json:"userPrivilegeLevel"`
	// UserRole           string  `json:"userRole"`
	// ServiceLevel       string  `json:"serviceLevel"`
}

// SignedDetails
type SignedDetails struct {
	Email              string
	FirstName          string
	LastName           string
	UserId             string
	UserPrivilegeLevel string
	jwt.StandardClaims
}
