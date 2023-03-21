package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"

	"github.com/hardikkheni/seedsuite-server/db"
	"github.com/hardikkheni/seedsuite-server/db/models"
)

func init() {
	goose.AddMigration(upCreateUserTable, downCreateUserTable)
}

func upCreateUserTable(tx *sql.Tx) error {
	db, err := db.DBFromTx(tx)
	if err != nil {
		return err
	}
	db.Migrator().CreateTable(&models.User{})
	return nil
}

func downCreateUserTable(tx *sql.Tx) error {
	db, err := db.DBFromTx(tx)
	if err != nil {
		return err
	}
	db.Migrator().DropTable(&models.User{})
	return nil
}
