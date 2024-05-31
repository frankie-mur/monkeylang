package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/frankie-mur/monkeylang/evaluator"
	"github.com/frankie-mur/monkeylang/lexer"
	"github.com/frankie-mur/monkeylang/object"
	"github.com/frankie-mur/monkeylang/parser"
)

const PROMPT = ">> "

// Start is the main entry point for the REPL (Read-Eval-Print Loop). It reads input from the provided io.Reader,
// tokenizes the input using the lexer, and prints the resulting tokens to the provided io.Writer.
// The REPL runs in an infinite loop, prompting the user for input and processing it until an error or EOF is encountered.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
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

		evauluated := evaluator.Eval(program, env)
		if evauluated != nil {
			io.WriteString(out, evauluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
