package dao

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type RawData struct {
	ID        bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt string        `json:"created_at" bson:"created_at"`
	Source    string        `json:"source" bson:"source"`
	Posts     []Post        `json:"posts" bson:"data"`
}

type Post struct {
	UserId uint64 `json:"userId" bson:"user_id"`
	Title  string `json:"title" bson:"title"`
	Body   string `json:"body" bson:"body"`
}
