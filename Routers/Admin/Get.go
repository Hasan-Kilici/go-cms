package Admin

import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
	"time"
)

func HomePage(c *fiber.Ctx) error {
	Token := c.Cookies("Token")
	redirect := c.Cookies("LastPath")

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}

	UserInfo := []interface{}{
		User.ID,
		User.Username,
		User.Email,
		User.Perm,
	}

	if(User.Perm != "Admin"){
		c.Redirect(redirect)
		return nil	
	} 

	cookie := new(fiber.Cookie)

	cookie.Name = "LastPath"
	cookie.Value = "/admin/home"
	cookie.Expires = time.Now().Add(time.Hour * 24 * 365)

	c.Cookie(cookie)

	c.Render("admin/index", map[string]interface{}{
		"title":"Admin Dashboard",
		"UserInfo": UserInfo,
	})
	return nil
}

func BlogsPage(c *fiber.Ctx) error{
	Token := c.Cookies("Token")
	redirect := c.Cookies("LastPath")

	Blogs, _ := Database.ListAllBlogs()
	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}

	UserInfo := []interface{}{
		User.ID,
		User.Username,
		User.Email,
		User.Perm,
	}

	if(User.Perm != "Admin"){
		c.Redirect(redirect)
		return nil	
	} 

	cookie := new(fiber.Cookie)

	cookie.Name = "LastPath"
	cookie.Value = "/admin/blogs"
	cookie.Expires = time.Now().Add(time.Hour * 24 * 365)

	c.Cookie(cookie)

	c.Render("admin/blog", fiber.Map{
		"title":"Admin Dashboard",
		"UserInfo": UserInfo,
		"Blogs":Blogs,
	})
	return nil
}

func ForumsPage() {
}

func E_CommercePage(){
}

func UsersPage(c *fiber.Ctx) error {
	Token := c.Cookies("Token")
	redirect := c.Cookies("LastPath")

	Users, _ := Database.ListAllUsers()
	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}

	UserInfo := []interface{}{
		User.ID,
		User.Username,
		User.Email,
		User.Perm,
	}

	if(User.Perm != "Admin"){
		c.Redirect(redirect)
		return nil	
	} 

	cookie := new(fiber.Cookie)

	cookie.Name = "LastPath"
	cookie.Value = "/admin/users"
	cookie.Expires = time.Now().Add(time.Hour * 24 * 365)

	c.Cookie(cookie)

	c.Render("admin/users", fiber.Map{
		"title":"Admin Dashboard",
		"UserInfo": UserInfo,
		"Users":Users,
	},"partials/")
	return nil
}