package Database

import (
	"CMS/Utils"
	"fmt"
	b64 "encoding/base64"
	"strings"
	"image"
)

func Register(Username, Email, Password string) bool {	
	if !FindUserByEmail(Email){
		Token := Utils.GenerateToken(31)
		NewPassword := b64.URLEncoding.EncodeToString([]byte(Password))

		res, err := db.Exec("INSERT INTO Users VALUES (?,?,?,?,?,?)","",Token,Username,Email,NewPassword,"User")
		if err != nil {
			fmt.Println(err)
			return false
		}
		
		rowCount, err := res.RowsAffected()
		if err != nil {
			fmt.Println(err)
			return false
		}
		
		fmt.Printf("%d satır eklendi\n", rowCount)
		return true
	} else {
		fmt.Println("Bu Email zaten Kullanılıyor")
		return false
	}
}

func CreateBlog(Title, HTML , Tags string){
	Token := Utils.GenerateToken(31)
	res, err := db.Exec("INSERT INTO Blogs VALUES (?,?,?,?,?,?)","",Token,Title,HTML,0,0)
	if err != nil {
		fmt.Println(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	SaveAllBlogTags(Token, Tags)
	
	fmt.Printf("%d satır eklendi\n", rowCount)
}

func SaveUserLike(BlogToken, UserToken string) {
	Token := Utils.GenerateToken(31)
	res, err := db.Exec("INSERT INTO likes VALUES (?,?,?,?)","",Token,UserToken,BlogToken)
	if err != nil {
		fmt.Println(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
		
	fmt.Printf("%d satır eklendi\n", rowCount)
}

func SaveAllBlogTags(BlogToken, Tag string) {
	Tags := strings.Split(Tag, ",")
	TagCount := len(Tags)
	for i := 0;i < TagCount;i++ {
		Token := Utils.GenerateToken(31)
		res, err := db.Exec("INSERT INTO tags VALUES (?,?,?,?,?)","",Token,Tags[i],BlogToken,"")
		if err != nil {
			fmt.Println(err)
		}

		rowCount, err := res.RowsAffected()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d satır eklendi\n", rowCount)
	}
} 

func UploadFile(img image.Image, path , dbpath , title, description string){
	Utils.SaveResizedImage(img, path)

	Token := Utils.GenerateToken(31)
	res, err := db.Exec("INSERT INTO galery VALUES (?,?,?)","",Token,dbpath)
	if err != nil {
		fmt.Println(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	
	SaveImageProps(Token,title,description)
	fmt.Printf("%d satır eklendi\n", rowCount)
}

func SaveImageProps(galeryToken, title, description string) {
	Token := Utils.GenerateToken(31)
	res, err := db.Exec("INSERT INTO galerypropertys VALUES (?,?,?,?,?)","",Token,title,description,galeryToken)
	if err != nil {
		fmt.Println(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("%d satır eklendi\n", rowCount)
}


func SaveAllProductTags(ProductToken, Tag string) {
	Tags := strings.Split(Tag, ",")
	TagCount := len(Tags)
	for i := 0;i < TagCount;i++ {
		Token := Utils.GenerateToken(31)
		res, err := db.Exec("INSERT INTO tags VALUES (?,?,?,?,?)","",Token,Tags[i],"",ProductToken)
		if err != nil {
			fmt.Println(err)
		}

		rowCount, err := res.RowsAffected()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d satır eklendi\n", rowCount)
	}
} 

func SaveAllProductImages(ProductToken, ImagePaths string) {
	Images := strings.Split(ImagePaths, ",")
	ImagesCount:= len(Images)

	for i := 0;i < ImagesCount;i++ {
		Token := Utils.GenerateToken(31)
		res, err := db.Exec("INSERT INTO productImages VALUES (?,?,?,?)","",Token,ProductToken,Images[i])
		if err != nil {
			fmt.Println(err)
		}

		rowCount, err := res.RowsAffected()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d satır eklendi\n", rowCount)
	}
}

func CreateProduct(Name, Description, Tags, imagePaths string, Price int){
	Token := Utils.GenerateToken(31)
	res, err := db.Exec("INSERT INTO products VALUES (?,?,?,?,?)","",Token,Name,Price,Description)
	if err != nil {
		fmt.Println(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	
	SaveAllProductTags(Token, Tags)
	SaveAllProductImages(Token, imagePaths)
	fmt.Printf("%d satır eklendi\n", rowCount)
}
