package Database

import (
	"fmt"
)

func GetBlogCount() int {
	var rowCount int
	err := db.QueryRow("SELECT COUNT(*) FROM blogs").Scan(&rowCount)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Satır sayısı: %d\n", rowCount)
	return rowCount
}

func GetUserCount() int {
	var rowCount int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&rowCount)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Satır sayısı: %d\n", rowCount)
	return rowCount
}

func GetGaleryCount() int {
	var rowCount int
	err := db.QueryRow("SELECT COUNT(*) FROM galery").Scan(&rowCount)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Satır sayısı: %d\n", rowCount)
	return rowCount
}

func GetProductCount() int {
	var rowCount int
	err := db.QueryRow("SELECT COUNT(*) FROM products").Scan(&rowCount)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Satır sayısı: %d\n", rowCount)
	return rowCount
}