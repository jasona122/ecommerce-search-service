package db

import (
	"github.com/jasona122/ecommerce-search-service/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func Init(dbConfig config.DatabaseConfig) error {
	tempDB, err := sqlx.Open(dbConfig.Driver, dbConfig.GetConnectionURL())
	if err != nil {
		return err
	}

	err = tempDB.Ping()
	if err != nil {
		return err
	}

	db = tempDB
	return nil
}

func GetDB() *sqlx.DB {
	return db
}
