package Database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
    "fmt"
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

type Tags struct {
    ID          int
    Token       string
    Tag         string
    BlogToken   string
}

type GaleryItem struct {
	Path        string
	Title       string
	Description string
    Token       string
}

type Products struct {
    ID              int
    Token           string
    Name            string
    Price           int
    Description     string
    ImagePath       string
}

type ProductPhoto struct {
    ID              int
    Token           string
    ProductToken    string
    Path            string
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

func ListAllBlogTags(BlogToken string) ([]Tags, error) {
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        return nil, err
    }
    defer db.Close()

    query := "SELECT * FROM tags WHERE BlogToken=?"
    rows, err := db.Query(query,BlogToken)

    defer rows.Close()

    tags := []Tags{}
    for rows.Next() {
        var tag Tags
        err := rows.Scan(&tag.ID, &tag.Token, &tag.Tag, &tag.BlogToken)
        if err != nil {
            return nil,err
        }
        tags = append(tags, tag)
    }

    if err != nil {
        return nil, err
    }
   
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return tags, nil
}

func ListBlogs(skip,length int) ([]Blogs, error) {
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        return nil, err
    }
    defer db.Close()

    query := "SELECT * FROM blogs LIMIT ?, ?"
    rows, err := db.Query(query,skip,length)
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

func ListUsers(skip, length int) ([]Users, error) {
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        return nil, err
    }
    defer db.Close()

    query := "SELECT * FROM users LIMIT ?, ?"
    rows, err := db.Query(query,skip,length)
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

func ListGaleryItems(skip, length int) ([]GaleryItem, error) {
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        return nil, err
    }
    defer db.Close()

    query := `
    SELECT galery.path, galerypropertys.title, galerypropertys.description, galery.Token
    FROM galery
    JOIN galerypropertys ON galery.Token = galerypropertys.galerytoken LIMIT ?,?
    `

    rows, err := db.Query(query, skip, length)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var galeryItems []GaleryItem
    for rows.Next() {
        var item GaleryItem
        err := rows.Scan(&item.Path, &item.Title, &item.Description, &item.Token)
        if err != nil {
            return nil, err
        }
        galeryItems = append(galeryItems, item)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }
    fmt.Println(galeryItems)
    return galeryItems, nil
}

func ListAllProducts(skip,length int) ([]Products, error){
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        return nil, err
    }
    defer db.Close()

    query := `
    SELECT p.*, (SELECT path FROM productimages WHERE productToken = p.token LIMIT 1) AS path
    FROM products p
    LIMIT ?, ?
    `

    rows, err := db.Query(query,skip,length)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    products := []Products{}
    for rows.Next() {
        var Product Products
        err := rows.Scan(&Product.ID, &Product.Token, &Product.Name, &Product.Price, &Product.Description, &Product.ImagePath)
        if err != nil {
            return nil, err
        }
        products = append(products, Product)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return products, nil
}

func ListAllProductImages(ProductToken string) ([]ProductPhoto, error) {
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/CMS")
    if err != nil {
        return nil, err
    }
    defer db.Close()

    query := "SELECT * FROM productImages WHERE ProductToken = ?"
    rows, err := db.Query(query, ProductToken)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    Images := []ProductPhoto{}
    for rows.Next() {
        var image ProductPhoto
        err := rows.Scan(&image.ID, &image.Token, &image.ProductToken, &image.Path)
        if err != nil {
            return nil, err
        }
        Images = append(Images, image)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return Images, nil
}