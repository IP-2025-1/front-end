package main

import (
	"fmt"
	"front-end/app/handlers"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("../static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/form", handlers.FormHandler)

	fmt.Printf("port running on http://127.0.0.1:3000/\n")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
