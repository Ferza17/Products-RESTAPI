package auth

import (
	"fmt"
	"github.com/Ferza17/Products-RESTAPI/domains/customers"
	"github.com/Ferza17/Products-RESTAPI/utils/env"
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
	"github.com/Ferza17/Products-RESTAPI/utils/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// Expiration Time in Minute
const ExpTimeMinute = 15

func CreateToken(customer customers.Customers) (string, int64, *errors.RestError) {
	// 15 minute to expire.
	expTime := time.Now().Add(time.Minute * ExpTimeMinute).Unix()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["customer_id"] = customer.Id
	atClaims["exp"] = expTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	resultToken, err := at.SignedString([]byte(env.GetEnvironmentVariable("JWTSECRET")))
	if err != nil {
		logger.Error("Error while creating token", err)
		return "", 0, errors.NewInternalServerError("Error while creating token")
	}

	return resultToken, expTime, nil
}

func ValidationToken(dataTime int64) bool {
	if dataTime > time.Now().Unix() {
		logger.Info("Token Not Expired")
		return false
	}

	logger.Info("Expired")
	return true
}

func Authentication(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	str := tokenString
	str = strings.ReplaceAll(str, "Bearer ", "")
	_, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("cant verify token")
		}
		return []byte(env.GetEnvironmentVariable("JWTSECRET")), nil
	})

	if err != nil {
		restErr := errors.NewUnauthorized("Unable to verify Token")
		c.JSON(restErr.Status, restErr)
		c.Abort()
		return
	}

	logger.Info("Token Verified!")
}
