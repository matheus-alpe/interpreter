package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/matheus-alpe/interpreter/internal/lexer"
	"github.com/matheus-alpe/interpreter/internal/parser"
)

const PROMPT = "> "

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
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "\nparser errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t")
		io.WriteString(out, msg)
		io.WriteString(out, "\n")
	}
}
