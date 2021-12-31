package management

import "fmt"

type User struct {
	ID       int
	FName    string
	LName    string
	Email    string
	IsActive bool
}

func (user User) Display() string {
	return fmt.Sprint(user)
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func (group Group) DisplayGroup() []User {

	return group.Users
}
