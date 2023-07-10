package Database

import (
	"fmt"
)

func UpdateBlog(Token,Title, HTML string) {
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

func UpdateProduct(Token , Name, Description, Tags string, Price int) {
	query := "UPDATE products SET name=?, price=?,description=? WHERE Token=?"
	res, err := db.Exec(query,Name,Price,Description,Token)
	if err != nil {
		fmt.Println(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d satır düzenlendi\n", rowCount)
	_, err = db.Exec("DELETE FROM tags WHERE producttoken = ?", Token)
	if err != nil {
		panic(err.Error())
	}

	SaveAllProductTags(Token, Tags)
}