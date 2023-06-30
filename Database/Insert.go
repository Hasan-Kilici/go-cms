package Database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"CMS/Utils"
	"fmt"
	b64 "encoding/base64"
)

func Register(Username, Email, Password string) bool {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()
	
	if !FindUserByEmail(Email){
		Token := Utils.GenerateToken(31)
		NewPassword := b64.URLEncoding.EncodeToString([]byte(Password))

		res, err := db.Exec("INSERT INTO Users VALUES (?,?,?,?,?,?)","",Token,Username,Email,NewPassword,"User")
		if err != nil {
			fmt.Println(err)
			return false
		}
		
		rowCount, err := res.RowsAffected()
		if err != nil {
			fmt.Println(err)
			return false
		}
		
		fmt.Printf("%d satır eklendi\n", rowCount)
		return true
	} else {
		fmt.Println("Bu Email zaten Kullanılıyor")
		return false
	}
}

func CreateBlog(Title, HTML string){
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

	Token := Utils.GenerateToken(31)
	res, err := db.Exec("INSERT INTO Blogs VALUES (?,?,?,?,?,?)","",Token,Title,HTML,0,0)
	if err != nil {
		fmt.Println(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
		
	fmt.Printf("%d satır eklendi\n", rowCount)
}

func SaveUserLike(BlogToken, UserToken string) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

	Token := Utils.GenerateToken(31)
	res, err := db.Exec("INSERT INTO likes VALUES (?,?,?,?)","",Token,UserToken,BlogToken)
	if err != nil {
		fmt.Println(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
		
	fmt.Printf("%d satır eklendi\n", rowCount)
}