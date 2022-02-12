package model

type User struct {
	ID       int    `json:"id,omitempty"`
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Root     int    `json:"root,omitempty"`
}

type Post struct {
	ID         int             `json:"id,omitempty"`
	Text       string          `json:"text,omitempty"`
	UserID     int             `json:"user_id,omitempty"`
	User       string          `json:"user,omitempty"`
	Date       string          `json:"date,omitempty"`
	Categories []*PostCategory `json:"categories,omitempty"`
	Likes      []*PostLike     `json:"likes,omitempty"`
}

type PostCategory struct {
	ID       int `json:"id,omitempty"`
	Post     int `json:"post,omitempty"`
	Category int `json:"category,omitempty"`
}

type PostLike struct {
	ID   int `json:"id,omitempty"`
	User int `json:"user,omitempty"`
	Post int `json:"post,omitempty"`
	Like int `json:"like,omitempty"`
}

type Category struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Session struct {
	ID    int
	Login string
	UUID  string
	Date  string
}
