package Tables

type Blog struct {
	ID				int
	Token			string
	Url				string
	Title			string
	Description		string
	LikeCount		string
}

type BlogLikes struct {
	ID				int
	Token			string
	BlogToken		string
	UserToken		string
}