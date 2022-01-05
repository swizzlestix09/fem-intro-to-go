package main

import "fmt"

func main() {
	sentance := "This is my house, bitch!"
	for idx, letter := range sentance {
		if idx%2 != 0 {
			fmt.Println(string(letter))
		}
	}
	//if you are not going to need a variable, in this case index, you can use the _ to notify go that this will not be needed.
	for _, letter := range sentance {
		fmt.Println(string(letter))

	}
}
