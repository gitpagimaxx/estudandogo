package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type Page struct {
	Title string
	Body  string
}

func home(w http.ResponseWriter, r *http.Request) {

	u := Page{
		Title: "Home Page",
		Body:  "Welcome to the home page",
	}

	templates.ExecuteTemplate(w, "home.html", u)
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	fmt.Println("Server is running at :5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
