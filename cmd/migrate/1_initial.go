package main

import (
	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	"jsbnch/pkg/model"
)

func init() {
	migrations.MustRegisterTx(
		// UP
		func(db migrations.DB) error {
			var transaction = db.(*pg.Tx)

			transaction.Model(&model.User{}).CreateTable(&orm.CreateTableOptions{IfNotExists: true})

			return nil
		},

		// DOWN
		func(db migrations.DB) error {
			var transaction = db.(*pg.Tx)

			transaction.Model(&model.User{}).DropTable(&orm.DropTableOptions{IfExists: true})

			return nil
		},
	)
}
