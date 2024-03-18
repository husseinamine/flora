package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

var UserNotFoundError = errors.New("User Not Found!")

type User struct {
	ID        int       `json:"-"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedOn time.Time `json:"-"`
	UpdatedOn time.Time `json:"-"`
}

type UserList []*User

func (ul *UserList) nextID() int {
	return len(*ul)
}

func (ul *UserList) findByID(id int) (int, error) {
	for i := range *ul {
		if i == id {
			return i, nil
		}
	}

	return -1, UserNotFoundError
}

func (ul *UserList) ToJSON(w io.Writer) error {
	jsonEncoder := json.NewEncoder(w)
	return jsonEncoder.Encode(ul)
}

func (ul *UserList) FromJSON(r io.Reader) (*User, error) {
	user := &User{
		ID:        ul.nextID(),
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
	}

	jsonDecoder := json.NewDecoder(r)
	if err := jsonDecoder.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (ul *UserList) AddUser(user *User) error {
	(*ul) = append(*ul, user)

	return nil
}

func (ul *UserList) UpdateUser(user *User, id int) error {
	index, err := ul.findByID(id)

	if err != nil {
		return err
	}

	(*ul)[index].Username = user.Username
	(*ul)[index].Email = user.Email
	(*ul)[index].Password = user.Password
	(*ul)[index].UpdatedOn = time.Now()

	return nil
}

var Users = UserList{
	{
		0,
		"hussein",
		"contact@husseinamine.com",
		"####",
		time.Now(),
		time.Now(),
	},
}
