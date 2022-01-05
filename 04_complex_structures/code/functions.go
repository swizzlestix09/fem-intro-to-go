package main

import "fmt"

//if returning one thing, can use this syntax:

// func checkAge(age int) int {
// 	fmt.Println(age)
// 	return age
// }

//if returning multple things, can announce multiple types i.e:

func checkAge(age int) (int, int) {
	fmt.Println(age)
	return 0, age
}

//can also name return values:

func printAge(age1, age2 int) (ageOfSally, ageOfBob int) {
	//notice that since these variables
	ageOfSally = age1
	ageOfBob = age2
	return
}

func main() {
	checkAge(26)
	// 	x, y := printAge(10, 21)
	// 	fmt.Println(x)
	// 	fmt.Println(y)
}
