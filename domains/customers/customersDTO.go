package customers

import (
	"fmt"
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
	"regexp"
)

type Customers struct {
	Id          string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Dob         string `json:"dob"`
	Sex         bool   `json:"sex"`
	Salt        string `json:"salt"`
	Password    string `json:"password"`
	CreatedDate string `json:"created_date"`
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (c *Customers) Validate() *errors.RestError {
	if c.Email == "" {
		return errors.NewBadRequestError("Please provide email!")
	}

	if !isEmailValid(c.Email) {
		return errors.NewBadRequestError(fmt.Sprintf("%s is not valid email!", c.Email))
	}

	if c.CustomerName == "" {
		return errors.NewBadRequestError("Please Provide Customer Name!")
	}

	if c.PhoneNumber == "" {
		return errors.NewBadRequestError("Please provide phone number!")
	}

	if c.Dob == "" {
		return errors.NewBadRequestError("Please provide dob!")
	}

	if c.Salt == "" {
		return errors.NewBadRequestError("Please provide salt!")
	}

	if c.Password == "" {
		return errors.NewBadRequestError("Please provide Password!")
	}
	return nil
}

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
