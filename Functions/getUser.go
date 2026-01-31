package main

import "errors"

type User struct {
	Name  string
	Email string
}

// Simulating DB to get user information
func getUser(id int) (User, error) {
	if id == 123 {
		return User{Name: "Timi Aji", Email: "test@gmail.com"}, nil
	}
	return User{}, errors.New("User not found")
}
