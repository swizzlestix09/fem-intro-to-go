package main

import "fmt"

func main() {
	scores := [5]float64{1, 2, 3, 4, 5}

	for _, values := range scores {
		fmt.Println(values)
	}
}

//Slices - segments of underlying array
//make - must be associated with space in memory
