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

func Login(Email, Password string) (string, error) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
	if err != nil {
		return "", err
	}
	defer db.Close()
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
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/cms")
    if err != nil {
        return false
    }
    defer db.Close()

    query := "SELECT COUNT(*) FROM users WHERE Email = ?"
    row := db.QueryRow(query, Email)

    var count int
    err = row.Scan(&count)
    if err != nil {
        return false
    }

    if count == 0 {
        return false
    }

    return true
}

func FindUserByUsername(Username string) bool {
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/cms")
    if err != nil {
        return false
    }
    defer db.Close()

    query := "SELECT COUNT(*) FROM users WHERE username = ?"
    row := db.QueryRow(query, Username)

    var count int
    err = row.Scan(&count)
    if err != nil {
        return false
    }

    if count == 0 {
        return false
    }

    return true
}

func FindUserByToken(Token string) (User, error){
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/cms")
    if err != nil {
        return User{}, err
    }
    defer db.Close()

    query := "SELECT * FROM users WHERE token = ?"
    row := db.QueryRow(query, Token)

    var user User
    err = row.Scan(&user.ID, &user.Token, &user.Username, &user.Email ,&user.Password, &user.Perm)
    if err != nil {
        if err == sql.ErrNoRows {
            return User{}, fmt.Errorf("kullanıcı bulunamadı")
        }
        return User{}, err
    }
    return user, nil
}