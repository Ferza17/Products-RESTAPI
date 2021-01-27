package Products

type Product struct {
	ProductId   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	BasicPrice  float64 `json:"basic_price"`
	CreatedDate string  `json:"created_date"`
}
