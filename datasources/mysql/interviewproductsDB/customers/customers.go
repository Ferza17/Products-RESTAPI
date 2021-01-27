package customers

import (
	"database/sql"
	"fmt"
	"github.com/Ferza17/Products-RESTAPI/utils/env"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	Client   *sql.DB
	username = env.GetEnvironmentVariable("mysqlUsersUsername")
	password = env.GetEnvironmentVariable("mysql_users_password")
	host     = env.GetEnvironmentVariable("mysql_users_host")
	database = env.GetEnvironmentVariable("DATABASE")
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		database,
	)

	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	log.Println("Database Successfully configured")
}
