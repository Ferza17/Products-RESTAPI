package oauth

import (
	"fmt"
	"github.com/Ferza17/Products-RESTAPI/datasources/mysql/oauthDB"
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
	"github.com/Ferza17/Products-RESTAPI/utils/logger"
)

const (
	QueryInsertToken = "INSERT INTO oauth(id_token, customer_id, order_id, token) VALUES (?,?,?,?)"
	QuerySearchToken = "SELECT COUNT(token) as total_token from oauth WHERE token=?"
)

func (o *Oauth) Save() *errors.RestError {
	stmt, err := oauthDB.Client.Prepare(QueryInsertToken)
	if err != nil {
		return errors.NewInternalServerError("Unable to prepare")
	}
	defer stmt.Close()

	_, err = stmt.Exec(o.IdToken, o.CustomerId, o.OrderId, o.Token)
	if err != nil {
		logger.Error("Err token", err)
		return errors.MysqlError(err)
	}
	return nil
}

func (o *Oauth) SearchToken() (bool, *errors.RestError) {
	// return true if found token in DB\
	stmt, err := oauthDB.Client.Prepare(QuerySearchToken)
	if err != nil {
		return false, errors.NewInternalServerError("unable to prepare query")
	}
	defer stmt.Close()

	rows, err := stmt.Query(o.Token)
	if err != nil {
		return false, errors.NewInternalServerError("Unable to exec query")
	}
	defer rows.Close()
	var count int64
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return false, errors.NewInternalServerError("Unable to scan data")
		}
	}

	logger.Info(fmt.Sprint(count))

	if count > 0 {
		return false, nil
	}
	return true, nil

}
