package main

import "fmt"

func main() {
	sentance := "This is my house, bitch!"
	for idx, letter := range sentance {
		if idx%2 != 0 {
			fmt.Println(string(letter))
		}
	}
}
