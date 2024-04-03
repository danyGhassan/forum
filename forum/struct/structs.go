package structs

type User struct {
	Uuid     string
	ID       int
	Username string
	Email    string
}

type Posts struct {
	ID       int
	Content  string
	Category string
	Likes    int
	Dislikes int
	IdUsers  int
	Comments []Comments
}

type Comments struct {
	ID       int
	IDPosts  int
	Content  string
	Likes    int
	Dislikes int
	IdUsers  int
}

type PageData struct {
	Users    []User
	Posts    []Posts
	Comments []Comments
}

var Connected string
var Connect bool
var IdConnected = 0
