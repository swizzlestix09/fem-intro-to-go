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

type Todo struct {
	Title   string
	Content string
}

type PageVariables struct {
	PageTitle string
	PageTodos []Todo
}

var todos []Todo
/Users/swizzlestix/Coinbase/fem-intro-to-go/08_errors
func getTodos(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		PageTitle: "Get Todos",
		PageTodos: todos,
	}

	t, err := template.ParseFiles("todos.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template error", err)
	}
	err = t.Execute(w, pageVariables)
}

func addTodo(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", getTodos)
	http.HandleFunc("/add-todos", addTodos)
	fmt.Println("Server is running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
