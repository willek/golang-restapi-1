package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //mysql driver
)

// DBConnect db
func DBConnect() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:12345678@/go_sb_1")

	if err != nil {
		panic(err.Error())
	}

	return db
}
