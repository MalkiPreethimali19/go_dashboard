package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB
// DB connection
func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "root:your_pasword@tcp(127.0.0.1:3306)/dashboard_db") // add your password here
	if err != nil {
        log.Fatal("Connection error:", err)
    }

	if err = DB.Ping(); err != nil {
        log.Fatal("Ping failed:", err)
    }
	
}