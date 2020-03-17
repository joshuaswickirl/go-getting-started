package models

import (
	"errors"
	"fmt"
)

// User of service
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers returns a slice of pointers to each user
func GetUsers() []*User {
	return users
}

// GetUserByID returns the user with the given ID
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("user with id '%v' not found", id)
}

// AddUser appends the given user to the users slice
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user must not include id")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// UpdateUser overwrites an existing user with the given user by ID
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("user with id '%v' not found", u.ID)
}

// RemoveUserByID splices the given user out of the user slice
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with id '%v' not found", id)
}
