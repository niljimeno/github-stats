package server

import (
	"fmt"
	"log"
	"net/http"
)

func Serve() {
	http.HandleFunc("/", generate)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
