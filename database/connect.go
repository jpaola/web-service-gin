package database

import (
	"database/sql"
	"fmt"
)

func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	hostname := "tcp(127.0.0.1:3306)"
	dbName := "database"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+hostname+"/"+dbName)
	
	// If there is an error during db connection print it to the console
	if err != nil {
		fmt.Println("DB Connection error", err.Error())
		return
	}

	return db
}