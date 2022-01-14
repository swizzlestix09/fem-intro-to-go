package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//http.ResponseWriter is a built-in type from http library
//http.Request is another builtin type that is a pointer
// func home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Home")
// }

type Todo struct {
	Title   string
	Content string
}

type PageVariables struct {
	PageTitle string
	PageTodos []Todo
}

var todos []Todo

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

func addTodos(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // parses through request
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Request parsing error: ", err)
	}

	todo := Todo{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	todos = append(todos, todo)
	log.Print(todos)
	http.Redirect(w, r, "/todos/", http.StatusSeeOther)

}

func main() {
	//http.HandleFunc("/", home)
	http.HandleFunc("/todos/", getTodos)
	http.HandleFunc("/add-todo/", addTodos)
	fmt.Println("Server is running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
