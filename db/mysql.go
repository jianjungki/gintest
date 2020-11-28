package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@/gintest")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
