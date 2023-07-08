package ECommerce

import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
	"CMS/Utils"
	"image"
)

type ProductForm struct {
	Name string `json:"Name" xml:"Name" form:"Name"`
	Description string `json:"Description" xml:"Description" form:"Description"`
	Price int `json:"Price" xml:"Price" form:"Price"`
	Tags string `json:"Tags" xml:"Tags" form:"Tags"`
}


func CreateProduct(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	Token := c.Cookies("Token")

	p := new(ProductForm)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}

	if User.Perm != "Admin" {
		c.Redirect(redirect)
		return nil
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	var fileNames string

	files := form.File["files"]
	for i, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		img, _, err := image.Decode(src)
		if err != nil {
			return err
		}

		fileName := Utils.GenerateToken(40) + ".png"
		err = Utils.SaveResizedImage(img, "./views/public/uploads/"+fileName)
		if err != nil {
			return err
		}

		fileNames += "/uploads/"+fileName

		if i < len(files)-1 {
			fileNames += ","
		}
	}

	Database.CreateProduct(p.Name, p.Description, p.Tags, fileNames, p.Price)
	c.Redirect(redirect)
	return nil
}