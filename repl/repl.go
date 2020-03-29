package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/roneetkumar/simple/evaluator"
	"github.com/roneetkumar/simple/lexer"
	"github.com/roneetkumar/simple/object"
	"github.com/roneetkumar/simple/parser"
)

// PROMPT const
const PROMPT = ">> "

//Start function
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

// Simple const
const SIMPLE_ERROR = `
	INVALID CODE BLOCK
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, SIMPLE_ERROR)
	io.WriteString(out, "Woops! We ran into some simple business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
