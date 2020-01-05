package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/roneetkumar/simple/repl"
)

func main() {
	user, err := user.Current()
	userName := strings.Join(strings.Split(user.Username, "\\")[1:], "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Simple programming language!\n",
		userName)
	fmt.Print("Simply GO ")
	repl.Start(os.Stdin, os.Stdout)
}
