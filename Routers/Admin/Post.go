package Admin

import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
)

type UserEditForm struct {
	Username string `json:"Username" xml:"Username" form:"Username"`
	Email string `json:"Email" xml:"Email" form:"Email"`
	Perm string `json:"Perm" xml:"Perm" form:"Perm"`
}

func EditUser(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	p := new(UserEditForm)

	userToken := c.Params("Token")
	Token := c.Cookies("Token")

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}

	if User.Perm != "Admin" {
		c.Redirect(redirect)
		return nil
	}
	
	if err := c.BodyParser(p); err != nil {
		return err
	}

	Database.UpdateUser(userToken,p.Username,p.Email,p.Perm)
	c.Redirect(redirect)
	return nil
}