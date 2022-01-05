// // Uncomment the entire file

package main

import "fmt"

func main() {

	var city string

	switch city { //looking at variable declared on line 9
	case "Des Moines":
		fmt.Println("You live in Iowa")
	case "Minneapolis,", "St Paul": // in cases using a comma is like using an or
		fmt.Println("You live in Minnesota")
	case "Madison":
		fmt.Println("You live in Wisconsin")
	default:
		fmt.Println("You're not from around here.")
	}

	// 	// ****************************
	var i int

	switch { //doesnt have a variable to the right - in the case statements we have our variable
	case i > 10:
		fmt.Println("Greater than 10")
	case i < 10:
		fmt.Println("Less than 10")
	default:
		fmt.Println("Is 10")
	}

	// 	// ****************************
	i = 9

	//in go, once a match is found it will not continue to check cases, unless there is a fallthrough declaration.
	switch {
	case i != 10:
		fmt.Println("Does not equal 10")
		fallthrough //our var is set to 9, the first line would return but that fallthrough line will allow go to continue checking cases.
	case i < 10:
		fmt.Println("Less than 10")
	case i > 10:
		fmt.Println("Greater than 10")
	default:
		fmt.Println("Is 10")
	}
}
