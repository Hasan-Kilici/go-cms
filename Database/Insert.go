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