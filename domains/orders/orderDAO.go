package orders

import (
	"github.com/Ferza17/Products-RESTAPI/datasources/mysql/interviewproductsDB"
	"github.com/Ferza17/Products-RESTAPI/domains/orderDetails"
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
	"github.com/Ferza17/Products-RESTAPI/utils/logger"
)

const (
	QueryInsertOrder        = "INSERT INTO orders(order_id,customer_id,order_number, order_date, payment_method_id  ) VALUES (?,?,?,?,?)"
	QueryInsertDetailsOrder = "INSERT INTO orderdetails (order_detail_id, order_id, product_id, qty) VALUES (?,?,?,?)"
)

// Implement with Transaction in Mysql
func (o *Order) CreateOrder(orderDetails []orderDetails.OrderDetail) (*OrderResponse, *errors.RestError) {
	result := &OrderResponse{}
	// Insert Order
	stmtOrder, err := interviewproductsDB.Client.Prepare(QueryInsertOrder)
	if err != nil {
		return nil, errors.NewInternalServerError("Error when trying to insert data")
	}

	_, err = stmtOrder.Exec(o.OrderId, o.CustomerId, o.OrderNumber, o.OrderDate, o.PaymentMethodId)
	if err != nil {
		return nil, errors.NewInternalServerError("Error when trying to execute")
	}
	_ = stmtOrder.Close()

	// Insert Order Details
	for _, orderItem := range orderDetails {
		stmtOrderDetails, err := interviewproductsDB.Client.Prepare(QueryInsertDetailsOrder)
		if err != nil {
			return nil, errors.NewInternalServerError("Error While trying insert order details")
		}

		_, err = stmtOrderDetails.Exec(orderItem.OrderDetailId, orderItem.OrderId, orderItem.ProductId, orderItem.Qty)
		if err != nil {
			logger.Error("err order details", err)
			return nil, errors.NewInternalServerError("Error while trying execute order id")
		}
		_ = stmtOrderDetails.Close()
	}

	result.Order = Order{
		OrderId:         o.OrderId,
		CustomerId:      o.CustomerId,
		PaymentMethodId: o.PaymentMethodId,
		OrderDate:       o.OrderDate,
		OrderNumber:     o.OrderNumber,
	}

	result.Message = "order Created!"
	result.OrderDetails = orderDetails

	return result, nil
}
