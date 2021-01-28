package oauth

type Oauth struct {
	IdToken    string `json:"id_token"`
	CustomerId string `json:"customer_id"`
	OrderId    string `json:"order_id"`
	Token      string `json:"token"`
}
