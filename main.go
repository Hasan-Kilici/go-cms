package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"CMS/Database"
	"CMS/Routers"
)

func main() {
	Database.Connect()

	engine := html.New("./views/pages", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	Routers.Initalize(app)
	
	err := app.Listen("127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
}
