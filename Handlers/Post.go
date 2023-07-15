package Handlers

import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
	"CMS/Utils"
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

	if !Utils.ValidateEmail(form.Email){
		c.Render("register", fiber.Map{
			"title":"Giriş yap",
			"error":"Bu Email geçersiz",
		})
		return nil
	}

	if !Utils.ValidatePassword(form.Pass) {
		c.Render("register", fiber.Map{
			"title":"Giriş yap",
			"error":"Şifre güçlü değil",
		})
		return nil
	}

	if !Utils.ValidateUsername(form.Username) {
		c.Render("register", fiber.Map{
			"title":"Giriş yap",
			"error":"Kullanıcı adı boşluk içeremez",
		})
		return nil
	}

	if !Database.Register(form.Username, form.Email, form.Pass){
		c.Render("register", fiber.Map{
			"title":"Giriş yap",
			"error":"Bu Email Kullanılıyor",
		})
		return fiber.ErrUnauthorized 
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
	user := Database.FindUserByEmail(form.Email)
	if !user {
		c.Render("login", fiber.Map{
			"title":"Giriş yap",
			"error":"Kullanıcı bulunamadı",
		})
		return fiber.ErrNotFound
	}

	login, err := Database.Login(form.Email,form.Pass)
	if err != nil{
		c.Render("login", fiber.Map{
			"title":"Giriş yap",
			"error":"Email ya da Şifre yanlış",
		})
		return fiber.ErrUnauthorized
	}
	
	cookie := new(fiber.Cookie)

	cookie.Name = "Token"
	cookie.Value = login
	cookie.Expires = time.Now().Add(time.Hour * 24 * 365)

	c.Cookie(cookie)

	c.Redirect(redirect)
	return nil
}

