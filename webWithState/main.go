package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	// _ "github.com/mattn/go-sqlite3"
)

type dataHolder struct {
	Values []string
}

var tmpl *template.Template

func main() {
	var err error
	tmpl, err = template.ParseFiles("templates/main.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", getHandler)
	err = http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, dataHolder{[]string{"foo", "bar"}})
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
}
