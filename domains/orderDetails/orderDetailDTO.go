package orderDetails

type OrderDetail struct {
	OrderDetailId string              `json:"order_detail_id"`
	OrderId       string              `json:"order_id"`
	ProductId     string              `json:"product_id"`
	Qty           int64               `json:"qty"`
	CreatedDate   string              `json:"created_date"`
}
