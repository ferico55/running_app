package repository

import (
	"fmt"

	"github.com/ferico55/running_app/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func openDBConnection() *sqlx.DB {
	db, err := sqlx.Open(config.DriverName, config.ConnectionString)
	check(err)
	return db
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
