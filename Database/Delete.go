package Database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
	"log"
)

func DeleteBlog(Token string){
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()
	
	query := "DELETE FROM blogs WHERE Token=?"
	res, err := db.Exec(query,Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	DeleteBlogTags(Token)
	DeleteLikes(Token)
	fmt.Printf("%d satır Silindi\n", rowCount)
}

func DeleteBlogTags(Token string) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()
	
	query := "DELETE FROM tags WHERE BlogToken=?"
	res, err := db.Exec(query,Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d satır Silindi\n", rowCount)	
}

func DeleteUser(Token string){
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()
	
	query := "DELETE FROM users WHERE Token=?" 
	res, err := db.Exec(query,Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d satır Silindi\n", rowCount)	
}

func DeleteUserLike(UserToken, BlogToken string){
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()
	
	query := "DELETE FROM likes WHERE UserToken=? AND BlogToken=?" 
	res, err := db.Exec(query,UserToken,BlogToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d satır Silindi\n", rowCount)	
}

func DeleteLikes(BlogToken string) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()
	
	query := "DELETE FROM likes WHERE BlogToken=?" 
	res, err := db.Exec(query,BlogToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d satır Silindi\n", rowCount)		
}

func DeleteGalery(Token string) {
	Photo, err := FindPhotoByToken(Token)
	if err != nil {
		fmt.Println("Galeride fotoğraf bulunamadı")
		return
	}

	Path := "./views/public"+Photo.Path
	e := os.Remove(Path)
    if e != nil {
        log.Fatal(e)
    }

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

	query := "DELETE FROM galery WHERE Token=?" 
	res, err := db.Exec(query,Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	DeleteGaleryProps(Token)

	fmt.Printf("%d satır Silindi\n", rowCount)	
}

func DeleteGaleryProps(GToken string) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

	query := "DELETE FROM galerypropertys WHERE galerytoken=?" 
	res, err := db.Exec(query,GToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d satır Silindi\n", rowCount)		
}