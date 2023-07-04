package Galery

import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
	"CMS/Utils"
	"image"
)

func UploadImage(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	Token := c.Cookies("Token")
	User, err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}
	if User.Perm != "Admin" {
		c.Redirect(redirect)
		return nil
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		return err
	}

	FileName := Utils.GenerateToken(40)
	Path := "./views/public/uploads/" + FileName + ".png"
	DBPath := "/uploads/" + FileName + ".png"

	Database.UploadFile(img, Path, DBPath)

	c.Redirect(redirect)
	return nil
}