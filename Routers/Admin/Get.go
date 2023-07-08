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
		return fiber.ErrUnauthorized
	}

	UserInfo := []interface{}{
		User.ID,
		User.Username,
		User.Email,
		User.Perm,
	}

	if(User.Perm != "Admin"){
		c.Redirect(redirect)
		return fiber.ErrForbidden
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

	Blogs, err := Database.ListBlogs(0,10)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	BlogCount := Database.GetBlogCount()

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrUnauthorized
	}

	UserInfo := []interface{}{
		User.ID,
		User.Username,
		User.Email,
		User.Perm,
	}

	if(User.Perm != "Admin"){
		c.Redirect(redirect)
		return fiber.ErrForbidden
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
	Page, err := strconv.Atoi(c.Params("Page"))
	if err != nil {
		return fiber.ErrBadRequest
	}
	Skip := Page * 10
	Blogs, err := Database.ListBlogs(Skip,10)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	BlogCount := Database.GetBlogCount()

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrUnauthorized
	}

	UserInfo := []interface{}{
		User.ID,
		User.Username,
		User.Email,
		User.Perm,
	}

	if(User.Perm != "Admin"){
		c.Redirect(redirect)
		return fiber.ErrForbidden
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

	Users, err := Database.ListUsers(0,10)
	if err != nil {
		return fiber.ErrInternalServerError 
	}

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrUnauthorized 
	}

	UserInfo := []interface{}{
		User.ID,
		User.Username,
		User.Email,
		User.Perm,
	}

	if(User.Perm != "Admin"){
		c.Redirect(redirect)
		return fiber.ErrForbidden 
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
	Page, err := strconv.Atoi(c.Params("Page"))
	if err != nil {
		return fiber.ErrBadRequest
	}
	Skip := Page * 10

	UserCount := Database.GetUserCount()

	Users, err := Database.ListUsers(Skip,10)
	if err != nil {
		return fiber.ErrInternalServerError 
	}
	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrUnauthorized 
	}

	UserInfo := []interface{}{
		User.ID,
		User.Username,
		User.Email,
		User.Perm,
	}

	if(User.Perm != "Admin"){
		c.Redirect(redirect)
		return fiber.ErrForbidden 
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
		return fiber.ErrUnauthorized 
	}

	if(User.Perm != "Admin"){
		c.Redirect(redirect)
		return fiber.ErrForbidden 
	} 

	EditUser, err := Database.FindUserByToken(UserToken)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrNotFound 
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

func GaleryPage(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	Token := c.Cookies("Token")

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrUnauthorized 
	}

	if(User.Perm != "Admin"){
		c.Redirect(redirect)
		return fiber.ErrForbidden 
	} 

	Galery, err := Database.ListGaleryItems(0,10)
	if err != nil {
		return fiber.ErrInternalServerError 
	}
	GaleryCount := Database.GetGaleryCount()

	cookie := new(fiber.Cookie)

	cookie.Name = "LastPath"
	cookie.Value = "/admin/galery"
	cookie.Expires = time.Now().Add(time.Hour * 24 * 365)

	c.Cookie(cookie)

	c.Render("admin/galery", fiber.Map{
		"title":"Admin galeri",
		"galery":Galery,
		"UserInfo":User,
		"galeryCount":GaleryCount,
	})

	return nil
}

func EditGaleryPage(c *fiber.Ctx) error {
	Token := c.Cookies("Token")
	redirect := c.Cookies("LastPath")
	GaleryToken := c.Params("Token")

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrUnauthorized 
	}

	if(User.Perm != "Admin"){
		c.Redirect(redirect)
		return fiber.ErrForbidden 
	} 

	Galery, err := Database.FindGaleryPropsByToken(GaleryToken)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrNotFound 
	}

	Image, err := Database.FindPhotoByToken(GaleryToken)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrNotFound 
	}

	c.Render("admin/editgalery", fiber.Map{
		"title":"Kullanıcıyı düzenle",
		"Image":Image.Path,
		"GTitle":Galery.Title,
		"Description":Galery.Description,
		"Token":Image.Token,
	})

	return nil
}

