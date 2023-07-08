package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"CMS/Database"
	"CMS/Routers"
	"CMS/Routers/Admin"
	"CMS/Routers/Blog"
	"CMS/Routers/Galery"
	"CMS/Routers/ECommerce"
)

func main() {
	Database.Connect()

	engine := html.New("./views/pages", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./views/public", fiber.Static{
		CacheDuration: 1 * time.Second,
	})

	// Content Pages
	app.Get("/", Routers.HomePage)
	app.Get("/login", Routers.LoginPage)
	app.Get("/register", Routers.RegisterPage)
	app.Get("/blog/:Token", Blog.BlogPage)

	// Admin Pages
	admin := app.Group("/admin")
	admin.Get("/home", Admin.HomePage)
	admin.Get("/blogs", Admin.BlogsPage)
	admin.Get("/users", Admin.UsersPage)
	admin.Get("/products", Admin.ProductsPage)
	admin.Get("/galery", Admin.GaleryPage)
	admin.Get("/blogs/:Page", Admin.BlogsPageWithPages)
	admin.Get("/users/:Page", Admin.UsersPageWithPages)
	admin.Get("/galery/:Page", Admin.EditGaleryPageWithPages)
	admin.Get("/products/:Page", Admin.ProductsPageWithPages)

	// Blog Properties
	app.Post("/like/blog/:Token", Blog.LikeBlog)
	app.Post("/unlike/blog/:Token", Blog.UnlikeBlog)

	// Normal Properties
	app.Post("/login", Routers.Login)
	app.Post("/register", Routers.Register)

	// Admin Properties
	app.Post("/create/blog", Blog.CreateBlog)
	app.Post("/upload/photo", Galery.UploadImage)
	app.Post("/create/product", ECommerce.CreateProduct)

	app.Get("/edit/blog/:Token", Blog.EditBlogPage)
	app.Get("/edit/product/:Token", Admin.EditProductPage)
	app.Get("/edit/user/:Token", Admin.EditUserPage)
	app.Get("/edit/galery/:Token", Admin.EditGaleryPage)

	app.Post("/edit/blog/:Token", Blog.EditBlog)
	app.Post("/edit/user/:Token", Admin.EditUser)
	app.Post("/edit/galery/:Token", Galery.EditGalery)
	app.Post("/edit/product/:Token", ECommerce.EditProduct)

	app.Delete("/delete/blog/:Token", Blog.DeleteBlog)
	app.Delete("/delete/product/:Token", ECommerce.DeleteProduct)
	app.Delete("/delete/user/:Token", Admin.DeleteUser)
	app.Delete("/delete/galery/:Token", Galery.DeleteGalery)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
