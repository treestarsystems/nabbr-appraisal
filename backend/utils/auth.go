package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var AuthSecretKey = os.Getenv("AUTH_SECRET_KEY")

// Vvalidate the userPrivilegeLevel
func ValidateUserPrivilegeLevel(c *gin.Context, userPrivilegeLevel string) bool {
	pattern := `^(ADMIN|APPRAISER|PETOWNER)$`
	matched, _ := regexp.MatchString(pattern, userPrivilegeLevel)
	return matched
}

// VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	if err != nil {
		return false, "error - Verify Password: Login or Passowrd is incorrect"
	}
	return true, ""
}

// HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatalf("error - HashPassword: (%s)", err)
	}
	return string(bytes)
}

// CheckUserType renews the user tokens when they login
func CheckUserType(c *gin.Context, role string) (err error) {
	userPrivilegeLevel := c.GetString("userPrivilegeLevel")
	err = nil
	if userPrivilegeLevel != role {
		return errors.New("error - Check User Privilege: Unauthorized to access this resource")
	}
	return err
}

// MatchUserTypeToUid only allows the user to access their data and no other data. Only the admin can access all user data
func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userPrivilegeLevel := c.GetString("userPrivilegeLevel")
	uid := c.GetString("userId")
	err = nil
	if userPrivilegeLevel == "USER" && uid != userId {
		return errors.New("error - Check User Privilege: Unauthorized to access this resource")
	}
	err = CheckUserType(c, userPrivilegeLevel)
	return err
}

// Auth validates token and authorizes users
func Authentication(c *gin.Context) {
	clientToken := c.Request.Header.Get("token")
	if clientToken == "" {
		apiResponse := NewAPIResponse(
			"failure",
			http.StatusInternalServerError,
			"error - Authentication: No Authorization header provided",
			[]string{},
		)
		c.JSON(http.StatusInternalServerError, apiResponse)
		c.Abort()
		return
	}

	claims, err := ValidateToken(clientToken)
	if err != "" {
		apiResponse := NewAPIResponse(
			"failure",
			http.StatusInternalServerError,
			fmt.Sprintf("error - Authentication: No Authorization header provided (%v)", err),
			[]string{},
		)
		c.JSON(http.StatusInternalServerError, apiResponse)
		c.Abort()
		return
	}

	c.Set("email", claims.Email)
	c.Set("firstName", claims.FirstName)
	c.Set("lastName", claims.LastName)
	c.Set("userId", claims.UserId)
	c.Set("userPrivilegeLevel", claims.UserPrivilegeLevel)
	c.Next()
}
