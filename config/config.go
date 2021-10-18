package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// HubToMySQL
func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/rakamin")

	if err != nil {
		return nil, err
	}

	return db, nil
}
