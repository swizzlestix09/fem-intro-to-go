package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// A method is a function that is called on an instance of a particular struct. Concerned with state and what state of the struct is being passed into it

//method - paramters after function declaration but in front of name
func (u *User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

//function - parameters after function declaration & name
func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func main() {
	user := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}

	desc := describeUser(user)
	desc2 := user.describe() // <-- how to use a method with an instance.
	fmt.Println(desc)
	fmt.Println((desc2))
}
