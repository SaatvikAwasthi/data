package domain

type Posts []Post

type Post struct {
	UserId uint64 `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
