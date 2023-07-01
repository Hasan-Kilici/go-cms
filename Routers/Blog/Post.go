package Blog

import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
)

type BlogForm struct {
	Title string `json:"Title" xml:"Title" form:"Title"`
	HTML string `json:"HTML" xml:"HTML" form:"HTML"`
	Tags string `json:"Tags" xml:"Tags" form:"Tags"`
}

func CreateBlog(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	Token := c.Cookies("Token")

	p := new(BlogForm)

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

	Database.CreateBlog(p.Title, p.HTML, p.Tags)
	c.Redirect(redirect)
	return nil
}

func EditBlog(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	blogToken := c.Params("Token")
	userToken := c.Cookies("Token")

	p := new(BlogForm)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	User , err := Database.FindUserByToken(userToken)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}

	if User.Perm != "Admin" {
		c.Redirect(redirect)
		return nil
	}

	Database.UpdateBlog(blogToken,p.Title,p.HTML)
	c.Redirect(redirect)
	return nil
}

func LikeBlog(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	blogToken := c.Params("Token")
	userToken := c.Cookies("Token")

	User , err := Database.FindUserByToken(userToken)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}

	if !Database.FindBlogLike(userToken, blogToken) {
		Blog, err := Database.FindBlogByToken(blogToken)
		if err != nil {
			c.Redirect(redirect)
			return nil
		}
		c.Redirect(redirect)

		NewBlogLikes := Blog.Like + 1
		Database.AddLikeToBlog(Blog.Token, NewBlogLikes)
		Database.SaveUserLike(Blog.Token, User.Token)

		return nil
	}

	c.Redirect(redirect)
	return nil
}

func UnlikeBlog(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	blogToken := c.Params("Token")
	userToken := c.Cookies("Token")

	User , err := Database.FindUserByToken(userToken)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}

	if Database.FindBlogLike(userToken, blogToken) {
		Blog, err := Database.FindBlogByToken(blogToken)
		if err != nil {
			c.Redirect(redirect)
			return nil
		}
		c.Redirect(redirect)

		NewBlogLikes := Blog.Like - 1
		Database.SubLikeToBlog(Blog.Token, NewBlogLikes)
		Database.DeleteUserLike(User.Token, Blog.Token)

		return nil
	}

	c.Redirect(redirect)
	return nil
}