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
	//will return implicitly without having to re write variables
	return
}

//if you're unsure of how many arguments your function will recieve there is a
//spread operator available - bear in mind that since multiple items are being returned, this is now a collection - in order to access items passed one must use a for loop (_ used since index is not needed):
func printAges(ages ...int) {
	for _, value := range ages {
		fmt.Println(value)
	}
	return
}
func main() {
	checkAge(26)
	printAges(14, 17, 43)
	x, y := printAge(10, 21)
	fmt.Println(x)
	fmt.Println(y)
}
