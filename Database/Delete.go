package Database

import (
	"fmt"
	"os"
	"log"
	"strings"
)

func DeleteBlog(Token string){
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

func DeleteAllProductImages(ImagePaths string) {
	Images := strings.Split(ImagePaths, ",")
	ImageCount := len(Images)

	for i := 0;i < ImageCount;i++ {

		Path := "./views/public"+Images[i]
		fmt.Println(Path)
		e := os.Remove(Path)
		if e != nil {
			log.Fatal(e)
		}
	}
} 

func DeleteProduct(Token string) {
	var deletedPhotoPaths string

	query := "SELECT path From ProductImages Where ProductToken = ?" 
	rows, err := db.Query(query,Token)
	if err != nil {
        panic(err.Error())
    }

	for rows.Next() {
		var path string
		err := rows.Scan(&path)
		if err != nil {
			panic(err.Error())
		}
		
		deletedPhotoPaths += path + ","
    }
	deletedPhotoPaths = strings.TrimRight(deletedPhotoPaths, ",")

	_, err = db.Exec("DELETE FROM productImages WHERE producttoken = ?", Token)
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("DELETE FROM tags WHERE producttoken = ?", Token)
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("DELETE FROM products WHERE token = ?", Token)
	if err != nil {
		panic(err.Error())
	}

	DeleteAllProductImages(deletedPhotoPaths)
}

func DeleteProductPhoto(Token string){
	var deletedPhotoPath string

	query := "SELECT path From ProductImages Where Token = ?" 
	rows, err := db.Query(query,Token)
	if err != nil {
        panic(err.Error())
    }

	for rows.Next() {
		var path string
		err := rows.Scan(&path)
		if err != nil {
			panic(err.Error())
		}
		
		deletedPhotoPath += path
    }

	Path := "./views/public"+deletedPhotoPath
	e := os.Remove(Path)
	if e != nil {
		log.Fatal(e)
	}

	_, err = db.Exec("DELETE FROM productImages WHERE token = ?", Token)
	if err != nil {
		panic(err.Error())
	}
}