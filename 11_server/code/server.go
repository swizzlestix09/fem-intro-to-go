package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//http.ResponseWriter is a built-in type from http library
//http.Request is another builtin type that is a pointer
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
}

func todos(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("todos.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template error", err)
	}
	err = t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", todos)
	fmt.Println("Server is running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
