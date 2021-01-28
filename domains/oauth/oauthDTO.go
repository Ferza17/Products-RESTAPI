package oauth

type Oauth struct {
	IdToken    string `json:"id_token"`
	CustomerId string `json:"customer_id"`
	OrderId    string `json:"order_id"`
	// varchar length(250)
	Token      string `json:"token"`
}
