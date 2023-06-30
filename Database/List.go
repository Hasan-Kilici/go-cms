package Database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct{
	ID			int
	Token		string
	Username	string
	Email		string
	Password	string
	Perm		string
}

type Blogs struct {
	ID				int
	Token			string
	Title			string
	HTML			string
	Views			int
	Like			int
}


func ListAllBlogs() ([]Blogs, error) {
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        return nil, err
    }
    defer db.Close()

    query := "SELECT * FROM blogs"
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    blogs := []Blogs{}
    for rows.Next() {
        var blog Blogs
        err := rows.Scan(&blog.ID, &blog.Token,&blog.Title,&blog.HTML,&blog.Views,&blog.Like)
        if err != nil {
            return nil, err
        }
        blogs = append(blogs, blog)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return blogs, nil
}

func ListAllUsers() ([]Users, error) {
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        return nil, err
    }
    defer db.Close()

    query := "SELECT * FROM users"
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    users := []Users{}
    for rows.Next() {
        var user Users
        err := rows.Scan(&user.ID, &user.Token, &user.Username, &user.Email ,&user.Password, &user.Perm)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}