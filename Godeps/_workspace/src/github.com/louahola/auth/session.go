package auth

type Session struct {
	UserId int
	Token string `bson:"_id,omitempty"`
}
