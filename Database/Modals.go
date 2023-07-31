package Database

type Users struct{
	ID			int
	Token		string
	Username	string
	Email		string
	Password	string
	Perm		string
}

type Blogs struct {
	ID				int
	Token			string
	Title			string
	HTML			string
	Views			int
	Like			int
}

type Tags struct {
    ID              int
    Token           string
    Tag             string
    BlogToken       string
    ProductToken    string
}


type GaleryItem struct {
	Path        string
	Title       string
	Description string
    Token       string
}

type Products struct {
    ID              int
    Token           string
    Name            string
    Price           int
    Description     string
    ImagePath       string
}

type ProductPhoto struct {
    ID              int
    Token           string
    ProductToken    string
    Path            string
}
