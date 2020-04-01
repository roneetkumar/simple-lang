package main

import (
	"fmt"
	"net/http"
	"os"
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
	mux.HandleFunc("/doc", docHandler).Methods("GET")

	http.Handle("/", mux)

	//DEFINING PORTS
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	http.ListenAndServe(":"+port, nil)

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

	code := r.FormValue("code")
	output := Start(code)

	data := struct {
		Code, Output string
	}{
		Code:   code,
		Output: output,
	}

	fmt.Println(data.Output)

	templates.ExecuteTemplate(w, "index.html", data)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>About</h1>")
}

func docHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "doc.html", nil)
}
