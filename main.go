package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, world.")

	handlerFunction := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", handlerFunction)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
