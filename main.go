package main

import (
	"log"
	"net/http"
)

func main() {
	// to add : url functionality
	// to add : color functionality
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ColorHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Colors</h1><style>body{background-color: #006400;}</style>")
}

func main() {
    ...
    http.HandleFunc("/color", ColorHandler)
    ...
}
