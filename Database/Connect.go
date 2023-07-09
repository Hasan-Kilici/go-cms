package Database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
	if err != nil {
		panic(err.Error())
	}
}

func closeDB() {
	if db != nil {
		db.Close()
	}
}


func Connect() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
}