package postgres

import (
	"fmt"
	"jsbnch/pkg/utils/env"

	"github.com/go-pg/pg/v10"
)

// NewConnection provides connected database hander using os env variables
func NewConnection() *pg.DB {
	return CreateConnectionHandler(
		env.GetDatabaseAddress(),
		env.GetDatabaseUser(),
		env.GetDatabasePassword(),
		env.GetDatabaseName(),
	)
}

// CreateConnectionHandler provieds database handler manualy, without env veriables
func CreateConnectionHandler(addr, user, pwd, db string) *pg.DB {
	var handler = pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: pwd,
		Database: db,
	})

	// Test connection
	_, err := handler.Exec("SELECT 1")

	if err != nil {
		panic("[DATABASE] ERROR: " + err.Error())
	} else {
		fmt.Println("[DATABASE] Successfuly connected to", addr)
	}

	return handler
}
