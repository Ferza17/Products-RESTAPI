package order

import (
	"fmt"
	oauthDomain "github.com/Ferza17/Products-RESTAPI/domains/oauth"
	"github.com/Ferza17/Products-RESTAPI/domains/orderDetails"
	orderDomain "github.com/Ferza17/Products-RESTAPI/domains/orders"
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
	"github.com/Ferza17/Products-RESTAPI/utils/generate"
	"time"
)

var Services orderInterface = &orderStruct{}

type (
	orderStruct struct {
	}
	orderInterface interface {
		CreateOrder(order orderDomain.OrderRequest, oauth oauthDomain.Oauth) (*orderDomain.OrderResponse, *errors.RestError)
	}
)

func (o *orderStruct) CreateOrder(order orderDomain.OrderRequest, oauth oauthDomain.Oauth) (*orderDomain.OrderResponse, *errors.RestError) {
	// Token Only one use
	dao := order.Order
	// Generate ID order and Id order Details
	dao.OrderId = generate.GetRandomString(64, "OrderIdAutoGenerate")
	//Order Number
	dao.OrderNumber = generate.GetOrderNumber()
	tm := time.Now()
	timeNow := fmt.Sprint(tm.Year(), "-", int(tm.Month()), "-", tm.Day())
	dao.OrderDate = timeNow

	var orderDetail []orderDetails.OrderDetail
	for _, detail := range order.OrderDetail {
		generateID := generate.GetRandomString(64, "OrderDetailsAutoGenerate")
		orderDetail = append(orderDetail, orderDetails.OrderDetail{
			OrderDetailId: generateID,
			ProductId:     detail.ProductId,
			OrderId:       dao.OrderId,
			Qty:           detail.Qty,
			CreatedDate:   timeNow,
		})
	}

	// Check if token has already been used then return error
	// Insert token, if token already exist return token only use once

	// insert token to oauthDB
	oauth.OrderId = dao.OrderId
	oauth.CustomerId = dao.CustomerId
	oauth.IdToken = generate.GetRandomString(64, "OauthIdToken")
	if oauthErr := oauth.Save(); oauthErr != nil {
		return nil, oauthErr
	}

	result, err := dao.CreateOrder(orderDetail)
	if err != nil {
		return nil, err
	}

	return result, nil
}
