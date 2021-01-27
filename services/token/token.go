package token

import (
	"fmt"
	"github.com/Ferza17/Products-RESTAPI/domains/customers"
	"github.com/Ferza17/Products-RESTAPI/utils/auth"
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
	"github.com/Ferza17/Products-RESTAPI/utils/logger"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strings"
)

var Services tokenInterface = &tokenStruct{}

type (
	tokenStruct struct {
	}
	tokenInterface interface {
		CreateToken(customer customers.Customers) (customers.CustomerLoginResponse, *errors.RestError)
		RefreshToken(tokenStr string) (customers.CustomerRefreshTokenResponse, *errors.RestError)
	}
)

func (t *tokenStruct) CreateToken(customer customers.Customers) (customers.CustomerLoginResponse, *errors.RestError) {
	token, expTime, err := auth.CreateToken(customer)
	if err != nil {
		return customers.CustomerLoginResponse{}, err
	}

	return customers.CustomerLoginResponse{
		TokenType:   "Bearer",
		AccessToken: token,
		Exp:         expTime,
	}, nil
}

func (t *tokenStruct) RefreshToken(tokenStr string) (customers.CustomerRefreshTokenResponse, *errors.RestError) {
	str := tokenStr
	str = strings.ReplaceAll(str, "Bearer ", "")
	token, _, err := new(jwt.Parser).ParseUnverified(str, jwt.MapClaims{})
	if err != nil {
		logger.Error("unable to decrypt jwt", err)
		return customers.CustomerRefreshTokenResponse{}, errors.NewBadRequestError("Bad Token Request")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Can't convert token's claims to standard claims")
		return customers.CustomerRefreshTokenResponse{}, errors.NewInternalServerError("Can't convert token's claims to standard claims")
	}

	// Validate
	customer := customers.Customers{
		Id: fmt.Sprint(claims["customer_id"]),
	}
	resultToken, _, tokenErr := auth.CreateToken(customer)
	if tokenErr != nil {
		return customers.CustomerRefreshTokenResponse{}, tokenErr
	}
	return customers.CustomerRefreshTokenResponse{
		TokenType:    "Bearer",
		AccessToken:  str,
		RefreshToken: resultToken,
	}, nil
}
