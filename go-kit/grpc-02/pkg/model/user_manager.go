package model

import (
	"errors"
)

// user
type User struct {
	Name    string
	Title   string
	Company string
}

// user manager
type UserManager struct {
	Users []*User
}

var ErrUserNotFound = errors.New("target user not found")
var ServerAddr = ":8080"

func (m *UserManager) GetUser(name string) (user *User, err error) {
	for _, u := range m.Users {
		if u.Name == name {
			return u, nil
		}
	}
	return nil, ErrUserNotFound
}

func (m *UserManager) Dispatch(name, company string) (err error) {
	for _, u := range m.Users {
		if u.Name == name {
			u.Company = company
			return nil
		}
	}
	return ErrUserNotFound
}

func (m *UserManager) SetTitle(name, title string) (err error) {
	for _, u := range m.Users {
		if u.Name == name {
			u.Title = title
			return nil
		}
	}
	return ErrUserNotFound
}
