package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/frankie-mur/monkeylang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome, %q!\n, this is the REPL for monkeylang\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
