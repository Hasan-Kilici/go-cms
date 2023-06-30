package Blog

import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
)

func DeleteBlog(c *fiber.Ctx) error {
	blogToken := c.Params("Token")
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

	Database.DeleteBlog(blogToken)
	c.Redirect(redirect)
	return nil
}