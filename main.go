package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/roneetkumar/simple/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Simple programming language!\n",
		user.Username)
	fmt.Printf("Simple GO >>\n")
	repl.Start(os.Stdin, os.Stdout)
}
