package db

import (
	"database/sql"
	"fmt"
	"main/config"

	_ "github.com/go-sql-driver/mysql"
)


// HubToMySQL
func MySQL(username string, pass string , dbName string ) (*sql.DB, error) {
	
var (
	dsn = fmt.Sprintf("%v:%v@/%v",username, pass, dbName)
)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
