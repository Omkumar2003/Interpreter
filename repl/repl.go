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
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	// If we are in file mode (!prompt), read the whole file at once.
	// Otherwise, for REPL mode, use the line-by-line scanner.
	if !prompt {
		// Read the entire file content.
		// NOTE: io.ReadAll is the modern way to do this.
		content, err := io.ReadAll(in)
		if err != nil {
			fmt.Fprintf(out, "Error reading from input: %s\n", err)
			return
		}

		// Execute the entire file content as one program.
		l := lexer.New(string(content))
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			return
		}

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)
		evaluated := evaluator.Eval(expanded, env)

		// In file mode, we often don't want to print the value of the last
		// expression, only what `puts` prints. For now, we'll keep it
		// simple and print the final value.
		if evaluated != nil && evaluated.Type() != object.NULL_OBJ {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
		return // We are done after executing the file.
	}

	// This is the original REPL loop, which is correct for interactive use.
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" || line == "EXIT" || line == "quit" {
			os.Exit(0)
		}

		l := lexer.New(line)
		p := parser.New(l)
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