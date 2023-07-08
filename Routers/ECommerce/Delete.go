package ECommerce
import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
)

func DeleteProduct(c *fiber.Ctx) error {
	productToken := c.Params("Token")
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

	Database.DeleteProduct(productToken)
	c.Redirect(redirect)
	return nil
}