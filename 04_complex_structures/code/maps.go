// // Uncomment the entire file
//map in go is like an object in js with key/value pairs
package main

import "fmt"

func main() {
	//type in brackets is the type of key, after it is the type of value.  [int] is key type, string is value type.
	//need to use make to allocate memory for values later.
	//dont need to specify space.

	// var userEmails map[int]string = make(map[int]string)

	// userEmails[1] = "user1@gmail.com"
	// userEmails[2] = "user2@gmail.com"

	// fmt.Println(userEmails)

	// ****************************

	// var userEmails map[int]string = make(map[int]string)
	// // userEmails := make(map[int]string)

	// userEmails[1] = "user1@gmail.com"
	// userEmails[2] = "user2@gmail.com"

	// fmt.Println(userEmails)

	// ****************************
	//can also intiialize map likeso:
	// userEmails := map[int]string{
	// 	1: "user1@gmail.com",
	// 	2: "user2@gmail.com",
	// }

	// fmt.Println(userEmails)

	// 	// ****************************

	// 	userEmails := map[int]string{
	// 		1: "user1@gmail.com",
	// 		2: "user2@gmail.com",
	// 	}

	// 	fmt.Println(userEmails)

	// 	fmt.Println(userEmails[1])

	// 	// ****************************

	// 	userEmails := map[int]string{
	// 		1: "user1@gmail.com",
	// 		2: "user2@gmail.com",
	// 	}

	// 	fmt.Println(userEmails)

	// userEmails[1] = "newUser1@gmail.com"

	// fmt.Println(userEmails)

	// fmt.Println(userEmails[3])

	// 	// ****************************

	// userEmails := map[int]string{
	// 	1: "user1@gmail.com",
	// 	2: "user2@gmail.com",
	// }
	// // can provide second argument to variable that will return a bool letting you know if value is there or not
	// email1, ok := userEmails[1]
	// fmt.Println("Email:", email1, "Present?", ok)

	// email3, ok := userEmails[3]
	// fmt.Println("Email", email3, "Present?", ok)

	// 	// ****************************

	// userEmails := map[int]string{
	// 	1: "user1@gmail.com",
	// 	2: "user2@gmail.com",
	// }

	// //the if will recieve whatever values are assigned to these variables at lookup. email = userEmails[1], ok = ok
	// if email, ok := userEmails[1]; ok {
	// 	fmt.Println(email)
	// } else {
	// 	fmt.Println("I don't know what you want from me")
	// }

	// //can also underscore first value - only checking for the boolean outcome
	// if _, ok := userEmails[1]; ok {
	// 	fmt.Println("exists")
	// } else {
	// 	fmt.Println("I don't know what you want from me")
	// }

	// 	// ****************************

	// userEmails := map[int]string{
	// 	1: "user1@gmail.com",
	// 	2: "user2@gmail.com",
	// }

	// delete(userEmails, 2) //what map are we talking about and which key we want to delete

	// fmt.Println(userEmails)
	// 	// ****************************

	userEmails := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}

	for k, v := range userEmails {
		fmt.Printf("%s has an ID of %d.\n", v, k)
	}
	// 	// ****************************
}
