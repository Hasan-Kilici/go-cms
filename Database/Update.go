package Database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func UpdateBlog(Token,Title, HTML string) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

	query := "UPDATE blogs SET Title=?, HTML=? WHERE Token=?"
	res, err := db.Exec(query,Title,HTML,Token)
	if err != nil {
		fmt.Println(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d satır düzenlendi\n", rowCount)
}

func AddViewToBlog(Token string, Views int){
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

	query := "UPDATE blogs SET Views=? WHERE Token=?"
	res, err := db.Exec(query,Views,Token)
	if err != nil {
		fmt.Println(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d satır düzenlendi\n", rowCount)
}

func AddLikeToBlog(Token string, Like int){
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

	query := "UPDATE blogs SET blike=? WHERE Token=?"
	res, err := db.Exec(query,Like,Token)
	if err != nil {
		fmt.Println(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d satır düzenlendi\n", rowCount)
}

func SubLikeToBlog(Token string, Like int){
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

	query := "UPDATE blogs SET blike=? WHERE Token=?"
	res, err := db.Exec(query,Like,Token)
	if err != nil {
		fmt.Println(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d satır düzenlendi\n", rowCount)
}

func UpdateUser(Token ,Username, Email, Perm string){
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

	query := "UPDATE users SET username=?, email=?, perm=? WHERE Token=?"
	res, err := db.Exec(query,Username,Email,Perm,Token)
	if err != nil {
		fmt.Println(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d satır düzenlendi\n", rowCount)
}

func UpdateGalery(Token,Title,Description string){
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        fmt.Println(err)
    }

    defer db.Close()

	query := "UPDATE galerypropertys SET title=?, description=? WHERE galeryToken=?"
	res, err := db.Exec(query,Title,Description,Token)
	if err != nil {
		fmt.Println(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d satır düzenlendi\n", rowCount)
}