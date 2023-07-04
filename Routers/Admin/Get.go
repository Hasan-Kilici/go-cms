package Admin

import (
	"github.com/gofiber/fiber/v2"
	"CMS/Database"
	"time"
	"strconv"
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

	Blogs, _ := Database.ListBlogs(0,10)
	BlogCount := Database.GetBlogCount()

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
		"BlogCount":BlogCount,
	})
	return nil
}

func BlogsPageWithPages(c *fiber.Ctx) error{
	Token := c.Cookies("Token")
	redirect := c.Cookies("LastPath")
	Page, _ := strconv.Atoi(c.Params("Page"))
	Skip := Page * 10
	Blogs, _ := Database.ListBlogs(Skip,10)
	BlogCount := Database.GetBlogCount()

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
		"BlogCount":BlogCount,
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
	UserCount := Database.GetUserCount()

	Users, _ := Database.ListUsers(0,10)
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
		"UserCount":UserCount,
	})
	return nil
}

func UsersPageWithPages(c *fiber.Ctx) error {
	Token := c.Cookies("Token")
	redirect := c.Cookies("LastPath")
	Page, _ := strconv.Atoi(c.Params("Page"))
	Skip := Page * 10

	UserCount := Database.GetUserCount()

	Users, _ := Database.ListUsers(Skip,10)
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
		"UserCount":UserCount,
	})
	return nil
}


func EditUserPage(c *fiber.Ctx) error {
	Token := c.Cookies("Token")
	redirect := c.Cookies("LastPath")
	UserToken := c.Params("Token")

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}

	if(User.Perm != "Admin"){
		c.Redirect(redirect)
		return nil	
	} 

	EditUser, err := Database.FindUserByToken(UserToken)
	if err != nil {
		c.Redirect(redirect)
		return nil
	}

	c.Render("admin/edituser", fiber.Map{
		"title":"Kullanıcıyı düzenle",
		"Username": EditUser.Username,
		"Email": EditUser.Email,
		"Password":EditUser.Password,
		"Perm":EditUser.Perm,
		"Token":EditUser.Token,
	})

	return nil
}