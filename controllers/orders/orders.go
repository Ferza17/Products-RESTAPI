package orders

import (
	"github.com/Ferza17/Products-RESTAPI/domains/oauth"
	orderDomain "github.com/Ferza17/Products-RESTAPI/domains/orders"
	"github.com/Ferza17/Products-RESTAPI/services/order"
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CreateOrder(c *gin.Context) {
	orderRequest := orderDomain.OrderRequest{}
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body!")
		c.JSON(restErr.Status, restErr)
		return
	}

	// Get Token
	tokenString := c.Request.Header.Get("Authorization")
	str := tokenString
	str = strings.ReplaceAll(str, "Bearer ", "")
	token := oauth.Oauth{
		Token: str,
	}

	result, err := order.Services.CreateOrder(orderRequest, token)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
	return
}
