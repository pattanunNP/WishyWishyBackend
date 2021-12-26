package models

type LineProfile struct {
	DisplayName   string `json:"displayName"`
	UserID        string `json:"userId"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}
type User struct {
	WishID string `json:"wish_id" bson:"wish_id"`
}

type UserWishs struct {
	UserID string `json:"user_id" bson:"user_id"`
	Wishs  string `json:"wishs" bson:"wishs"`
}
