package Routers

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func HomePage(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)

	cookie.Name = "LastPath"
	cookie.Value = "/"
	cookie.Expires = time.Now().Add(time.Hour * 24 * 365)

	c.Cookie(cookie)

	c.Render("./views/pages/index.html", fiber.Map{
		"title":"Anasayfa",
	})
	return nil
}

func LoginPage(c *fiber.Ctx) error {
	c.Render("./views/pages/login.html", fiber.Map{
		"title":"Giriş yap",
	})
	return nil
}

func RegisterPage(c *fiber.Ctx) error {
	c.Render("./views/pages/register.html", fiber.Map{
		"title":"Giriş yap",
	})
	return nil
}