package utils

import "fmt"

func printNum(num int) {
	fmt.Println("Current Number:", num)
}

//Add adds multiple numbers, returning total.
func Add(nums ...int) int {
	total := 0
	for _, c := range nums {
		printNum(c)
		total += c
	}
	return total
}
