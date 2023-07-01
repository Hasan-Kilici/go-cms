package Blog

import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
	"time"
)

func EditBlogPage(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	blogToken := c.Params("Token")
	userToken := c.Cookies("Token")

	User , err := Database.FindUserByToken(userToken)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}

	if User.Perm != "Admin" {
		c.Redirect(redirect)
		return nil
	}

	Blog, _ := Database.FindBlogByToken(blogToken)

	c.Render("admin/editblog", fiber.Map{
		"ID":Blog.ID,
		"BTitle":Blog.Title,
		"Token":Blog.Token,
		"HTML":Blog.HTML,
		"Views":Blog.Views,
		"Like":Blog.Like,
		"Title":"Blogu düzenle",
	})
	return nil
}

func BlogPage(c *fiber.Ctx) error {
	blogToken := c.Params("Token")
	userToken := c.Cookies("Token")
	Blog, err := Database.FindBlogByToken(blogToken)
	Tags , _ := Database.ListAllBlogTags(blogToken)
	if err != nil {
		c.Render("error/404",fiber.Map{
			"Title":"Sayfa bulunamadı",
		})
		return nil
	}
	
	UserLike := Database.FindBlogLike(userToken,blogToken)
	
	User , err := Database.FindUserByToken(userToken)
	if err != nil {
		c.Render("blog/blog", fiber.Map{
			"UserStatus":false,
			"ID":Blog.ID,
			"Title":Blog.Title,
			"Token":Blog.Token,
			"HTML":Blog.HTML,
			"Views":Blog.Views,
			"Like":Blog.Like,
			"Tags":Tags,
		})
		return nil
	}
	
	c.Render("blog/blog", fiber.Map{
		"UserStatus":true,
		"UserLike": UserLike,
		"Username":User.Username,
		"UserToken":User.Token,
		"Email": User.Email,
		"ID":Blog.ID,
		"Title":Blog.Title,
		"Token":Blog.Token,
		"HTML":Blog.HTML,
		"Views":Blog.Views,
		"Like":Blog.Like,
		"Tags":Tags,
	})

	cookie := new(fiber.Cookie)

	cookie.Name = "LastPath"
	cookie.Value = "/blog/"+Blog.Token
	cookie.Expires = time.Now().Add(time.Hour * 24 * 365)

	c.Cookie(cookie)

	NewView := Blog.Views + 1
	Database.AddViewToBlog(Blog.Token, NewView)
	return nil
}