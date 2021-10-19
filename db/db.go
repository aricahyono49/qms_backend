package db

import (
	"database/sql"
	"fmt"
	"main/config"

	_ "github.com/go-sql-driver/mysql"
)

// HubToMySQL
func MySQL() (*sql.DB, error) {

	var (
		dsn = fmt.Sprintf("%v:%v@/%v", config.Configuration.DBUsername, config.Configuration.DBPassword, config.Configuration.DBName)
	)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
