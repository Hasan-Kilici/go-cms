package Admin

import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
)

func HomePage(c *fiber.Ctx) error {
	Token := c.Cookies("Token")
	redirect := c.Cookies("LastPath")

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}

	if(User.Perm != "Admin"){
		c.Redirect(redirect)
		return nil	
	}

	c.Render("./views/pages/admin/index.html", fiber.Map{
		"title":"Admin Dashboard",
		"Username":User.Username,
		"Email":User.Email,
		"Perm":User.Perm,
	})
	return nil
}