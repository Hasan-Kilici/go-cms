package Admin

import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
	"io/ioutil"
)

type UserEditForm struct {
	Username string `json:"Username" xml:"Username" form:"Username"`
	Email string `json:"Email" xml:"Email" form:"Email"`
	Perm string `json:"Perm" xml:"Perm" form:"Perm"`
}

type EditForm struct {
	HTML string `json:"HTML" xml:"HTML" form:"HTML"`
	Path string `json:"Path" xml:"Path" form:"Path"`
}

func EditUser(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	p := new(UserEditForm)

	userToken := c.Params("Token")
	Token := c.Cookies("Token")

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrUnauthorized
	}

	if User.Perm != "Admin" {
		c.Redirect(redirect)
		return fiber.ErrForbidden
	}
	
	if err := c.BodyParser(p); err != nil {
		return fiber.ErrBadRequest
	}

	Database.UpdateUser(userToken,p.Username,p.Email,p.Perm)
	c.Redirect(redirect)
	return nil
}

func EditPage(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	p := new(EditForm)

	Token := c.Cookies("Token")

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrUnauthorized
	}

	if User.Perm != "Admin" {
		c.Redirect(redirect)
		return fiber.ErrForbidden
	}

	if err := c.BodyParser(p); err != nil {
		return err
	}

	err = ioutil.WriteFile(p.Path, []byte(p.HTML), 0644)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}
    
	c.Redirect(redirect)
	return nil
}