package Galery

import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
	"CMS/Utils"
	"image"
)

type GaleryForm struct {
	Title string `json:"Title" xml:"Title" form:"Title"`
	Description string `json:"Description" xml:"Description" form:"Description"`
}

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

	p := new(GaleryForm)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	Database.UploadFile(img, Path, DBPath, p.Title, p.Description)
	c.Redirect(redirect)
	return nil
}

func EditGalery(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	UserToken := c.Cookies("Token")
	GaleryToken := c.Params("Token")
	User, err := Database.FindUserByToken(UserToken)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}

	if User.Perm != "Admin" {
		c.Redirect(redirect)
		return nil
	}

	p := new(GaleryForm)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	Database.UpdateGalery(GaleryToken, p.Title, p.Description)
	c.Redirect(redirect)
	return nil
}