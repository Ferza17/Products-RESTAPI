package orders

import "github.com/Ferza17/Products-RESTAPI/domains/orderDetails"

type OrderResponse struct {
	Message      string                     `json:"message"`
	Order        Order                      `json:"order"`
	OrderDetails []orderDetails.OrderDetail `json:"order_details"`
}
