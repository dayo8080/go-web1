package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("The server is running on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)

}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")

}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}


func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplated, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplated.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates:", err)
	}
}