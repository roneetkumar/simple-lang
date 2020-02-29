package main

import (
	"fmt"
	"net/http"
	"os/user"

	"html/template"

	"github.com/gorilla/mux"
	"github.com/roneetkumar/simple/evaluator"
	"github.com/roneetkumar/simple/lexer"
	"github.com/roneetkumar/simple/object"
	"github.com/roneetkumar/simple/parser"
)

var templates *template.Template

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Simple programming language!\n",
		user.Username)
	fmt.Printf("Simply GO... \n")

	//PARSING HTML
	templates = template.Must(template.ParseGlob("web/templates/*.html"))

	//PARSING CSS
	fs := http.FileServer(http.Dir("web/templates/assets/"))
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

// Start func
func Start(str string) string {
	env := object.NewEnvironment()

	l := lexer.New(str)
	p := parser.New(l)
	program := p.ParseProgram()

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		return evaluated.Inspect()
	}
	return ""
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	sourceCode := r.FormValue("code")
	output := Start(sourceCode)

	fmt.Println(output)

	templates.ExecuteTemplate(w, "index.html", output)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>About</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact</h1>")
}
