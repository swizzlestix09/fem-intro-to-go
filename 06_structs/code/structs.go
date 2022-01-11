package main

import "fmt"

// for struct to be available in other packages ( imports ) struct &  properties should be uppercase

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// // User is a user type
// type User struct {
// ID					       int
// FirstName, LastName, Email string
// }

//Group represents a group of users
type Group struct {
	Role           string
	Users          []User
	NewestUser     User
	SpaceAvailable bool
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s E-mail: %s ", u.FirstName, u.LastName, u.Email)
	return desc
}

func describeGroup(g Group) string {
	desc := fmt.Sprintf("This user group has %b users. The newest user is %s %s. Space Available? %t", len(g.Users), g.NewestUser.FirstName, g.NewestUser.LastName, g.SpaceAvailable)
	return desc
}

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 2, FirstName: "Jane", LastName: "Eyre", Email: "jedizzle@gmail.com"}

	g := Group{
		Role:           "Admin",
		Users:          []User{u, u2},
		NewestUser:     u2,
		SpaceAvailable: true,
	}

	fmt.Println(u)
	fmt.Println(u.FirstName) //like objects, can access individual properties with dot notation
	fmt.Println(describeUser(u))
	fmt.Println(describeGroup(g))
}
