package user

import (
	"net/url"
)

var UList []User

type User struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

func IsValid(form url.Values) bool {
	fname := form["fname"][0]
	lname := form["lname"][0]
	return len(fname) > 0 && len(lname) > 0
}

func ParseUser(form url.Values) User {
	user := User{
		Fname: form["fname"][0],
		Lname: form["lname"][0],
	}
	return user
}
