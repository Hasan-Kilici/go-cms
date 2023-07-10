package Galery

import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
)

func DeleteGalery(c *fiber.Ctx) error{
	GaleryToken := c.Params("Token")
	UserToken := c.Cookies("Token")
	redirect := c.Cookies("LastPath")

	User, err := Database.FindUserByToken(UserToken)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}

	if User.Perm != "Admin" {
		c.Redirect(redirect)
		return nil	
	}

	Database.DeleteGalery(GaleryToken)
	c.Redirect(redirect)

	return nil
}