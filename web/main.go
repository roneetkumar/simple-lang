package main

import (
	"fmt"
	"net/http"

	"html/template"

	"github.com/gorilla/mux"
)

var templates *template.Template



func main() {

	//PARSING HTML
	templates = template.Must(template.ParseGlob("templates/*.html"))

	//PARSING CSS
	fs := http.FileServer(http.Dir("templates/assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//CREATING ROUTER INSTANCE
	mux := mux.NewRouter()

	//DEFINING ROUTES
	mux.HandleFunc("/", indexHandler).Methods("GET")
	mux.HandleFunc("/about", aboutHandler).Methods("GET")
	mux.HandleFunc("/contact", contactHandler).Methods("GET")

	http.Handle("/", mux)

	//DEFINING PORTS
	http.ListenAndServe(":80", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// sourceCode := r.FormValue("code")
	// fmt.Println(sourceCode)

	//SOURCE CODE TO BE EXECUTED
	code := `
package main

import(
        "fmt"
)

func main() {
    fmt.Println("Hello, playground")
}`

	templates.ExecuteTemplate(w, "index.html", code)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>About</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact</h1>")
}
