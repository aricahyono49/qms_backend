package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "root"
	password string = ""
	database string = "qms_backend?parseTime=True"
	
	// username string = "backend-c-queue-management@tcp(35.201.240.187:3306)"
	// password string = "9MghTbFu5NCRt6lc"
	// database string = "backend-c-queue-management?parseTime=True"
)

var (
	dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)
)

// HubToMySQL
func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
