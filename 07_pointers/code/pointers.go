package main

// func main() {
// 	var name string
// 	var namePointer *string

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println(&name)
// }

// // ******************************************************

// func main() {
// 	var name string = "Beyonce"
// 	var namePointer *string = &name
// 	var nameValue = *namePointer

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name Value:", nameValue)

// }

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

//Coordinates represent a latitude and longitiude
type Coordinates struct {
	X, Y float64
}

var c = Coordinates{X: 10, Y: 20}

func main() {
  coordinatedMemAdd :=
}
