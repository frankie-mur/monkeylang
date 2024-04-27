package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/frankie-mur/monkeylang/lexer"
	"github.com/frankie-mur/monkeylang/token"
)

const PROMPT = ">> "

// Start is the main entry point for the REPL (Read-Eval-Print Loop). It reads input from the provided io.Reader,
// tokenizes the input using the lexer, and prints the resulting tokens to the provided io.Writer.
// The REPL runs in an infinite loop, prompting the user for input and processing it until an error or EOF is encountered.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

	}
}
