package main

import(
	"github.com/gofiber/fiber/v2"
	"CMS/Routers"
	"CMS/Database"
	"CMS/Routers/Admin"
	"github.com/gofiber/template/html/v2"
	"CMS/Routers/Blog"
)

func main(){
	Database.Connect()
	
	engine := html.New("./views/pages", ".html",)
	app := fiber.New(fiber.Config{
        Views: engine,
    })

	app.Static("/", "./views/public")

	//Content Pages
	app.Get("/", Routers.HomePage)
	app.Get("/login", Routers.LoginPage)
	app.Get("/register", Routers.RegisterPage)
	app.Get("/blog/:Token", Blog.BlogPage)

	//Admin Pages
	app.Get("/admin/home", Admin.HomePage)
	app.Get("/admin/blogs", Admin.BlogsPage)
	app.Get("/admin/users", Admin.UsersPage)

	//Blog Propertys
	app.Post("/like/blog/:Token", Blog.LikeBlog)
	app.Post("/unlike/blog/:Token", Blog.UnlikeBlog)

	//Normal Propertys
	app.Post("/login", Routers.Login)
	app.Post("/register", Routers.Register)

	//Admin Propertys
	app.Post("/create/blog", Blog.CreateBlog)
	app.Post("/delete/blog/:Token", Blog.DeleteBlog)
	app.Get("/edit/blog/:Token", Blog.EditBlogPage) //Blog Edit Page
	app.Post("/edit/blog/:Token", Blog.EditBlog)
	app.Post("/delete/user/:Token", Admin.DeleteUser)
	app.Get("/edit/user/:Token", Admin.EditUserPage) //User Edit Page
	app.Post("/edit/user/:Token", Admin.EditUser)

	err := app.Listen(":4000")
	if err != nil {
		panic(err)
	}
}