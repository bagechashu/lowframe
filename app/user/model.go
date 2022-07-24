package user

import (
	"net/url"
)

var UList []User

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func IsValid(form url.Values) bool {
	username := form["username"][0]
	password := form["password"][0]
	return len(username) > 0 && len(password) > 0
}

func ParseUser(form url.Values) User {
	user := User{
		Username: form["username"][0],
		Password: form["password"][0],
	}
	return user
}
