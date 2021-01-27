package auth

import (
	"encoding/json"
	"fmt"
	"github.com/Ferza17/Products-RESTAPI/domains/customers"
	"github.com/Ferza17/Products-RESTAPI/utils/env"
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
	"github.com/Ferza17/Products-RESTAPI/utils/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func ParseTimeToUnix(data interface{}) int64 {
	var tm time.Time
	switch iat := data.(type) {
	case float64:
		tm = time.Unix(int64(iat), 0)
	case json.Number:
		v, _ := iat.Int64()
		tm = time.Unix(v, 0)
	}
	return tm.Unix()
}

func Authentication(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}

		return []byte(env.GetEnvironmentVariable("JWTSECRET")), nil
	})

	if token != nil && err == nil {
		fmt.Println("Token Verified")
	} else {
		restErr := errors.NewUnauthorized("Invalid Bearer Token")
		c.JSON(restErr.Status, restErr)
	}
}
