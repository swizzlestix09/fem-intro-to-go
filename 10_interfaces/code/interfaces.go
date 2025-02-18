package main

import "fmt"

//An interface is list of methods describing behavior of particular types

// Add a Describer interface

type Describer interface { //describer only has one function but I wrote two to satisfy conditions. So for an interface to work, would that mean that I have to write multiple functions with the same name for the structs that are involved with the interface?
	describe() string
}

// User is a single user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// Group is a group of Users
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

// These two structs have different implementations of the `describe()` method.

func (u *User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
	return desc
}

func (g *Group) describe() string {
	desc := fmt.Sprintf("The %s user group has %d users. Newest user:  %s, Accepting New Users:  %t", g.role, len(g.users), g.newestUser.FirstName, g.spaceAvailable)
	return desc
}

// Create a function that doesn't care what type you pass in as long as the type "satisfies the interface"
func DoTheDescribing(d Describer) string {
	return d.describe()
}

func main() {
	u1 := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 1, FirstName: "Humphrey", LastName: "Bogart", Email: "humphrey.bogart@gmail.com"}
	g := Group{role: "admin", users: []User{u1, u2}, newestUser: u2, spaceAvailable: true}

	userDescriptionWInterface := DoTheDescribing(&u1)
	groupDescriptionWInterface := DoTheDescribing(&g)

	fmt.Println(userDescriptionWInterface)
	fmt.Println(groupDescriptionWInterface)
}

// EMPTY interface
/*
An empty interface is written like: interface{}
Specified some kind of struct
Specifies zero methods.
Can hold values of any type - can be used by code that expects an unknown type
Allows you to call methods and functions on types when you arent sure what will be expected
Similar to the any type in typescript
*/
