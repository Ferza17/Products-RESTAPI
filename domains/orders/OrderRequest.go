package orders

import "github.com/Ferza17/Products-RESTAPI/domains/orderDetails"

type OrderRequest struct {
	Order       Order                      `json:"order"`
	OrderDetail []orderDetails.OrderDetail `json:"order_detail"`
}
