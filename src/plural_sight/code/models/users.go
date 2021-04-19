package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID int
	FirstName string
	LastName string
}

var (
	users []*User
	nextID = 1
)
func GetUsers() []*User {
	return users	
}
func AddUser(u User) (User, error) {
	if u.ID !=0 {
		return User{} , errors.New("new User must not include id or it must be greater than 1")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found",id)
}
func UpdateUser(u User) (User, error) {
	for i, canidate := range users {
		if canidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found",u.ID)
}
func RemoveUserByID(id int) error {
	for i, u  := range users {
		if u.ID == id {
			users = append(users[:1], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found",id)
}
