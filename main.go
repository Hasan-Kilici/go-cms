package main

import(
	"github.com/gofiber/fiber/v2"
	"CMS/Routers"
	"CMS/Database"
	"CMS/Routers/Admin"
	"github.com/gofiber/template/html/v2"
)

func main(){
	Database.Connect()
	
	engine := html.New("./views/pages", ".html",)
	app := fiber.New(fiber.Config{
        Views: engine,
    })

	app.Static("/", "./views/public")

	app.Get("/", Routers.HomePage)
	app.Get("/login", Routers.LoginPage)
	app.Get("/register", Routers.RegisterPage)

	app.Get("/admin/home", Admin.HomePage)
	app.Get("/admin/blogs", Admin.BlogsPage)
	app.Get("/admin/users", Admin.UsersPage)

	app.Post("/login", Routers.Login)
	app.Post("/register", Routers.Register)

	err := app.Listen(":4000")
	if err != nil {
		panic(err)
	}
}