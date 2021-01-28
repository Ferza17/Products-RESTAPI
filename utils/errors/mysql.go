package errors

import "strings"

func MysqlError(err error) *RestError {
	if strings.Contains(err.Error(), "Duplicate entry") {
		if strings.Contains(err.Error(), "email") {
			return NewBadRequestError("Email Already Exist!")
		}

		if strings.Contains(err.Error(), "phone_number") {
			return NewBadRequestError("phone_number Already Exist!")
		}

		if strings.Contains(err.Error(), "token") {
			return NewBadRequestError("Token Only use Once !")
		}
	}

	return nil
}
