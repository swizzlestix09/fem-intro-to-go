// // Uncomment the entire file

package main

import "fmt"

func main() {

	// 	// ****************************

	// i := 1

	// for i := 1; i <= 100; i++ {
	// 	fmt.Println(i)
	// }

	// 	// 	// ****************************

	// i := 1

	// for i <= 100 { //will behave like a while loop
	// 	fmt.Println(i)
	// 	// This will behave like a while loop
	// 	i += 1
	// }

	// 	// 	// ****************************

	//this for loop is using the range keyword
	//for loop has some variables that are getting assigned
	//in the for loop we are iterating over each letter in the string
	//index is printing out, and each letter is a number. returning bytes. type conversion would be handy here to reconvert to a letter
	//range keyword - indeicated we want to iterate over something
	//want to for each my way around a particular thing
	var mySentence = "This is a sentence"

	for index, letter := range mySentence {
		fmt.Println("Index:", index, "Letter:", letter, string(letter))
	}
}
