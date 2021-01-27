package customers

import (
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
	"encoding/json"
)

type (
	PublicCustomer struct {
		Id          string `json:"customer_id"`
		CreatedDate string `json:"created_date"`
	}

	PrivateCustomer struct {
		Id          string `json:"customer_id"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Dob         string `json:"dob"`
		Sex         bool   `json:"sex"`
		Salt        string `json:"salt"`
		CreatedDate string `json:"created_date"`
	}
)

func (c *Customers) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicCustomer{
			Id:          c.Id,
			CreatedDate: c.CreatedDate,
		}
	}

	customerJSON, _ := json.Marshal(c)
	var privateCustomer PrivateCustomer
	if err := json.Unmarshal(customerJSON, &privateCustomer); err != nil {
		return errors.NewInternalServerError("Error When unmarshalling Body")
	}
	return privateCustomer
}
