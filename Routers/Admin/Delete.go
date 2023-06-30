package Admin

import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
)

func DeleteUser(c *fiber.Ctx) error {
	duserToken := c.Params("Token")
	userToken := c.Cookies("Token")

	redirect := c.Cookies("LastPath")

	User , err := Database.FindUserByToken(userToken)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}

	if User.Perm != "Admin" {
		c.Redirect(redirect)
		return nil
	}

	Database.DeleteUser(duserToken)
	c.Redirect(redirect)
	return nil
}