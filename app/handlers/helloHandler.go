package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(response, "404 not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(response, "method is not suported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(response, "Hello You")
}
