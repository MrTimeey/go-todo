package main

import (
	"fmt"
	"github.com/google/uuid"
	"html/template"
	"log"
	"net/http"
	"time"
)

type ToDo struct {
	Id    string
	Title string
	Done  bool
}

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
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "todo-list-element", ToDo{Id: uuid.New().String(), Title: title, Done: false})
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/*", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", templateFunction)
	http.HandleFunc("POST /add-todo/", addFunction)

	fmt.Println("Server up and running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
