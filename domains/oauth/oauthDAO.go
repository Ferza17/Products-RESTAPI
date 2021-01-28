package oauth

import (
	"github.com/Ferza17/Products-RESTAPI/datasources/mysql/oauthDB"
	"github.com/Ferza17/Products-RESTAPI/utils/errors"
	"github.com/Ferza17/Products-RESTAPI/utils/logger"
)

const (
	QueryInsertToken = "INSERT INTO oauth(id_token, customer_id, order_id, token) VALUES (?,?,?,?)"
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