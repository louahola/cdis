package auth

type AccountType string
const (
	FACEBOOK = "facebook"
	GOOGLE = "google"
)

type UserType string
const (
	GUEST = "guest"
	MEMBER = "member"
	ADMIN = "admin"
)

type User struct {
	Id int `json:"id" bson:"_id,omitempty"`
	Name string
	Email string
	AccountType AccountType
	UserType UserType
}
