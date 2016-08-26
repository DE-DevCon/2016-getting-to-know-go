package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Thank you for visiting %s!", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8082", nil)
}
