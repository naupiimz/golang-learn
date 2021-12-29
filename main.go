package main

type User struct {
	ID       int
	FName    string
	LName    string
	Email    string
	IsActive bool
}

func main() {
	user := User{}
	user.ID = 1980
}

func displayUser(user string) string {
	return "vallid"
}
