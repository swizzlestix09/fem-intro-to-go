package main

import "fmt"

func main() {

	var myArray [5]int //set to 5 element array of ints
	// var mySlice []int  //doesnt have length - a 'slice' - go will not know how much memory to allocate
	var mySlice []int = make([]int, 5, 10) //we dont know whats going to be in the array yet but carve out 5 spaces in memory for it
	//first argument is type of array, second is how many values to initialize with, optional third argument is max amount of values in array

	myArray[0] = 1
	mySlice[0] = 1 //because there is not memory allocated to this, we will run into an error. index[0] out of range[0] with length 0

	fmt.Println(myArray) //
	fmt.Println(mySlice)

	// ***************************

	// var myArray [5]int
	// var mySlice []int = make([]int, 5)

	// fmt.Println(myArray)
	// fmt.Println(mySlice)

	// ***************************

	// var myArray [5]int
	// // var mySlice []int = make([]int, 5)
	// var mySlice []int = make([]int, 5, 10)
	// // var mySlice = make([]int, 5, 10)

	// myArray[0] = 1
	// mySlice[0] = 1

	// fmt.Println(myArray)
	// fmt.Println(mySlice)
	// fmt.Println(len(mySlice))
	// fmt.Println(cap(mySlice))

	// ***************************

	fruitArray := [5]string{"banana", "pear", "apple", "kumquat", "peach"}

	var splicedFruit []string = fruitArray[1:3]        // ==> ["pear", "apple",]
	var anotherSplicedFruit []string = fruitArray[2:5] // ==>
	// in order to slice and store from index 1 to 3 of fruitArray, you would create a new variable and declare it of array type, then set it to the array you'd like to copy and in square brackets announce what you want to have.
	fmt.Println(splicedFruit)
	fmt.Println(splicedFruit)
	fmt.Println(anotherSplicedFruit)

	fmt.Println(len(splicedFruit))
	fmt.Println(cap(splicedFruit))
	fmt.Println(cap(anotherSplicedFruit))

	fruitToAdd := append(fruitArray[:], "cherries", "pomegrante", "watermelon")
	fmt.Println("After appending", fruitToAdd)
	fmt.Println(len(fruitToAdd))
	//appending to array doubles array size. go anticipates you may need more space, so it allocates space for any extra additions.
	fmt.Println(cap(fruitToAdd))
	// ***************************

	// SEE SLIDE

	// ***************************

	slice1 := []int{1, 2, 3}
	//append is similar to push
	slice2 := append(slice1, 4, 5)

	fmt.Println(slice1, slice2)
	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(len(slice2), cap(slice2))

	// ***************************

	// originalSlice := []int{1, 2, 3}
	// destination := make([]int, len(originalSlice))

	// fmt.Println("Before Copy:", originalSlice, destination)

	// mysteryValue := copy(destination, originalSlice)

	// // fmt.Println("After Copy:", originalSlice, destination, mysteryValue)
}
