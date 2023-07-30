package model

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type User struct {
	Username       string `json:"username"`
	Name           string `json:"name"`
	About          string `json:"about"`
	Email          string `json:"email"`
	AvatarImageURL string `json:"avatarImageURL"`
	Posts          Post   `json:"Posts"`
	Following      int    `json:"following"`
	Follower       int    `json:"follower"`
}

type Post struct {
	id          string
	postedBy    User
	imageURL    string
	description string
	likes       int
	comments    Comment
}

type Comment struct {
	id        string
	text      string
	commentBy User
	commentOn Post
}
