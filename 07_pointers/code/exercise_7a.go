package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

func updateEmail(user *User) {
	user.Email = "DollyTheQueen@yahoo.com"
	return
}

func main() {
	u := User{ID: 1, FirstName: "Dolly", LastName: "Parton", Email: "notJolene@gmail.com"}

	fmt.Println("Before Changes: ", u)
	updateEmail(&u)
	fmt.Println("After Changes:  ", u)

}
