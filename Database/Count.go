package Database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)


func GetBlogCount() int {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        panic(err.Error())
    }

	defer db.Close()

	var rowCount int
	err = db.QueryRow("SELECT COUNT(*) FROM blogs").Scan(&rowCount)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Satır sayısı: %d\n", rowCount)
	return rowCount
}

func GetUserCount() int {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        panic(err.Error())
    }

	defer db.Close()

	var rowCount int
	err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&rowCount)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Satır sayısı: %d\n", rowCount)
	return rowCount
}