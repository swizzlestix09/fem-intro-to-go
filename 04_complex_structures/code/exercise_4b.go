package main

import "fmt"

func average(grades ...float64) float64 {
	var total float64
	for _, grade := range grades {
		total += grade
	}
	return total / float64(len(grades))

}

func main() {
	totalAverage := average(88, 96, 70)
	fmt.Println(totalAverage)
}
