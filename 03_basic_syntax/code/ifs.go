// // Uncomment this entire file

package main

import (
	"errors"
	"fmt"
)

func someFunction() error {
	return errors.New("some error")
}

func main() {

	var someVar = 9

	if someVar > 10 { //parenthesis not required in go - curlies, yes
		fmt.Println(someVar)
	}

	// ****************************

	if someVar > 100 {
		fmt.Println("Greater than 100")
	} else if someVar == 100 {
		fmt.Println("Equals 100")
	} else {
		fmt.Println("Less than 100")
	}

	// ****************************
	err := someFunction() //err represents a error. if func returns error save to var so we can figure it out
	// => If this function returns a value,
	// => it will be an  error of type Error

	// ****************************
	if err != nil {
		fmt.Println(err.Error()) // <-- very common if block use for errors in go
	}

	//set err to result of some function and then check to see if nil.
	//cool thing about this is that err is being created in if block -
	//meaning it is not accessible anywhere else so var name can be used again
	//for other errors

	if err := someFunction(); err != nil {
		fmt.Println(err.Error())
	}

	// End of file curly brace
}
