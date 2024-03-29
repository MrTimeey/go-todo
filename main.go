package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type ToDo struct {
	Id    string
	Title string
	Done  bool
}

var tmpl = template.Must(template.ParseFiles("index.html"))

func main() {
	templateFunction := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		todos := map[string][]ToDo{
			"Todos": {
				{Id: "3f665f1a-64bd-4507-ace7-a10ebc1b7da9", Title: "Finish project proposal", Done: false},
				{Id: "96e9c5e3-1d81-4a43-9cb5-39e596a5b348", Title: "Buy groceries", Done: false},
				{Id: "ce3270c7-e1fe-481d-9eb8-04938dd9d395", Title: "Go for a run", Done: true},
			},
		}
		tmpl.Execute(w, todos)
	}
	addFunction := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		tmpl.ExecuteTemplate(w, "todo-list-element", ToDo{Id: uuid.New().String(), Title: title, Done: false})
	}
	checkFunction := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		currentState, _ := strconv.ParseBool(r.PostFormValue("currentState"))
		id := r.PathValue("id")
		tmpl.ExecuteTemplate(w, "todo-list-element", ToDo{Id: id, Title: title, Done: !currentState})
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/*", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", templateFunction)
	http.HandleFunc("POST /add-todo/", addFunction)
	http.HandleFunc("POST /check-todo/{id}", checkFunction)

	fmt.Println("Server up and running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
