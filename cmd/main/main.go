package main

import "fmt"

type User struct {
	name string // private attribute of user struct
}

// importers can call using user.New()
func New(name string) *User {
	return &User{
		name: name,
	}
}

// using the suffix `Interface` here
type UserInterface interface {
	Name() string
}

// public function for user
func (u *User) Name() string {
	return "Name: " + u.name
}

// in main.go - to make this code work change package name to main instead
func main() {
	userA := New("kishen")
	fmt.Println(userA.Name())
}
