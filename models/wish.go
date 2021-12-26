package models

import "time"

type Wish struct {
	Content     string    `json:"content"`
	Creator     string    `json:"creator"`
	Hidename    bool      `json:"hidename"`
	CreatedTime time.Time `json:"created_time"`
}

type Wishy struct {
	WishID      string    `json:"wish_id" bson:"wish_id"`
	Content     string    `json:"content" bson:"content"`
	Creator     string    `json:"creator" bson:"creator"`
	CreaterId   string    `json:"creator_id" bson:"creator_id"`
	Hidename    bool      `json:"hidename" bson:"hidename"`
	CreatedTime time.Time `json:"create_at" bson:"created_at"`
}
