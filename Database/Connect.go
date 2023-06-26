package Database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
}