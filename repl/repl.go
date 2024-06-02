package repl

import (
	"bufio"
	"fmt"

	"io"
	"os"

	"github.com/om/interpreter/evaluator"
	"github.com/om/interpreter/lexer"
	"github.com/om/interpreter/object"
	"github.com/om/interpreter/parser"
)

const ANGRY_FACE = `【≽ܫ≼】`

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, prompt bool) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	for {
		if prompt {
			fmt.Fprint(out, PROMPT)

		} else {
			fmt.Fprint(out)
		}

		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		
		if line == "exit"||line=="EXIT"||line=="quit"{
			os.Exit(0)
		}
		//if u r making a variable first time u should use variable declaration or short hand declaration 
		//u cannot use = , when defining the variable first time
		// later on to change/assign a new value u can use =
		// next token return the value in token.Token structure

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)

		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, ANGRY_FACE)
	io.WriteString(out, "kuch galat hai\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
