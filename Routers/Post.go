package Routers

import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
	"time"
)

type LoginForm struct {
	Email string `json:"email" xml:"email" form:"email"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}

type RegisterForm struct {
	Username string `json:"username" xml:"email" form:"username"`
	Email string `json:"email" xml:"email" form:"email"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}

func Register(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")

	form := new(RegisterForm)
	if err := c.BodyParser(form); err != nil {
		return err
	}

	if !Database.Register(form.Username, form.Email, form.Pass){
		c.Render("register", fiber.Map{
			"title":"Giriş yap",
			"error":"Bu Email Kullanılıyor",
		})
		return nil
	}

	login, _ := Database.Login(form.Email,form.Pass)

	cookie := new(fiber.Cookie)

	cookie.Name = "Token"
	cookie.Value = login
	cookie.Expires = time.Now().Add(time.Hour * 24 * 365)

	c.Cookie(cookie)

	c.Redirect(redirect)
	return nil
}

func Login(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")

	form := new(LoginForm)
	if err := c.BodyParser(form); err != nil {
		return err
	}

	login, err := Database.Login(form.Email,form.Pass)
	if err != nil{
		c.Render("login", fiber.Map{
			"title":"Giriş yap",
			"error":"Email ya da Şifre yanlış",
		})
		return nil
	}
	
	cookie := new(fiber.Cookie)

	cookie.Name = "Token"
	cookie.Value = login
	cookie.Expires = time.Now().Add(time.Hour * 24 * 365)

	c.Cookie(cookie)

	c.Redirect(redirect)
	return nil
}

