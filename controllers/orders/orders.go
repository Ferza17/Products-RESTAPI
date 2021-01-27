package orders

import (
	orderDomain "github.com/Ferza17/Products-RESTAPI/domains/orders"
	"github.com/Ferza17/Products-RESTAPI/services/order"
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	orderRequest := orderDomain.OrderRequest{}
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body!")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, err := order.Services.CreateOrder(orderRequest)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
