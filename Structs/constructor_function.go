package main

/*
	A constructor function is a function that creates and initializes a struct instance.

This can be useful for performing more complex initialization logic or for enforcing certain constraints on the struct's fields.
In Go, we can define a constructor function by convention, usually starting with "New" followed by the struct name.
The constructor function typically returns a pointer to the struct instance.
*/
type User struct {
	Name   string
	Age    int
	Gender string
}

func NewUser(name string, age int, gender string) *User {
	if age < 0 {
		return nil // If age is invalid.
	}
	return &User{
		Name:   name,
		Age:    age,
		Gender: gender,
	}
}
