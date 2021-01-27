package orders

type Order struct {
	OrderId         string                      `json:"order_id"`
	CustomerId      string                      `json:"customer_id"`
	OrderNumber     string                      `json:"order_number"`
	OrderDate       string                      `json:"order_date"`
	PaymentMethodId string                      `json:"payment_method_id"`
}
