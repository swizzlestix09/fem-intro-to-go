package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

type Group struct {
	Role           string
	Users          []User
	NewestUser     User
	SpaceAvailable bool
}

func (u *User) describeUser() string {
	desc := fmt.Sprintf("Name: %s %s E-mail: %s ", u.FirstName, u.LastName, u.Email)
	return desc
}

func (g *Group) describe() string {
	if len(g.Users) > 2 {
		g.SpaceAvailable = false
	}

	desc := fmt.Sprintf("This user group has %b users. The newest user is %s %s. Space Available? %t", len(g.Users), g.NewestUser.FirstName, g.NewestUser.LastName, g.SpaceAvailable)

	return desc
}

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 2, FirstName: "Jane", LastName: "Eyre", Email: "jedizzle@gmail.com"}
	u3 := User{ID: 3, FirstName: "Audrey", LastName: "Hepburn", Email: "aHep01@gmail.com"}

	g := Group{
		Role:           "Admin",
		Users:          []User{u, u2, u3},
		NewestUser:     u2,
		SpaceAvailable: true,
	}

	fmt.Println(u)
	fmt.Println(u.FirstName) //like objects, can access individual properties with dot notation
	fmt.Println(u.describeUser())
	fmt.Println(g.describe())
}
