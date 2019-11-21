package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"calcwithantlr/calc"
	"calcwithantlr/parser"
)

func main() {
	// Setup channels to quit cleanly
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	// Listen for OS signals
	signal.Notify(quit, os.Interrupt)

	go shutdown(quit, done)

	go repl()

	<-done
}

// setup REPL.
func repl() {
	// defer is always executed at end of routine
	// recover from errors and go back into loop
	// if no error, do nothing and exit
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("# Invalid input ...\n")
			repl()
		}
	}()

	// read input string from prompt
	get := func(r *bufio.Reader) string {
		t, _ := r.ReadString('\n')
		return strings.TrimSpace(t)
	}

	// wait to read from cli
	fmt.Printf("> ")
	reader := bufio.NewReader(os.Stdin)
	text := get(reader)

	// parse and calculate/present response
	res := calculate(text)
	fmt.Printf("# %s\n", res.Simplify())

	// go back into repl
	repl()
}

// this routine blocks at first line until an OS interrupt signal is read
// then closes the done channel so main can continue and finish
func shutdown(quit <-chan os.Signal, done chan<- bool) {
	<-quit
	close(done)
}

func calculate(in string) *calc.Numby {
	// Create input stream
	is := antlr.NewInputStream(in)

	// Create the Lexer
	lexer := parser.NewMixedLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	parser := parser.NewMixedParser(stream)

	// Finally parse the expression
	var listener calc.Listener
	antlr.ParseTreeWalkerDefault.Walk(&listener, parser.Start())

	return listener.Pop()
}
