package customers

import (
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
	"github.com/Ferza17/Products-RESTAPI/utils/logger"
	customersDB "github.com/Ferza17/Products-RESTAPI/datasources/mysql/interviewproductsDB/customers"
)

const (
	queryInsert              = "INSERT INTO customers(customer_id, customer_name, email, phone_number, dob, sex, salt, password) VALUES (?,?,?,?,?,?,?,?)"
	queryFindByEmailAndPhone = "SELECT customer_id FROM customers WHERE phone_number=? OR email=? AND password=?"
)

func (c *Customers) Save() *errors.RestError {
	stmt, err := customersDB.Client.Prepare(queryInsert)
	if err != nil {
		logger.Error("error while preparing query", err)
		return errors.NewInternalServerError("Error processing data")
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.Id, c.CustomerName, c.Email, c.PhoneNumber, c.Dob, c.Sex, c.Salt, c.Password)
	if err != nil {
		logger.Error("Error while trying to exec query ", err)
		return errors.MysqlError(err)
	}
	return nil
}

func (c *Customers) Login() *errors.RestError {
	stmt, err := customersDB.Client.Prepare(queryFindByEmailAndPhone)
	if err != nil {
		logger.Error("error while preparing query", err)
		return errors.NewInternalServerError("Error processing data")
	}
	defer stmt.Close()

	row := stmt.QueryRow(c.Email, c.PhoneNumber, c.Password)
	if getErr := row.Scan(&c.Id); getErr != nil {
		return errors.NewInternalServerError("Error while scan data!")
	}

	return nil
}
