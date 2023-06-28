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

	c.Render("index", fiber.Map{
		"title":"Anasayfa",
	})
	return nil
}

func LoginPage(c *fiber.Ctx) error {
	Token := c.Cookies("Token")
	redirect := c.Cookies("LastPath")
	
	if Token != "" {
		c.Redirect(redirect)
		return nil
	}
	
	c.Render("login", fiber.Map{
		"title":"Giriş yap",
	})
	return nil
}

func RegisterPage(c *fiber.Ctx) error {
	Token := c.Cookies("Token")
	redirect := c.Cookies("LastPath")

	if Token != "" {
		c.Redirect(redirect)
		return nil
	}

	c.Render("register", fiber.Map{
		"title":"Giriş yap",
	})
	return nil
}