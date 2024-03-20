package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, world.")

	handlerFunction := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world\n")
		io.WriteString(w, r.Method)
	}

	http.HandleFunc("/", handlerFunction)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
