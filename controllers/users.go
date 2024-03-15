package controllers

import (
	"encoding/json"
	"io"
	"time"
)

type User struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedOn time.Time `json:"-"`
	UpdatedOn time.Time `json:"-"`
}

type UserList []*User

func (ul *UserList) ToJSON(w io.Writer) error {
	jsonEncoder := json.NewEncoder(w)
	return jsonEncoder.Encode(ul)
}

var Users = UserList{
	{
		"hussein",
		"contact@husseinamine.com",
		"####",
		time.Now(),
		time.Now(),
	},
}
