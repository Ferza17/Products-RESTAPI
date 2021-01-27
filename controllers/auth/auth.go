package auth

import (
	"github.com/Ferza17/Products-RESTAPI/domains/customers"
	customerServices "github.com/Ferza17/Products-RESTAPI/services/customers"
	tokenServices "github.com/Ferza17/Products-RESTAPI/services/token"
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)
const authHeader = "Authorization"

func CreateAuthToken(c *gin.Context) {
	var customerRequest customers.CustomerLoginRequest
	if err := c.ShouldBindJSON(&customerRequest); err != nil {
		restErr := errors.NewBadRequestError("Invalid json Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := customerServices.Service.Login(customerRequest)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}

func RefreshToken(c *gin.Context) {
	token := c.Request.Header.Get(authHeader)
	res, err := tokenServices.Services.RefreshToken(token)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
