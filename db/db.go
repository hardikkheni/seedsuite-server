package db

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBFromTx(tx *sql.Tx) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New((postgres.Config{
		Conn: tx,
	})))
	if err != nil {
		return nil, err
	}
	return db, nil
}
