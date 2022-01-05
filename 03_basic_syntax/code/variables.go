package main

import "fmt"

// func main() {
// 	var name string = "Angie"
// 	fmt.Println(name) //will log angie
// }

// golang does have inference - though string can be announced can also use an = and will assume its a string if its a simple type
// var name = "Angie" also works

// func main() {
// 	var name = "Angie"
// 	fmt.Println(name) //will log angie
// }

//var name string <-- can do this well. Similar to let; declared a variable and we're holding on to it for use alter. A default value will be assigned as a placeholder.

// func main() {
// 	var name string
// 	name = "Angie"
// 	fmt.Println(name) //will log angie
// }

//can declare a variable outside of function and retain access to it

// var name string = "Angie"

// func main() {
// 	fmt.Println(name) //will log angie
// }

//can assign multiple variables at the same time
// var name1, name2 = "Angie", "Mike"

// func main() {
// 	fmt.Println(name1, name2)
// }

//as long as created in func, can drop var  and type when delaring a variable using := syntax

func main() {
	name := "Danny"
	fmt.Println(name)
}
