package postgres

import (
	"touchstone-api/external/migrations"
	"touchstone-api/pkg/utils/env"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// Connection represents connection handler or transaction
type Connection interface {
	Model(model ...interface{}) *orm.Query
}

// NewConnection provides connected database hander using os env variables
func NewConnection(runMigrations bool) *pg.DB {
	db := CreateConnectionHandler(
		env.GetDatabaseAddress(),
		env.GetDatabaseUser(),
		env.GetDatabasePassword(),
		env.GetDatabaseName(),
	)

	if runMigrations == true {
		migrations.RunMigrations(db)
	}

	return db
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
	}

	return handler
}
