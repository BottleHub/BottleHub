// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Chatboard struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	ImageURL    string     `json:"imageURL"`
	Description *string    `json:"description,omitempty"`
	Members     []*User    `json:"members,omitempty"`
	Messages    []*Message `json:"messages,omitempty"`
}

type Comment struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	CommentBy *User  `json:"commentBy"`
	CommentOn *Post  `json:"commentOn"`
}

type FetchChatboard struct {
	ID string `json:"id"`
}

type FetchComment struct {
	ID string `json:"id"`
}

type FetchMessage struct {
	ID string `json:"id"`
}

type FetchPost struct {
	ID string `json:"id"`
}

type FetchUser struct {
	ID string `json:"id"`
}

type Link struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	User    *User  `json:"user"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Message struct {
	ID        string     `json:"id"`
	Text      *string    `json:"text,omitempty"`
	FileURL   *string    `json:"fileURL,omitempty"`
	MessageBy *User      `json:"messageBy"`
	MessageOn *Chatboard `json:"messageOn"`
}

type NewChatboard struct {
	ID          string  `json:"_id"`
	Name        string  `json:"name"`
	ImageURL    string  `json:"imageURL"`
	Description *string `json:"description,omitempty"`
}

type NewComment struct {
	Text      string `json:"text"`
	CommentBy string `json:"commentBy"`
	CommentOn string `json:"commentOn"`
}

type NewLink struct {
	Title   string `json:"title"`
	Address string `json:"address"`
}

type NewMessage struct {
	ID        string  `json:"_id"`
	Text      *string `json:"text,omitempty"`
	FileURL   *string `json:"fileURL,omitempty"`
	MessageBy string  `json:"messageBy"`
	MessageOn string  `json:"messageOn"`
}

type NewPost struct {
	PostedBy    string  `json:"postedBy"`
	ImageURL    string  `json:"imageURL"`
	Description *string `json:"description,omitempty"`
	Likes       int     `json:"likes"`
}

type NewUser struct {
	Username       string  `json:"username"`
	Name           string  `json:"name"`
	About          *string `json:"about,omitempty"`
	Email          string  `json:"email"`
	AvatarImageURL string  `json:"avatarImageURL"`
	Password       string  `json:"password"`
	PublicWallet   string  `json:"publicWallet"`
	PrivateWallet  string  `json:"privateWallet"`
}

type Post struct {
	ID          string     `json:"id"`
	PostedBy    *User      `json:"postedBy"`
	ImageURL    string     `json:"imageURL"`
	Description *string    `json:"description,omitempty"`
	Likes       int        `json:"likes"`
	Comments    []*Comment `json:"comments,omitempty"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type User struct {
	ID             string  `json:"id"`
	Username       string  `json:"username"`
	Name           string  `json:"name"`
	About          *string `json:"about,omitempty"`
	Email          string  `json:"email"`
	AvatarImageURL string  `json:"avatarImageURL"`
	Posts          []*Post `json:"posts,omitempty"`
	Following      []*User `json:"following,omitempty"`
	Followers      []*User `json:"followers,omitempty"`
	PublicWallet   string  `json:"publicWallet"`
	PrivateWallet  string  `json:"privateWallet"`
}

type Status string

const (
	StatusNotStarted Status = "NOT_STARTED"
	StatusInProgress Status = "IN_PROGRESS"
	StatusCompleted  Status = "COMPLETED"
)

var AllStatus = []Status{
	StatusNotStarted,
	StatusInProgress,
	StatusCompleted,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusNotStarted, StatusInProgress, StatusCompleted:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
