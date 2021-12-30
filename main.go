package main

import "fmt"

type User struct {
	ID       int
	FName    string
	LName    string
	Email    string
	IsActive bool
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func (user User) display() string {
	return fmt.Sprint(user)
}

func (group Group) displayGroup() []User {

	return group.Users
}

func main() {
	user1 := User{12, "yoyo", "man", "yman", true}
	user2 := User{11, "ww", "qw", "qq", false}

	userss := []User{user1, user2}

	group := Group{"yoyo", user1, userss, false}

	result := group.displayGroup()

	fmt.Println(result)
}
