package main

import "fmt"

func average(grades [5]float64) float64 {
	var total float64
	for _, grade := range grades {
		total += grade
	}
	return total / float64(len(grades))
}

func main() {
	scores := [5]float64{87, 79, 88, 91, 69}
	pets := map[string]string{
		"aimee": "cat",
		"yago":  "cat",
		"puppy": "cat",
		"betty": "dog",
	}

	fmt.Println(average(scores))
}
