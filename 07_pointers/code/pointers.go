package main

import "fmt"

// func main() {
// 	var name string
// 	var namePointer *string

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println(&name)
// }

// // ******************************************************

func main() {
	var name string = "Beyonce"     //name is set to Beyonce
	var namePointer *string = &name //namePointer is pointing to address in memory
	var nameValue = *namePointer    //* is defererencing namePointer to what's held there, which will print Beyonce

	fmt.Println("Name:", name)
	fmt.Println("Name *:", namePointer)
	fmt.Println("Name Value:", nameValue)

}

// // ******************************************************

// func changeName(n *string) {
// 	*n = strings.ToUpper(*n) //the * takes the memory address and produces the value for modification
// }

// func main() {
// 	name := "Elvis"
// 	changeName(&name) //want to send actual address where name is stored to change it properly
// 	fmt.Println(name)
// }

// // ******************************************************

// //Coordinates represent a latitude and longitiude
// type Coordinates struct {
// 	X, Y float64
// }

// var c = Coordinates{X: 10, Y: 20}

// func main() {
// 	coordinatedMemAdd := &c //address location of var c which holds coordinates
// 	coordinatedMemAdd.X = 200
// 	coordinatedMemAdd.Y = 400
// 	fmt.Println(coordinatedMemAdd, c) //both will be changed since the memory address was associated with the change
// 	//there is a shortcut for structs with go were just referencing the struct will change it, like below
// 	newCoordinates := c
// 	newCoordinates.X = 80.12
// 	newCoordinates.Y = 1123.3212
// 	fmt.Println(newCoordinates) //coordinated in c has been changed.
// }
