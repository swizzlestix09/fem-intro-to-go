package main

import "fmt"

func average(grades [5]float64) float64 {
	var total float64
	for _, grade := range grades {
		total += grade
	}
	return total / float64(len(grades))
}

var pets = map[string]string{
	"aimee": "cat",
	"yago":  "cat",
	"puppy": "cat",
	"betty": "dog",
}

func petValidation(pet string) bool {
	_, isValid := pets[pet]
	return isValid
}

var groceries = []string{"corn", "cookies", "salad", "onions", "sirloin", "grapes", "weed"}

func addToGroceries(foods ...string) []string {
	listOfFoods := groceries
	for _, food := range foods {
		listOfFoods = append(groceries, food)
	}
	return listOfFoods

}

func main() {
	scores := [5]float64{87, 79, 88, 91, 69}

	fmt.Println(average(scores))
	fmt.Println(petValidation("aimee"))
	fmt.Println(addToGroceries("Kiwi", "Pinapples", "mamahuevos"))
}