func EditGaleryPageWithPages(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	Token := c.Cookies("Token")

	Page, err := strconv.Atoi(c.Params("Page"))
	if err != nil {
		return fiber.ErrBadRequest
	}
	Skip := Page * 10

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrUnauthorized 
	}

	if(User.Perm != "Admin"){
		c.Redirect(redirect)
		return fiber.ErrForbidden
	} 

	Galery, err := Database.ListGaleryItems(Skip,10)
	if err != nil {
		return fiber.ErrInternalServerError 
	}
	GaleryCount := Database.GetGaleryCount()

	cookie := new(fiber.Cookie)

	cookie.Name = "LastPath"
	cookie.Value = "/admin/galery"
	cookie.Expires = time.Now().Add(time.Hour * 24 * 365)

	c.Cookie(cookie)

	c.Render("admin/galery", fiber.Map{
		"title":"Admin galeri",
		"galery":Galery,
		"UserInfo":User,
		"galeryCount":GaleryCount,
	})

	return nil
}

func ProductsPage(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	Token := c.Cookies("Token")

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrUnauthorized 
	}

	if User.Perm != "Admin"{
		c.Redirect(redirect)
		return fiber.ErrForbidden 
	}

	ProductCount := Database.GetProductCount()
	Products , err := Database.ListAllProducts(0,10)
	if err != nil {
		return fiber.ErrInternalServerError 
	}

	UserInfo := []interface{}{
		User.ID,
		User.Username,
		User.Email,
		User.Perm,
	}

	cookie := new(fiber.Cookie)

	cookie.Name = "LastPath"
	cookie.Value = "/admin/products"
	cookie.Expires = time.Now().Add(time.Hour * 24 * 365)

	c.Cookie(cookie)

	c.Render("admin/products", fiber.Map{
		"Title":"Admin Ürünler",
		"Products": Products,
		"UserInfo":UserInfo,
		"ProductCount":ProductCount,
	})
	return nil
}

func ProductsPageWithPages(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	Token := c.Cookies("Token")

	Page, err := strconv.Atoi(c.Params("Page"))
	if err != nil {
		return fiber.ErrBadRequest
	}
	Skip := Page * 10

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrUnauthorized 
	}

	if User.Perm != "Admin"{
		c.Redirect(redirect)
		return fiber.ErrForbidden 
	}

	ProductCount := Database.GetProductCount()
	Products , err := Database.ListAllProducts(Skip,10)
	if err != nil {
		return fiber.ErrInternalServerError 
	}

	UserInfo := []interface{}{
		User.ID,
		User.Username,
		User.Email,
		User.Perm,
	}

	cookie := new(fiber.Cookie)

	cookie.Name = "LastPath"
	cookie.Value = "/admin/products/"+strconv.Itoa(Page)
	cookie.Expires = time.Now().Add(time.Hour * 24 * 365)

	c.Cookie(cookie)

	c.Render("admin/products", fiber.Map{
		"Title":"Admin Ürünler",
		"Products": Products,
		"UserInfo":UserInfo,
		"ProductCount":ProductCount,
	})
	return nil
}

func EditProductPage(c *fiber.Ctx) error {
	redirect := c.Cookies("LastPath")
	ProductToken := c.Params("Token")
	Token := c.Cookies("Token")

	User , err := Database.FindUserByToken(Token)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrUnauthorized 
	}

	if User.Perm != "Admin"{
		c.Redirect(redirect)
		return fiber.ErrForbidden 
	}

	Product, err := Database.FindProductsByToken(ProductToken)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrNotFound 
	}

	ProductImages , err := Database.ListAllProductImages(ProductToken)
	if err != nil {
		c.Redirect(redirect)
		return fiber.ErrNotFound 
	}

	ProductTags, _ := Database.ListAllProductTags(ProductToken)

	cookie := new(fiber.Cookie)

	cookie.Name = "LastPath"
	cookie.Value = "/edit/product/"+ProductToken
	cookie.Expires = time.Now().Add(time.Hour * 24 * 365)

	c.Cookie(cookie)

	c.Render("admin/editproduct", fiber.Map{
		"Title":"Ürün düzenle",
		"Name": Product.Name,
		"Price": Product.Price,
		"Description": Product.Description,
		"PToken": Product.Token,
		"Images": ProductImages,
		"Tags": ProductTags,
	})
	return nil
}