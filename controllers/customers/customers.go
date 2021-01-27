package customers

import (
	customerDomain "github.com/Ferza17/Products-RESTAPI/domains/customers"
	customerService "github.com/Ferza17/Products-RESTAPI/services/customers"
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCustomer(c *gin.Context) {
	var customer customerDomain.Customers
	if err := c.ShouldBindJSON(&customer); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body!")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, serviceErr := customerService.Service.CreateCustomer(customer)

	if serviceErr != nil {
		c.JSON(serviceErr.Status, serviceErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

