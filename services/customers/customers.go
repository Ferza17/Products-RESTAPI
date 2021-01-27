package customers

import (
	"fmt"
	customerDomain "github.com/Ferza17/Products-RESTAPI/domains/customers"
	"github.com/Ferza17/Products-RESTAPI/services/token"
	"github.com/Ferza17/Products-RESTAPI/utils/crypt"
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
	"github.com/Ferza17/Products-RESTAPI/utils/generate"
	"time"
)

var Service customersInterface = &customersStruct{}

type (
	customersStruct struct {
	}
	customersInterface interface {
		CreateCustomer(customers customerDomain.Customers) (*customerDomain.Customers, *errors.RestError)
		Login(request customerDomain.CustomerLoginRequest) (customerDomain.CustomerLoginResponse, *errors.RestError)
	}
)

func (c *customersStruct) CreateCustomer(customer customerDomain.Customers) (*customerDomain.Customers, *errors.RestError) {
	customer.Id = generate.GetRandomString(64)
	customer.Password = crypt.GetSHA265(customer.Password)
	t := time.Now()
	customer.CreatedDate = fmt.Sprintf("%d-%s-%d", t.Year(), t.Month(), t.Day())
	if err := customer.Validate(); err != nil {
		return nil, err
	}

	if err := customer.Save(); err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *customersStruct) Login(request customerDomain.CustomerLoginRequest) (customerDomain.CustomerLoginResponse, *errors.RestError) {

	if err := request.Validation(); err != nil {
		return customerDomain.CustomerLoginResponse{}, err
	}

	customer := customerDomain.Customers{
		Email:       request.PhoneNumberOrEmail,
		PhoneNumber: request.PhoneNumberOrEmail,
		Password:    crypt.GetSHA265(request.Password),
	}

	if err := customer.Login(); err != nil {
		return customerDomain.CustomerLoginResponse{}, err
	}

	result, err := token.Services.CreateToken(customer)
	if err != nil {
		return customerDomain.CustomerLoginResponse{}, err
	}

	return result, nil
}
