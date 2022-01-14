package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

//Planet is a planet type
type Planet struct {
	Name       string `json:"name"`
	Terrain    string `json:"terrain"`
	Population string `json:"population"`
}

//Person is a person type
type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet
}

//Allpeople is a collection of Person types
type AllPeople struct {
	People []Person `json:"results"`
}

func (p *Person) getHomeworld() {
	res, err := http.Get(p.HomeworldURL)

	if err != nil {
		log.Print("error fetching homeworld", err)
	}
	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
		log.Print("Error reading res.Body", err)
	}

	json.Unmarshal(bytes, &p.Homeworld)
}

//Route to get people fmt.fPrint that
func getPeople(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "gettin people")
	res, err := http.Get(BaseURL + "people")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to req star wars people")
	}

	fmt.Println(res)

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("failed to parse request body")
	}

	var people AllPeople

	fmt.Println(string(bytes))
	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("error parsing json", err)
	}

	fmt.Println(people)

	for _, pers := range people.People {
		pers.getHomeworld()
		fmt.Println(pers)
	}
}

func main() {
	fmt.Println(BaseURL)
	http.HandleFunc("/people", getPeople)
	fmt.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
