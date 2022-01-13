package main

import "fmt"

func average(grade1, grade2, grade3 float64) (totalAverage float64) {
	totalAverage = (grade1 + grade2 + grade3) / 3
	return
}

func main() {
	fmt.Println(average(88, 96, 70))
}
