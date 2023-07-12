package Routers

import (
	"CMS/Middleware"
	"CMS/Handlers"
	"CMS/Handlers/Admin"
	"CMS/Handlers/Blog"
	"CMS/Handlers/ECommerce"
	"CMS/Handlers/Galery"
	"github.com/gofiber/fiber/v2"
)

func Initalize(app *fiber.App) {

	app.Static("/", "./views/public", fiber.Static{
		Compress: false,
		CacheDuration: -1,
	})

	app.Get("/", Handlers.HomePage)
	app.Get("/login", Handlers.LoginPage)
	app.Get("/register", Handlers.RegisterPage)
	app.Get("/blog/:Token", Blog.BlogPage)
	
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
	admin.Get("/edit/html/partials", Admin.EditPartialsHTML)
	admin.Get("/edit/html/blog", Admin.EditBlogPagesHTML)
	admin.Get("/edit/html/forum", Admin.EditForumPagesHTML)
	admin.Get("/edit/html/error", Admin.EditErrorPagesHTML)
	admin.Get("/edit/html/e-commerce", Admin.EditECommercePagesHTML)

	// Blog Properties
	app.Post("/like/blog/:Token", Blog.LikeBlog)
	app.Post("/unlike/blog/:Token", Blog.UnlikeBlog)

	// Normal Properties
	app.Post("/login", Handlers.Login)
	app.Post("/register", Handlers.Register)

	// Admin Properties
	//Create 
	app.Post("/create/blog", Blog.CreateBlog)
	app.Post("/upload/photo", Galery.UploadImage)
	app.Post("/create/product", ECommerce.CreateProduct)
	//Edit Pages
	app.Get("/edit/blog/:Token", Blog.EditBlogPage)
	app.Get("/edit/product/:Token", Admin.EditProductPage)
	app.Get("/edit/user/:Token", Admin.EditUserPage)
	app.Get("/edit/galery/:Token", Admin.EditGaleryPage)
	//Edit
	app.Post("/edit/blog/:Token", Blog.EditBlog)
	app.Post("/edit/user/:Token", Admin.EditUser)
	app.Post("/edit/galery/:Token", Galery.EditGalery)
	app.Post("/edit/product/:Token", ECommerce.EditProduct)
	app.Post("/edit/page", Admin.EditPage)
	//Delete
	app.Delete("/delete/blog/:Token", Blog.DeleteBlog)
	app.Delete("/delete/product/:Token", ECommerce.DeleteProduct)
	app.Delete("/delete/user/:Token", Admin.DeleteUser)
	app.Delete("/delete/galery/:Token", Galery.DeleteGalery)
	app.Delete("/delete/productphoto/:Token", ECommerce.DeleteProductImage)

	app.Use(Middleware.Security)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})
}