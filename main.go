package main

import(
	"github.com/gofiber/fiber/v2"
	"CMS/Routers"
	"CMS/Database"
	"CMS/Routers/Admin"
)

func main(){
	Database.Connect()

	app := fiber.New()
	app.Static("/", "./views/public")

	app.Get("/", Routers.HomePage)
	app.Get("/login", Routers.LoginPage)
	app.Get("/register", Routers.RegisterPage)

	app.Get("/admin/home", Admin.HomePage)

	app.Post("/login", Routers.Login)
	app.Post("/register", Routers.Register)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}