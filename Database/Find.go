package Database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	b64 "encoding/base64"
	"errors"
	"fmt"
)

type User struct{
	ID			int
	Token		string
	Username	string
	Email		string
	Password	string
	Perm		string
}

type Blog struct {
	ID				int
	Token			string
	Title			string
	HTML		    string
    Views           int
	Like		    int
}

type Galery struct {
    ID              int
    Token           string
    Path            string
}

type GaleryProps struct {
    ID              int
    Token           string
    Title           string
    Description     string
    GaleryToken     string
}

type Product struct {
    ID              int
    Token           string
    Name            string
    Price           int
    Description     string
    ImagePath       string
}

type EditProduct struct {
    ID              int
    Token           string
    Name            string
    Price           int
    Description     string
}

func Login(Email, Password string) (string, error) {
	query := "SELECT Token, Password FROM users WHERE Email = ?"
	row := db.QueryRow(query, Email)

	var token, hashedPassword string
	row.Scan(&token, &hashedPassword)
	RealPassword, _ := b64.URLEncoding.DecodeString(hashedPassword)

	if(string(RealPassword) == Password){
		return token, nil
	}

	return "Kullanıcı Bulunamadı", errors.New("Email ya da şifre yanlış")
}

func FindUserByEmail(Email string) bool {
    query := "SELECT COUNT(*) FROM users WHERE Email = ?"
    row := db.QueryRow(query, Email)

    var count int
    err := row.Scan(&count)
    if err != nil {
        return false
    }

    if count == 0 {
        return false
    }

    return true
}

func FindUserByUsername(Username string) bool {
    query := "SELECT COUNT(*) FROM users WHERE username = ?"
    row := db.QueryRow(query, Username)

    var count int
    err := row.Scan(&count)
    if err != nil {
        return false
    }

    if count == 0 {
        return false
    }

    return true
}

func FindUserByToken(Token string) (User, error){
    query := "SELECT * FROM users WHERE token = ?"
    row := db.QueryRow(query, Token)

    var user User
    err := row.Scan(&user.ID, &user.Token, &user.Username, &user.Email ,&user.Password, &user.Perm)
    if err != nil {
        if err == sql.ErrNoRows {
            return User{}, fmt.Errorf("kullanıcı bulunamadı")
        }
        return User{}, err
    }
    return user, nil
}

func FindBlogByToken(Token string) (Blog, error) {
    query := "SELECT * FROM blogs WHERE token = ?"
    row := db.QueryRow(query, Token)

    var blog Blog
    err := row.Scan(&blog.ID, &blog.Token, &blog.Title, &blog.HTML ,&blog.Views, &blog.Like)
    if err != nil {
        if err == sql.ErrNoRows {
            return Blog{}, fmt.Errorf("kullanıcı bulunamadı")
        }
        return Blog{}, err
    }
    return blog, nil
}

func FindBlogLike(UserToken, BlogToken string) bool {
    query := "SELECT COUNT(*) FROM likes WHERE UserToken = ? AND BlogToken = ?"
    row := db.QueryRow(query, UserToken, BlogToken)
    
    var count int
    err := row.Scan(&count)
    if err != nil {
        fmt.Println("Like bulunamadı")
        return false
    }
    
    if count == 0 {
        fmt.Println("Like bulunamadı")
        return false
    } else {
        fmt.Println("Like bulundu")
        return true
    }
    
    return true
}

func FindPhotoByToken(Token string) (Galery, error){
    query := "SELECT * FROM galery WHERE token = ?"
    row := db.QueryRow(query, Token)

    var galery Galery
    err := row.Scan(&galery.ID, &galery.Token, &galery.Path)
    if err != nil {
        if err == sql.ErrNoRows {
            return Galery{}, err
        }
        return Galery{}, err
    }

    return galery, nil
}

func FindGaleryPropsByToken(Token string) (GaleryProps, error){
    query := "SELECT * FROM galerypropertys WHERE galerytoken = ?"
    row := db.QueryRow(query, Token)

    var galeryProps GaleryProps
    err := row.Scan(&galeryProps.ID, &galeryProps.Token, &galeryProps.Title, &galeryProps.Description, &galeryProps.GaleryToken)
    if err != nil {
        if err == sql.ErrNoRows {
            return GaleryProps{}, err
        }
        return GaleryProps{}, err
    }

    return galeryProps, nil
}

func FindProductsByToken(Token string) (EditProduct, error){
    query := "SELECT * FROM products WHERE Token = ?"
    row := db.QueryRow(query, Token)

    var products EditProduct
    err := row.Scan(&products.ID, &products.Token, &products.Name, &products.Price, &products.Description)
    if err != nil {
        if err == sql.ErrNoRows {
            return EditProduct{}, err
        }
        return EditProduct{}, err
    }
    return products, nil
}