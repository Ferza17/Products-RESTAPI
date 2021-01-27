package customers

import (
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
)

type CustomerLoginRequest struct {
	PhoneNumberOrEmail string `json:"phone_number_or_email"`
	Password           string `json:"password"`
}


func (r *CustomerLoginRequest) Validation() *errors.RestError {
	if r.PhoneNumberOrEmail == "" {
		return errors.NewBadRequestError("Invalid phone number or email")
	}

	if r.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
